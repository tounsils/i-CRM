package domain

import (
	"errors"
	"fmt"
)

// DocType is the kind of document. Proposals, estimates and invoices are one
// type with three values rather than three near-identical tables, because the
// product is the lifecycle between them: a proposal becomes an estimate becomes
// an invoice, carrying its lines along.
type DocType string

const (
	DocProposal DocType = "proposal"
	DocEstimate DocType = "estimate"
	DocInvoice  DocType = "invoice"
)

// DocStatus is the position of a document in its lifecycle. The legal moves are
// declared in transitions below, not scattered across handlers.
type DocStatus string

const (
	StatusDraft         DocStatus = "draft"
	StatusSent          DocStatus = "sent"
	StatusAccepted      DocStatus = "accepted"
	StatusDeclined      DocStatus = "declined"
	StatusConverted     DocStatus = "converted"
	StatusPartiallyPaid DocStatus = "partially_paid"
	StatusPaid          DocStatus = "paid"
	StatusVoid          DocStatus = "void"
)

var (
	ErrUnknownDocType    = errors.New("domain: unknown document type")
	ErrIllegalTransition = errors.New("domain: illegal status transition")
	ErrNotConvertible    = errors.New("domain: document cannot be converted")
)

// ParseDocType validates an untrusted string against the closed set, so a
// handler never writes an unknown type into the store.
func ParseDocType(s string) (DocType, error) {
	t := DocType(s)
	switch t {
	case DocProposal, DocEstimate, DocInvoice:
		return t, nil
	}
	return "", fmt.Errorf("%w: %q", ErrUnknownDocType, s)
}

// Label is the human name for the type, for headings and buttons.
func (t DocType) Label() string {
	switch t {
	case DocProposal:
		return "Proposal"
	case DocEstimate:
		return "Estimate"
	case DocInvoice:
		return "Invoice"
	}
	return string(t)
}

// Label is the human name for the status.
func (s DocStatus) Label() string {
	switch s {
	case StatusDraft:
		return "Draft"
	case StatusSent:
		return "Sent"
	case StatusAccepted:
		return "Accepted"
	case StatusDeclined:
		return "Declined"
	case StatusConverted:
		return "Converted"
	case StatusPartiallyPaid:
		return "Partially paid"
	case StatusPaid:
		return "Paid"
	case StatusVoid:
		return "Void"
	}
	return string(s)
}

// transitions maps each type's current status to the statuses it may move to.
// A status present with an empty slice is terminal and says so deliberately;
// a status absent from the map is a programming error, not a terminal state.
var transitions = map[DocType]map[DocStatus][]DocStatus{
	DocProposal: {
		StatusDraft:     {StatusSent, StatusVoid},
		StatusSent:      {StatusAccepted, StatusDeclined, StatusVoid},
		StatusAccepted:  {StatusConverted, StatusVoid},
		StatusDeclined:  {},
		StatusConverted: {},
		StatusVoid:      {},
	},
	DocEstimate: {
		StatusDraft:     {StatusSent, StatusVoid},
		StatusSent:      {StatusAccepted, StatusDeclined, StatusVoid},
		StatusAccepted:  {StatusConverted, StatusVoid},
		StatusDeclined:  {},
		StatusConverted: {},
		StatusVoid:      {},
	},
	DocInvoice: {
		StatusDraft:         {StatusSent, StatusVoid},
		StatusSent:          {StatusPartiallyPaid, StatusPaid, StatusVoid},
		StatusPartiallyPaid: {StatusPaid, StatusVoid},
		StatusPaid:          {},
		StatusVoid:          {},
	},
}

// InitialStatus is the status every document is created in. Nothing is ever
// created already sent, so there is always a point at which lines can be fixed.
const InitialStatus = StatusDraft

// CanTransition reports whether a document of this type may move between these
// statuses.
func CanTransition(t DocType, from, to DocStatus) bool {
	byStatus, ok := transitions[t]
	if !ok {
		return false
	}
	for _, allowed := range byStatus[from] {
		if allowed == to {
			return true
		}
	}
	return false
}

// Transition returns the new status, or an error naming both ends of the move
// it refused. Handlers call this instead of assigning a status directly.
func Transition(t DocType, from, to DocStatus) (DocStatus, error) {
	if !CanTransition(t, from, to) {
		return from, fmt.Errorf("%w: %s cannot go from %s to %s", ErrIllegalTransition, t, from, to)
	}
	return to, nil
}

// NextStatuses lists the moves currently open to a document, so the UI offers
// exactly the buttons the rules allow rather than a fixed set it then rejects.
func NextStatuses(t DocType, from DocStatus) []DocStatus {
	byStatus, ok := transitions[t]
	if !ok {
		return nil
	}
	out := make([]DocStatus, len(byStatus[from]))
	copy(out, byStatus[from])
	return out
}

// ConversionTarget gives the type a document converts into: a proposal becomes
// an estimate, an estimate becomes an invoice. Invoices are the end of the
// chain and convert into nothing.
func ConversionTarget(t DocType) (DocType, error) {
	switch t {
	case DocProposal:
		return DocEstimate, nil
	case DocEstimate:
		return DocInvoice, nil
	case DocInvoice:
		return "", fmt.Errorf("%w: an invoice is the end of the chain", ErrNotConvertible)
	}
	return "", fmt.Errorf("%w: %q", ErrUnknownDocType, t)
}

// Line is one priced row of a document. The tax name and rate are copied onto
// the line rather than referenced, because editing a tax rate next year must
// not silently restate an invoice that was already sent.
type Line struct {
	Description string
	Quantity    Quantity
	UnitPrice   Money
	TaxName     string
	TaxRate     Rate
}

// LineTotals is one line's contribution to the document.
type LineTotals struct {
	Subtotal Money
	Tax      Money
}

// Total is what the line adds to the document total.
func (lt LineTotals) Total() Money {
	return lt.Subtotal + lt.Tax
}

// Totals computes this line in isolation.
func (l Line) Totals() LineTotals {
	subtotal := l.UnitPrice.MulQuantity(l.Quantity)
	return LineTotals{Subtotal: subtotal, Tax: l.TaxRate.Apply(subtotal)}
}

// Totals is the footer of a document.
type Totals struct {
	Subtotal Money
	Tax      Money
	Total    Money
}

// ComputeTotals sums the lines. See Rate.Apply for why tax is rounded per line
// and then added, rather than taken once on the subtotal.
func ComputeTotals(lines []Line) Totals {
	var t Totals
	for _, l := range lines {
		lt := l.Totals()
		t.Subtotal += lt.Subtotal
		t.Tax += lt.Tax
	}
	t.Total = t.Subtotal + t.Tax
	return t
}

// StatusForPayments is the status an invoice holds given its total and the sum
// of payments recorded against it. Payment is the one place where status is
// derived from data rather than chosen by a user, so overpayment and exact
// payment both land on paid and neither needs a human decision.
//
// A void invoice stays void: voiding is an accounting act, and a payment
// arriving afterwards is a refund to handle, not a reason to revive the
// document.
func StatusForPayments(total, paid Money, current DocStatus) DocStatus {
	if current == StatusVoid {
		return StatusVoid
	}
	switch {
	case paid <= 0:
		return current
	case paid >= total:
		return StatusPaid
	default:
		return StatusPartiallyPaid
	}
}
