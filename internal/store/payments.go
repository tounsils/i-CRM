package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/tounsils/i-CRM/internal/domain"
)

// ErrNotPayable is returned when a payment is recorded against something that
// cannot receive one: anything that is not an invoice, a draft that has not
// been sent, or a voided invoice.
var ErrNotPayable = errors.New("store: document cannot take a payment")

// RecordPayment inserts a payment and re-derives the invoice status from the
// payments actually on file.
//
// The status is derived rather than passed in, which is what stops a
// part-payment from being marked paid by a mistaken click: the only way to
// reach paid is for the sum of payments to reach the total.
func (s *Store) RecordPayment(ctx context.Context, p Payment) (int64, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("store: recording payment: %w", err)
	}
	defer tx.Rollback() //nolint:errcheck // no-op once Commit has succeeded

	docType, current, err := documentTypeAndStatus(ctx, tx, p.DocumentID)
	if err != nil {
		return 0, err
	}
	if docType != domain.DocInvoice {
		return 0, fmt.Errorf("%w: %s %d is not an invoice", ErrNotPayable, docType, p.DocumentID)
	}
	if current == domain.StatusDraft || current == domain.StatusVoid {
		return 0, fmt.Errorf("%w: invoice %d is %s", ErrNotPayable, p.DocumentID, current)
	}

	const insert = `INSERT INTO payments
	                (document_id, amount_minor, method, reference, paid_on, created_at)
	                VALUES (?, ?, ?, ?, ?, ?)`
	res, err := tx.ExecContext(ctx, insert, p.DocumentID, int64(p.Amount),
		p.Method, p.Reference, formatTime(p.PaidOn), formatTime(time.Now()))
	if err != nil {
		return 0, fmt.Errorf("store: recording payment: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("store: recording payment: %w", err)
	}

	if err := refreshPaymentStatus(ctx, tx, p.DocumentID, docType, current); err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("store: recording payment: %w", err)
	}
	return id, nil
}

// refreshPaymentStatus re-derives and stores an invoice's status from its total
// and the payments on file.
func refreshPaymentStatus(
	ctx context.Context, tx *sql.Tx, docID int64, docType domain.DocType, current domain.DocStatus,
) error {
	const totals = `SELECT d.total_minor,
	                       COALESCE((SELECT SUM(p.amount_minor) FROM payments p
	                                 WHERE p.document_id = d.id), 0)
	                FROM documents d WHERE d.id = ?`
	var total, paid int64
	if err := tx.QueryRowContext(ctx, totals, docID).Scan(&total, &paid); err != nil {
		return fmt.Errorf("store: totalling payments for document %d: %w", docID, err)
	}

	next := domain.StatusForPayments(domain.Money(total), domain.Money(paid), current)
	if next == current {
		return nil
	}
	// The derived status still has to be a legal move, so the lifecycle has one
	// owner. A payment can take a sent invoice to partially_paid or paid, and
	// nothing else.
	if _, err := domain.Transition(docType, current, next); err != nil {
		return err
	}
	const q = `UPDATE documents SET status = ? WHERE id = ?`
	if _, err := tx.ExecContext(ctx, q, string(next), docID); err != nil {
		return fmt.Errorf("store: updating status of document %d: %w", docID, err)
	}
	return nil
}

// PaymentsForDocument lists the payments on file, oldest first.
func (s *Store) PaymentsForDocument(ctx context.Context, docID int64) ([]Payment, error) {
	const q = `SELECT id, document_id, amount_minor, method, reference, paid_on, created_at
	           FROM payments WHERE document_id = ? ORDER BY paid_on, id`
	rows, err := s.db.QueryContext(ctx, q, docID)
	if err != nil {
		return nil, fmt.Errorf("store: listing payments for document %d: %w", docID, err)
	}
	defer rows.Close()

	var out []Payment
	for rows.Next() {
		var (
			p         Payment
			amount    int64
			paidOn    string
			createdAt string
		)
		err := rows.Scan(&p.ID, &p.DocumentID, &amount, &p.Method, &p.Reference, &paidOn, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("store: reading a payment: %w", err)
		}
		p.Amount = domain.Money(amount)
		if p.PaidOn, err = parseTime(paidOn); err != nil {
			return nil, fmt.Errorf("store: reading payment date: %w", err)
		}
		if p.CreatedAt, err = parseTime(createdAt); err != nil {
			return nil, fmt.Errorf("store: reading payment timestamp: %w", err)
		}
		out = append(out, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("store: listing payments for document %d: %w", docID, err)
	}
	return out, nil
}

// OutstandingTotal is the sum still owed across all live invoices, for the
// dashboard. Draft and void invoices are excluded because neither is money
// anybody owes.
func (s *Store) OutstandingTotal(ctx context.Context) (domain.Money, error) {
	placeholders, args := unsettledInvoiceArgs()
	q := `SELECT COALESCE(SUM(d.total_minor -
	          (SELECT COALESCE(SUM(p.amount_minor), 0) FROM payments p
	           WHERE p.document_id = d.id)), 0)
	      FROM documents d
	      WHERE d.type = ? AND d.status IN (` + placeholders + `)`

	var outstanding int64
	if err := s.db.QueryRowContext(ctx, q, args...).Scan(&outstanding); err != nil {
		return 0, fmt.Errorf("store: totalling outstanding invoices: %w", err)
	}
	return domain.Money(outstanding), nil
}
