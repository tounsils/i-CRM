package store

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/tounsils/i-CRM/internal/domain"
)

func newTestStore(t *testing.T) *Store {
	t.Helper()
	// A real file rather than :memory:, so the test exercises the same open
	// path, pragmas and WAL behaviour the running server uses.
	s, err := Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	t.Cleanup(func() {
		if err := s.Close(); err != nil {
			t.Errorf("Close: %v", err)
		}
	})
	return s
}

func mustCustomer(t *testing.T, s *Store) int64 {
	t.Helper()
	id, err := s.CreateCustomer(context.Background(), Customer{
		Name:     "Northwind Trading",
		Email:    "ap@northwind.example",
		Status:   domain.CustomerLead,
		Currency: "USD",
	})
	if err != nil {
		t.Fatalf("CreateCustomer: %v", err)
	}
	return id
}

func sampleLines() []domain.Line {
	return []domain.Line{
		{Description: "Implementation", Quantity: 10 * domain.QuantityScale, UnitPrice: 15000,
			TaxName: "VAT", TaxRate: 2000},
		{Description: "Support", Quantity: 1 * domain.QuantityScale, UnitPrice: 4999,
			TaxName: "VAT", TaxRate: 2000},
	}
}

func mustDocument(t *testing.T, s *Store, docType domain.DocType, customerID int64) int64 {
	t.Helper()
	id, err := s.CreateDocument(context.Background(), Document{
		Type:       docType,
		CustomerID: customerID,
		Currency:   "USD",
		IssueDate:  time.Now(),
	}, sampleLines())
	if err != nil {
		t.Fatalf("CreateDocument(%s): %v", docType, err)
	}
	return id
}

func TestOpenIsIdempotent(t *testing.T) {
	// Migrations must be safe to run on every boot, because that is when they
	// run. A second Open over the same file re-applying 0001 would fail on the
	// duplicate CREATE TABLE, which is exactly the regression worth catching.
	dir := t.TempDir()
	path := filepath.Join(dir, "twice.db")

	first, err := Open(path)
	if err != nil {
		t.Fatalf("first Open: %v", err)
	}
	if err := first.Close(); err != nil {
		t.Fatalf("first Close: %v", err)
	}

	second, err := Open(path)
	if err != nil {
		t.Fatalf("second Open: %v", err)
	}
	if err := second.Close(); err != nil {
		t.Fatalf("second Close: %v", err)
	}
}

func TestForeignKeysAreEnforced(t *testing.T) {
	// SQLite defaults foreign_keys to OFF. If the pragma in the DSN ever stops
	// being applied, orphaned lines and payments become possible and nothing
	// else in the suite would notice, so this asserts the pragma directly.
	s := newTestStore(t)
	ctx := context.Background()

	var on int
	if err := s.db.QueryRowContext(ctx, `PRAGMA foreign_keys`).Scan(&on); err != nil {
		t.Fatalf("reading pragma: %v", err)
	}
	if on != 1 {
		t.Fatalf("foreign_keys = %d, want 1", on)
	}

	const orphan = `INSERT INTO document_lines (document_id, description) VALUES (9999, 'orphan')`
	if _, err := s.db.ExecContext(ctx, orphan); err == nil {
		t.Fatal("inserting a line against a non-existent document was allowed")
	}
}

func TestDocumentNumbersAreSequentialPerType(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()
	customerID := mustCustomer(t, s)

	var got []string
	for _, docType := range []domain.DocType{
		domain.DocInvoice, domain.DocInvoice, domain.DocEstimate, domain.DocInvoice, domain.DocProposal,
	} {
		id := mustDocument(t, s, docType, customerID)
		d, err := s.DocumentByID(ctx, id)
		if err != nil {
			t.Fatalf("DocumentByID: %v", err)
		}
		got = append(got, d.Number)
	}

	want := []string{"INV-0001", "INV-0002", "EST-0001", "INV-0003", "PRO-0001"}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("document %d numbered %q, want %q", i, got[i], want[i])
		}
	}
}

