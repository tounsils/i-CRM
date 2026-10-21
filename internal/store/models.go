package store

import (
	"time"

	"github.com/tounsils/i-CRM/internal/domain"
)

// User is a staff member who can sign in.
type User struct {
	ID           int64
	Email        string
	Name         string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}

// Customer covers the whole pipeline. A lead and a customer are one record at
// different Status values, so converting a lead keeps its history instead of
// copying it into a second table.
type Customer struct {
	ID          int64
	Name        string
	ContactName string
	Email       string
	Phone       string
	Address     string
	Status      domain.CustomerStatus
	Currency    string
	Notes       string
	CreatedAt   time.Time
}

// Tax is a named rate that can be applied to a line.
type Tax struct {
	ID       int64
	Name     string
	Rate     domain.Rate
	Archived bool
}

// Item is a catalogue entry used to fill a line quickly. Its price is a
// starting point: once copied onto a line, the line owns the figure.
type Item struct {
	ID          int64
	Name        string
	Description string
	UnitPrice   domain.Money
	TaxID       *int64
	Archived    bool
}

// Line is a stored document line. The priced part is domain.Line so that the
// totals code works on exactly the same values the database holds.
type Line struct {
	ID         int64
	DocumentID int64
	Position   int
	domain.Line
}

// Document is a proposal, estimate or invoice.
//
// Totals is read from the cached columns on the row rather than recomputed on
// load, so a list of 200 invoices costs one query. document_lines remains the
// source of truth and recomputeTotals is the only writer of those columns.
type Document struct {
	ID            int64
	Type          domain.DocType
	Number        string
	CustomerID    int64
	Status        domain.DocStatus
	Currency      string
	IssueDate     time.Time
	DueDate       *time.Time
	Notes         string
	Totals        domain.Totals
	ConvertedToID *int64
	CreatedAt     time.Time

	// Filled by the queries that join or aggregate; zero otherwise.
	CustomerName string
	Lines        []Line
	Paid         domain.Money
}

// Outstanding is what is still owed on the document.
func (d Document) Outstanding() domain.Money {
	return d.Totals.Total - d.Paid
}

// Payment is money received against an invoice.
type Payment struct {
	ID         int64
	DocumentID int64
	Amount     domain.Money
	Method     string
	Reference  string
	PaidOn     time.Time
	CreatedAt  time.Time
}

// Message is one entry in a customer's correspondence timeline.
type Message struct {
	ID         int64
	CustomerID int64
	DocumentID *int64
	AuthorID   int64
	Direction  domain.MessageDirection
	Subject    string
	Body       string
	CreatedAt  time.Time

	// Filled by the timeline query; zero otherwise.
	AuthorName string
}
