package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/tounsils/i-CRM/internal/domain"
)

// ErrNotEditable is returned when lines are changed on a document that has left
// draft. An issued document is a record of what was sent, so it is not edited.
var ErrNotEditable = errors.New("store: document is no longer editable")

const documentSelect = `
SELECT d.id, d.type, d.number, d.customer_id, d.status, d.currency,
       d.issue_date, d.due_date, d.notes,
       d.subtotal_minor, d.tax_minor, d.total_minor,
       d.converted_to_id, d.created_at,
       c.name,
       COALESCE((SELECT SUM(p.amount_minor) FROM payments p WHERE p.document_id = d.id), 0)
FROM documents d
JOIN customers c ON c.id = d.customer_id`

// DocumentFilter narrows a document list. Zero values mean no filter, so one
// method serves the invoice list, the proposal list and a customer's history
// rather than three near-identical queries.
type DocumentFilter struct {
	Type       domain.DocType
	Status     domain.DocStatus
	CustomerID int64
}

// querier is satisfied by *sql.DB and *sql.Tx, so the line and payment reads
// work both inside a transaction and outside one.
type querier interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

// CreateDocument inserts a document with its lines and returns the new id.
//
// The status is forced to domain.InitialStatus rather than taken from the
// argument: nothing is ever created already sent, so there is always a point at
// which the lines can still be corrected.
func (s *Store) CreateDocument(ctx context.Context, d Document, lines []domain.Line) (int64, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("store: creating document: %w", err)
	}
	defer tx.Rollback() //nolint:errcheck // no-op once Commit has succeeded

	id, err := insertDocument(ctx, tx, d, lines)
	if err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("store: creating document: %w", err)
	}
	return id, nil
}

// insertDocument writes one document and its lines inside an existing
// transaction. Create and Convert both need this, and sharing it keeps
// numbering and totals from having two implementations.
func insertDocument(ctx context.Context, tx *sql.Tx, d Document, lines []domain.Line) (int64, error) {
	number, err := nextDocumentNumber(ctx, tx, d.Type)
	if err != nil {
		return 0, err
	}

	const q = `INSERT INTO documents
	           (type, number, customer_id, status, currency, issue_date, due_date, notes, created_at)
	           VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := tx.ExecContext(ctx, q,
		string(d.Type), number, d.CustomerID, string(domain.InitialStatus), d.Currency,
		formatTime(d.IssueDate), nullTime(d.DueDate), d.Notes, formatTime(time.Now()))
	if err != nil {
		return 0, fmt.Errorf("store: creating document: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("store: creating document: %w", err)
	}

	if err := insertLines(ctx, tx, id, lines); err != nil {
		return 0, err
	}
	if err := recomputeTotals(ctx, tx, id); err != nil {
		return 0, err
	}
	return id, nil
}

// nextDocumentNumber allocates the next reference for a type from its counter.
func nextDocumentNumber(ctx context.Context, tx *sql.Tx, t domain.DocType) (string, error) {
	const q = `INSERT INTO document_counters (type, next_seq) VALUES (?, 1)
	           ON CONFLICT (type) DO UPDATE SET next_seq = next_seq + 1
	           RETURNING next_seq`
	var seq int64
	if err := tx.QueryRowContext(ctx, q, string(t)).Scan(&seq); err != nil {
		return "", fmt.Errorf("store: allocating a %s number: %w", t, err)
	}
	return fmt.Sprintf("%s-%04d", numberPrefix(t), seq), nil
}

func numberPrefix(t domain.DocType) string {
	switch t {
	case domain.DocProposal:
		return "PRO"
	case domain.DocEstimate:
		return "EST"
	case domain.DocInvoice:
		return "INV"
	}
	return "DOC"
}

func insertLines(ctx context.Context, tx *sql.Tx, docID int64, lines []domain.Line) error {
	const q = `INSERT INTO document_lines
	           (document_id, position, description, quantity_milli, unit_price_minor, tax_name, tax_rate_bp)
	           VALUES (?, ?, ?, ?, ?, ?, ?)`
	for i, l := range lines {
		_, err := tx.ExecContext(ctx, q, docID, i, l.Description,
			int64(l.Quantity), int64(l.UnitPrice), l.TaxName, int64(l.TaxRate))
		if err != nil {
			return fmt.Errorf("store: inserting line %d of document %d: %w", i, docID, err)
		}
	}
	return nil
}

// recomputeTotals is the only writer of the cached total columns. Every path
// that changes lines calls it in the same transaction, so the cache cannot
// disagree with document_lines.
func recomputeTotals(ctx context.Context, tx *sql.Tx, docID int64) error {
	lines, err := linesForDocument(ctx, tx, docID)
	if err != nil {
		return err
	}
	priced := make([]domain.Line, len(lines))
	for i, l := range lines {
		priced[i] = l.Line
	}
	totals := domain.ComputeTotals(priced)

	const q = `UPDATE documents SET subtotal_minor = ?, tax_minor = ?, total_minor = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, q,
		int64(totals.Subtotal), int64(totals.Tax), int64(totals.Total), docID)
	if err != nil {
		return fmt.Errorf("store: caching totals for document %d: %w", docID, err)
	}
	return nil
}