func TestCreateDocumentCachesTotalsThatMatchItsLines(t *testing.T) {
	// The cached columns have one writer, recomputeTotals. The invariant is
	// that the cache always equals a fresh computation over the stored lines.
	s := newTestStore(t)
	ctx := context.Background()
	id := mustDocument(t, s, domain.DocInvoice, mustCustomer(t, s))

	d, err := s.DocumentByID(ctx, id)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}

	priced := make([]domain.Line, len(d.Lines))
	for i, l := range d.Lines {
		priced[i] = l.Line
	}
	if fresh := domain.ComputeTotals(priced); fresh != d.Totals {
		t.Fatalf("cached totals %+v do not match a fresh computation %+v", d.Totals, fresh)
	}
	if d.Totals.Total != d.Totals.Subtotal+d.Totals.Tax {
		t.Fatalf("stored total %s is not subtotal %s plus tax %s",
			d.Totals.Total, d.Totals.Subtotal, d.Totals.Tax)
	}
}

func TestReplaceLinesRecomputesTotalsAndRefusesAfterSending(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()
	id := mustDocument(t, s, domain.DocInvoice, mustCustomer(t, s))

	replacement := []domain.Line{
		{Description: "One thing", Quantity: domain.QuantityScale, UnitPrice: 10000, TaxRate: 0},
	}
	if err := s.ReplaceLines(ctx, id, replacement); err != nil {
		t.Fatalf("ReplaceLines on a draft: %v", err)
	}
	d, err := s.DocumentByID(ctx, id)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}
	if len(d.Lines) != 1 || d.Totals.Total != 10000 {
		t.Fatalf("after replacing lines: %d lines, total %s; want 1 line and 100.00",
			len(d.Lines), d.Totals.Total)
	}

	if err := s.SetStatus(ctx, id, domain.StatusSent); err != nil {
		t.Fatalf("SetStatus sent: %v", err)
	}
	err = s.ReplaceLines(ctx, id, replacement)
	if !errors.Is(err, ErrNotEditable) {
		t.Fatalf("editing a sent invoice returned %v, want ErrNotEditable", err)
	}
}

func TestSetStatusRefusesAnIllegalMove(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()
	id := mustDocument(t, s, domain.DocInvoice, mustCustomer(t, s))

	// Draft straight to paid skips sending, and an invoice has no "accepted".
	if err := s.SetStatus(ctx, id, domain.StatusPaid); !errors.Is(err, domain.ErrIllegalTransition) {
		t.Fatalf("draft to paid returned %v, want ErrIllegalTransition", err)
	}
	if err := s.SetStatus(ctx, id, domain.StatusAccepted); !errors.Is(err, domain.ErrIllegalTransition) {
		t.Fatalf("invoice to accepted returned %v, want ErrIllegalTransition", err)
	}
}

func TestPaymentsDeriveInvoiceStatus(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()
	id := mustDocument(t, s, domain.DocInvoice, mustCustomer(t, s))

	if err := s.SetStatus(ctx, id, domain.StatusSent); err != nil {
		t.Fatalf("SetStatus sent: %v", err)
	}
	before, err := s.DocumentByID(ctx, id)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}

	part := before.Totals.Total / 4
	if _, err := s.RecordPayment(ctx, Payment{DocumentID: id, Amount: part, PaidOn: time.Now()}); err != nil {
		t.Fatalf("RecordPayment part: %v", err)
	}
	partial, err := s.DocumentByID(ctx, id)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}
	if partial.Status != domain.StatusPartiallyPaid {
		t.Fatalf("after a quarter payment status is %s, want %s", partial.Status, domain.StatusPartiallyPaid)
	}
	if partial.Outstanding() != before.Totals.Total-part {
		t.Fatalf("outstanding is %s, want %s", partial.Outstanding(), before.Totals.Total-part)
	}

	rest := before.Totals.Total - part
	if _, err := s.RecordPayment(ctx, Payment{DocumentID: id, Amount: rest, PaidOn: time.Now()}); err != nil {
		t.Fatalf("RecordPayment rest: %v", err)
	}
	paid, err := s.DocumentByID(ctx, id)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}
	if paid.Status != domain.StatusPaid {
		t.Fatalf("after full payment status is %s, want %s", paid.Status, domain.StatusPaid)
	}
	if paid.Outstanding() != 0 {
		t.Fatalf("outstanding after full payment is %s, want 0.00", paid.Outstanding())
	}
}

func TestPaymentsRefusedOnDraftVoidAndNonInvoices(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()
	customerID := mustCustomer(t, s)

	draft := mustDocument(t, s, domain.DocInvoice, customerID)
	pay := Payment{DocumentID: draft, Amount: 100, PaidOn: time.Now()}
	if _, err := s.RecordPayment(ctx, pay); !errors.Is(err, ErrNotPayable) {
		t.Fatalf("paying a draft returned %v, want ErrNotPayable", err)
	}

	estimate := mustDocument(t, s, domain.DocEstimate, customerID)
	pay.DocumentID = estimate
	if _, err := s.RecordPayment(ctx, pay); !errors.Is(err, ErrNotPayable) {
		t.Fatalf("paying an estimate returned %v, want ErrNotPayable", err)
	}

	voided := mustDocument(t, s, domain.DocInvoice, customerID)
	if err := s.SetStatus(ctx, voided, domain.StatusVoid); err != nil {
		t.Fatalf("SetStatus void: %v", err)
	}
	pay.DocumentID = voided
	if _, err := s.RecordPayment(ctx, pay); !errors.Is(err, ErrNotPayable) {
		t.Fatalf("paying a void invoice returned %v, want ErrNotPayable", err)
	}
}

func TestProposalConvertsThroughEstimateToInvoice(t *testing.T) {
	// The advertised product is this chain, so it gets an end-to-end test in
	// the store as well as at the HTTP layer.
	s := newTestStore(t)
	ctx := context.Background()
	customerID := mustCustomer(t, s)
	proposalID := mustDocument(t, s, domain.DocProposal, customerID)

	if _, err := s.Convert(ctx, proposalID, nil); !errors.Is(err, domain.ErrIllegalTransition) {
		t.Fatalf("converting a draft proposal returned %v, want ErrIllegalTransition", err)
	}

	for _, to := range []domain.DocStatus{domain.StatusSent, domain.StatusAccepted} {
		if err := s.SetStatus(ctx, proposalID, to); err != nil {
			t.Fatalf("SetStatus %s: %v", to, err)
		}
	}

	estimateID, err := s.Convert(ctx, proposalID, nil)
	if err != nil {
		t.Fatalf("Convert proposal: %v", err)
	}
	proposal, err := s.DocumentByID(ctx, proposalID)
	if err != nil {
		t.Fatalf("DocumentByID proposal: %v", err)
	}
	if proposal.Status != domain.StatusConverted {
		t.Fatalf("source proposal is %s, want converted", proposal.Status)
	}
	if proposal.ConvertedToID == nil || *proposal.ConvertedToID != estimateID {
		t.Fatalf("proposal.ConvertedToID = %v, want %d", proposal.ConvertedToID, estimateID)
	}

	estimate, err := s.DocumentByID(ctx, estimateID)
	if err != nil {
		t.Fatalf("DocumentByID estimate: %v", err)
	}
	if estimate.Type != domain.DocEstimate || estimate.Status != domain.StatusDraft {
		t.Fatalf("converted document is %s/%s, want estimate/draft", estimate.Type, estimate.Status)
	}
	if estimate.Totals != proposal.Totals {
		t.Fatalf("totals changed across conversion: %+v became %+v", proposal.Totals, estimate.Totals)
	}
	if len(estimate.Lines) != len(proposal.Lines) {
		t.Fatalf("estimate has %d lines, proposal had %d", len(estimate.Lines), len(proposal.Lines))
	}

	for _, to := range []domain.DocStatus{domain.StatusSent, domain.StatusAccepted} {
		if err := s.SetStatus(ctx, estimateID, to); err != nil {
			t.Fatalf("SetStatus %s on estimate: %v", to, err)
		}
	}
	due := time.Now().Add(30 * 24 * time.Hour)
	invoiceID, err := s.Convert(ctx, estimateID, &due)
	if err != nil {
		t.Fatalf("Convert estimate: %v", err)
	}
	invoice, err := s.DocumentByID(ctx, invoiceID)
	if err != nil {
		t.Fatalf("DocumentByID invoice: %v", err)
	}
	if invoice.Type != domain.DocInvoice {
		t.Fatalf("converted estimate produced a %s, want an invoice", invoice.Type)
	}
	if invoice.DueDate == nil {
		t.Fatal("invoice has no due date, but one was supplied to Convert")
	}
	if _, err := s.Convert(ctx, invoiceID, nil); !errors.Is(err, domain.ErrNotConvertible) {
		t.Fatalf("converting an invoice returned %v, want ErrNotConvertible", err)
	}
}