func linesForDocument(ctx context.Context, q querier, docID int64) ([]Line, error) {
	const sel = `SELECT id, document_id, position, description, quantity_milli,
	                    unit_price_minor, tax_name, tax_rate_bp
	             FROM document_lines WHERE document_id = ? ORDER BY position, id`
	rows, err := q.QueryContext(ctx, sel, docID)
	if err != nil {
		return nil, fmt.Errorf("store: reading lines of document %d: %w", docID, err)
	}
	defer rows.Close()

	var out []Line
	for rows.Next() {
		var (
			l     Line
			qty   int64
			price int64
			rate  int64
		)
		err := rows.Scan(&l.ID, &l.DocumentID, &l.Position, &l.Description,
			&qty, &price, &l.TaxName, &rate)
		if err != nil {
			return nil, fmt.Errorf("store: reading a line of document %d: %w", docID, err)
		}
		l.Quantity = domain.Quantity(qty)
		l.UnitPrice = domain.Money(price)
		l.TaxRate = domain.Rate(rate)
		out = append(out, l)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("store: reading lines of document %d: %w", docID, err)
	}
	return out, nil
}

// DocumentByID loads one document with its lines and the sum of its payments.
func (s *Store) DocumentByID(ctx context.Context, id int64) (Document, error) {
	d, err := scanDocument(s.db.QueryRowContext(ctx, documentSelect+` WHERE d.id = ?`, id))
	if err != nil {
		return Document{}, err
	}
	if d.Lines, err = linesForDocument(ctx, s.db, id); err != nil {
		return Document{}, err
	}
	return d, nil
}

// ListDocuments returns matching documents, newest first, without their lines.
func (s *Store) ListDocuments(ctx context.Context, f DocumentFilter) ([]Document, error) {
	q := documentSelect
	args := []any{}
	where := ""
	addClause := func(clause string, arg any) {
		if where == "" {
			where = ` WHERE ` + clause
		} else {
			where += ` AND ` + clause
		}
		args = append(args, arg)
	}
	if f.Type != "" {
		addClause(`d.type = ?`, string(f.Type))
	}
	if f.Status != "" {
		addClause(`d.status = ?`, string(f.Status))
	}
	if f.CustomerID != 0 {
		addClause(`d.customer_id = ?`, f.CustomerID)
	}
	q += where + ` ORDER BY d.issue_date DESC, d.id DESC`

	rows, err := s.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("store: listing documents: %w", err)
	}
	defer rows.Close()

	var out []Document
	for rows.Next() {
		d, err := scanDocumentRows(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, d)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("store: listing documents: %w", err)
	}
	return out, nil
}

// ReplaceLines swaps the lines of a draft document and recomputes its totals.
func (s *Store) ReplaceLines(ctx context.Context, docID int64, lines []domain.Line) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("store: replacing lines of document %d: %w", docID, err)
	}
	defer tx.Rollback() //nolint:errcheck // no-op once Commit has succeeded

	docType, current, err := documentTypeAndStatus(ctx, tx, docID)
	if err != nil {
		return err
	}
	if current != domain.StatusDraft {
		return fmt.Errorf("%w: %s %d is %s", ErrNotEditable, docType, docID, current)
	}

	if _, err := tx.ExecContext(ctx, `DELETE FROM document_lines WHERE document_id = ?`, docID); err != nil {
		return fmt.Errorf("store: clearing lines of document %d: %w", docID, err)
	}
	if err := insertLines(ctx, tx, docID, lines); err != nil {
		return err
	}
	if err := recomputeTotals(ctx, tx, docID); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("store: replacing lines of document %d: %w", docID, err)
	}
	return nil
}

// SetStatus moves a document through its lifecycle. The move is checked by
// domain.Transition, so an illegal one is refused here rather than in a handler.
func (s *Store) SetStatus(ctx context.Context, docID int64, to domain.DocStatus) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("store: setting status of document %d: %w", docID, err)
	}
	defer tx.Rollback() //nolint:errcheck // no-op once Commit has succeeded

	docType, current, err := documentTypeAndStatus(ctx, tx, docID)
	if err != nil {
		return err
	}
	next, err := domain.Transition(docType, current, to)
	if err != nil {
		return err
	}

	const q = `UPDATE documents SET status = ? WHERE id = ?`
	res, err := tx.ExecContext(ctx, q, string(next), docID)
	if err != nil {
		return fmt.Errorf("store: setting status of document %d: %w", docID, err)
	}
	if err := requireOneRow(res, fmt.Sprintf("document %d", docID)); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("store: setting status of document %d: %w", docID, err)
	}
	return nil
}