func TestOutstandingTotalCountsOnlyLiveInvoices(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()
	customerID := mustCustomer(t, s)

	// A draft invoice and an estimate are not money anybody owes.
	mustDocument(t, s, domain.DocInvoice, customerID)
	mustDocument(t, s, domain.DocEstimate, customerID)

	sent := mustDocument(t, s, domain.DocInvoice, customerID)
	if err := s.SetStatus(ctx, sent, domain.StatusSent); err != nil {
		t.Fatalf("SetStatus sent: %v", err)
	}
	d, err := s.DocumentByID(ctx, sent)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}

	outstanding, err := s.OutstandingTotal(ctx)
	if err != nil {
		t.Fatalf("OutstandingTotal: %v", err)
	}
	if outstanding != d.Totals.Total {
		t.Fatalf("outstanding = %s, want the one sent invoice's %s", outstanding, d.Totals.Total)
	}

	if _, err := s.RecordPayment(ctx, Payment{
		DocumentID: sent, Amount: d.Totals.Total, PaidOn: time.Now(),
	}); err != nil {
		t.Fatalf("RecordPayment: %v", err)
	}
	outstanding, err = s.OutstandingTotal(ctx)
	if err != nil {
		t.Fatalf("OutstandingTotal after payment: %v", err)
	}
	if outstanding != 0 {
		t.Fatalf("outstanding after paying the only live invoice = %s, want 0.00", outstanding)
	}
}

func TestUnsettledInvoicesExcludeTheSettledAndUnissued(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()
	customerID := mustCustomer(t, s)

	issued := time.Now().AddDate(0, 0, -40)
	past := time.Now().AddDate(0, 0, -10)
	future := time.Now().AddDate(0, 0, 30)

	invoice := func(due time.Time) int64 {
		t.Helper()
		id, err := s.CreateDocument(ctx, Document{
			Type:       domain.DocInvoice,
			CustomerID: customerID,
			Currency:   "USD",
			IssueDate:  issued,
			DueDate:    &due,
		}, sampleLines())
		if err != nil {
			t.Fatalf("CreateDocument: %v", err)
		}
		return id
	}

	late := invoice(past)
	notYetDue := invoice(future)
	settled := invoice(past)
	neverIssued := invoice(past)

	for _, id := range []int64{late, notYetDue, settled} {
		if err := s.SetStatus(ctx, id, domain.StatusSent); err != nil {
			t.Fatalf("SetStatus sent: %v", err)
		}
	}
	paid, err := s.DocumentByID(ctx, settled)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}
	if _, err := s.RecordPayment(ctx, Payment{
		DocumentID: settled, Amount: paid.Totals.Total, PaidOn: time.Now(),
	}); err != nil {
		t.Fatalf("RecordPayment: %v", err)
	}

	unsettled, err := s.ListUnsettledInvoices(ctx)
	if err != nil {
		t.Fatalf("ListUnsettledInvoices: %v", err)
	}
	if len(unsettled) != 2 {
		t.Fatalf("got %d unsettled invoices, want 2 (the sent ones that are not paid)", len(unsettled))
	}

	byID := make(map[int64]Document, len(unsettled))
	for _, d := range unsettled {
		byID[d.ID] = d
	}
	if _, ok := byID[neverIssued]; ok {
		t.Error("a draft invoice appeared in the unsettled list despite its due date having passed")
	}
	if _, ok := byID[settled]; ok {
		t.Error("a paid invoice appeared in the unsettled list")
	}

	// Soonest due first, which is the order the dashboard relies on.
	if unsettled[0].ID != late {
		t.Errorf("unsettled[0] is %d, want the earlier-due invoice %d", unsettled[0].ID, late)
	}

	asOf := time.Now()
	if !byID[late].Overdue(asOf) {
		t.Error("an invoice sent and due ten days ago is not reported overdue")
	}
	if byID[notYetDue].Overdue(asOf) {
		t.Error("an invoice due in thirty days is reported overdue")
	}
	if paidAgain, err := s.DocumentByID(ctx, settled); err != nil {
		t.Fatalf("DocumentByID: %v", err)
	} else if paidAgain.Overdue(asOf) {
		t.Error("a paid invoice is reported overdue because its due date passed")
	}
}

func TestCustomerAndMessageTimeline(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()
	customerID := mustCustomer(t, s)

	authorID, err := s.CreateUser(ctx, User{
		Email: "staff@example.com", Name: "Staff", PasswordHash: "x", Role: "admin",
	})
	if err != nil {
		t.Fatalf("CreateUser: %v", err)
	}

	for _, d := range []domain.MessageDirection{domain.MessageOutbound, domain.MessageInbound} {
		_, err := s.CreateMessage(ctx, Message{
			CustomerID: customerID, AuthorID: authorID, Direction: d,
			Subject: "Proposal", Body: "Body of " + string(d),
		})
		if err != nil {
			t.Fatalf("CreateMessage %s: %v", d, err)
		}
	}

	timeline, err := s.MessagesForCustomer(ctx, customerID)
	if err != nil {
		t.Fatalf("MessagesForCustomer: %v", err)
	}
	if len(timeline) != 2 {
		t.Fatalf("timeline has %d entries, want 2", len(timeline))
	}
	if timeline[0].AuthorName != "Staff" {
		t.Fatalf("author name = %q, want %q", timeline[0].AuthorName, "Staff")
	}

	promoted, err := s.CustomerByID(ctx, customerID)
	if err != nil {
		t.Fatalf("CustomerByID: %v", err)
	}
	promoted.Status = domain.CustomerActive
	if err := s.UpdateCustomer(ctx, promoted); err != nil {
		t.Fatalf("UpdateCustomer: %v", err)
	}
	counts, err := s.CountCustomersByStatus(ctx)
	if err != nil {
		t.Fatalf("CountCustomersByStatus: %v", err)
	}
	if counts[domain.CustomerActive] != 1 || counts[domain.CustomerLead] != 0 {
		t.Fatalf("counts = %v, want one active and no leads", counts)
	}
}

func TestMissingRecordsReportNotFound(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()

	if _, err := s.CustomerByID(ctx, 404); !errors.Is(err, ErrNotFound) {
		t.Errorf("CustomerByID(404) = %v, want ErrNotFound", err)
	}
	if _, err := s.DocumentByID(ctx, 404); !errors.Is(err, ErrNotFound) {
		t.Errorf("DocumentByID(404) = %v, want ErrNotFound", err)
	}
	if _, err := s.UserByEmail(ctx, "nobody@example.com"); !errors.Is(err, ErrNotFound) {
		t.Errorf("UserByEmail = %v, want ErrNotFound", err)
	}
	if err := s.SetStatus(ctx, 404, domain.StatusSent); !errors.Is(err, ErrNotFound) {
		t.Errorf("SetStatus(404) = %v, want ErrNotFound", err)
	}
}