// Convert turns an accepted proposal into an estimate, or an accepted estimate
// into an invoice, copying the lines across. The source is marked converted and
// linked to the new document so the chain stays walkable.
//
// dueDate is passed in rather than derived here because payment terms are a
// configuration value owned by the config package, and the store must not grow
// a second opinion about them.
func (s *Store) Convert(ctx context.Context, docID int64, dueDate *time.Time) (int64, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("store: converting document %d: %w", docID, err)
	}
	defer tx.Rollback() //nolint:errcheck // no-op once Commit has succeeded

	source, err := scanDocument(tx.QueryRowContext(ctx, documentSelect+` WHERE d.id = ?`, docID))
	if err != nil {
		return 0, err
	}
	target, err := domain.ConversionTarget(source.Type)
	if err != nil {
		return 0, err
	}
	if source.Status != domain.StatusAccepted {
		return 0, fmt.Errorf("%w: only an accepted %s converts, and %s is %s",
			domain.ErrIllegalTransition, source.Type, source.Number, source.Status)
	}

	lines, err := linesForDocument(ctx, tx, docID)
	if err != nil {
		return 0, err
	}
	priced := make([]domain.Line, len(lines))
	for i, l := range lines {
		priced[i] = l.Line
	}

	newID, err := insertDocument(ctx, tx, Document{
		Type:       target,
		CustomerID: source.CustomerID,
		Currency:   source.Currency,
		IssueDate:  time.Now(),
		DueDate:    dueDate,
		Notes:      source.Notes,
	}, priced)
	if err != nil {
		return 0, err
	}

	if _, err := domain.Transition(source.Type, source.Status, domain.StatusConverted); err != nil {
		return 0, err
	}
	const markConverted = `UPDATE documents SET status = ?, converted_to_id = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, markConverted, string(domain.StatusConverted), newID, docID)
	if err != nil {
		return 0, fmt.Errorf("store: marking document %d converted: %w", docID, err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("store: converting document %d: %w", docID, err)
	}
	return newID, nil
}

func documentTypeAndStatus(
	ctx context.Context, q querier, docID int64,
) (domain.DocType, domain.DocStatus, error) {
	var docType, status string
	const sel = `SELECT type, status FROM documents WHERE id = ?`
	err := q.QueryRowContext(ctx, sel, docID).Scan(&docType, &status)
	if errors.Is(err, sql.ErrNoRows) {
		return "", "", fmt.Errorf("%w: document %d", ErrNotFound, docID)
	}
	if err != nil {
		return "", "", fmt.Errorf("store: reading document %d: %w", docID, err)
	}
	return domain.DocType(docType), domain.DocStatus(status), nil
}

func scanDocument(row *sql.Row) (Document, error) {
	d, err := scanDocumentRows(row)
	if errors.Is(err, sql.ErrNoRows) {
		return Document{}, ErrNotFound
	}
	return d, err
}

func scanDocumentRows(row rowScanner) (Document, error) {
	var (
		d           Document
		docType     string
		status      string
		issueDate   string
		dueDate     sql.NullString
		subtotal    int64
		tax         int64
		total       int64
		convertedTo sql.NullInt64
		createdAt   string
		paid        int64
	)
	err := row.Scan(&d.ID, &docType, &d.Number, &d.CustomerID, &status, &d.Currency,
		&issueDate, &dueDate, &d.Notes, &subtotal, &tax, &total,
		&convertedTo, &createdAt, &d.CustomerName, &paid)
	if errors.Is(err, sql.ErrNoRows) {
		return Document{}, sql.ErrNoRows
	}
	if err != nil {
		return Document{}, fmt.Errorf("store: reading document: %w", err)
	}

	d.Type = domain.DocType(docType)
	d.Status = domain.DocStatus(status)
	d.Totals = domain.Totals{
		Subtotal: domain.Money(subtotal),
		Tax:      domain.Money(tax),
		Total:    domain.Money(total),
	}
	d.ConvertedToID = idFromNull(convertedTo)
	d.Paid = domain.Money(paid)

	if d.IssueDate, err = parseTime(issueDate); err != nil {
		return Document{}, fmt.Errorf("store: reading document issue date: %w", err)
	}
	if d.DueDate, err = parseNullTime(dueDate); err != nil {
		return Document{}, fmt.Errorf("store: reading document due date: %w", err)
	}
	if d.CreatedAt, err = parseTime(createdAt); err != nil {
		return Document{}, fmt.Errorf("store: reading document timestamp: %w", err)
	}
	return d, nil
}
