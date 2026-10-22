// Package docview turns a stored document into the rows a human-facing
// rendering needs: labelled fields and already-formatted strings, ready to
// place on a page.
//
// It exists because a document is now rendered twice, as HTML for the print
// view and as a PDF for download, and those two must never disagree about what
// the document says. Both render from one of these, so which rows a totals
// block carries and how a figure is formatted are decided once. Styling can
// still differ between the two; the content cannot.
//
// Nothing here computes a total. The figures come from the document, which got
// them from domain.ComputeTotals, so this package cannot introduce a third
// opinion about the arithmetic.
package docview

import (
	"strings"

	"github.com/tounsils/i-CRM/internal/domain"
	"github.com/tounsils/i-CRM/internal/store"
)

// dateFormat is how every date on a document reads. One constant, so the PDF
// and the HTML cannot format the same date two ways.
const dateFormat = "2 Jan 2006"

// absent is what stands in for an optional value that was not supplied, so a
// table cell is never mysteriously blank.
const absent = "-"

// Field is a labelled value, used for the metadata and totals blocks.
type Field struct {
	Label string
	Value string

	// Emphasis marks the rows a reader's eye should land on: the document
	// total, and anything still outstanding.
	Emphasis bool
}

// Line is one priced row with every figure already formatted.
type Line struct {
	Description string
	Quantity    string
	UnitPrice   string
	Tax         string
	Net         string
	TaxAmount   string
}

// Payment is one payment already received.
type Payment struct {
	Date      string
	Method    string
	Reference string
	Amount    string
}

// Document is everything a rendering needs and nothing it does not.
type Document struct {
	Kind           string
	Number         string
	CompanyName    string
	CompanyAddress []string
	CustomerLines  []string
	Meta           []Field
	Lines          []Line
	Totals         []Field
	Payments       []Payment
	Notes          string
}

// Input is what Build needs, as a struct rather than five parameters.
type Input struct {
	Document       store.Document
	Customer       store.Customer
	Payments       []store.Payment
	CompanyName    string
	CompanyAddress string
}

// Build assembles the view of a document.
func Build(in Input) Document {
	d := in.Document

	out := Document{
		Kind:           d.Type.Label(),
		Number:         d.Number,
		CompanyName:    in.CompanyName,
		CompanyAddress: splitLines(in.CompanyAddress),
		CustomerLines:  customerLines(in.Customer),
		Notes:          strings.TrimSpace(d.Notes),
	}

	out.Meta = append(out.Meta, Field{Label: "Issued", Value: d.IssueDate.Format(dateFormat)})
	if d.DueDate != nil {
		out.Meta = append(out.Meta, Field{Label: "Due", Value: d.DueDate.Format(dateFormat)})
	}
	out.Meta = append(out.Meta, Field{Label: "Currency", Value: d.Currency})

	for _, line := range d.Lines {
		totals := line.Totals()
		out.Lines = append(out.Lines, Line{
			Description: line.Description,
			Quantity:    line.Quantity.String(),
			UnitPrice:   line.UnitPrice.String(),
			Tax:         taxLabel(line.TaxName, line.TaxRate),
			Net:         totals.Subtotal.String(),
			TaxAmount:   totals.Tax.String(),
		})
	}

	out.Totals = []Field{
		{Label: "Subtotal", Value: withCurrency(d.Currency, d.Totals.Subtotal)},
		{Label: "Tax", Value: withCurrency(d.Currency, d.Totals.Tax)},
		{Label: "Total", Value: withCurrency(d.Currency, d.Totals.Total), Emphasis: true},
	}

	// Paid and outstanding appear only once something has been received. On a
	// document with no payments they would both be noise, and "Outstanding"
	// equal to the total invites the reader to think a payment was missed.
	if len(in.Payments) > 0 {
		out.Totals = append(out.Totals,
			Field{Label: "Paid", Value: withCurrency(d.Currency, d.Paid)},
			Field{
				Label:    "Outstanding",
				Value:    withCurrency(d.Currency, d.Outstanding()),
				Emphasis: true,
			},
		)
	}

	for _, payment := range in.Payments {
		out.Payments = append(out.Payments, Payment{
			Date:      payment.PaidOn.Format(dateFormat),
			Method:    orAbsent(payment.Method),
			Reference: orAbsent(payment.Reference),
			Amount:    payment.Amount.String(),
		})
	}

	return out
}

// customerLines is the address block a document is addressed to, with the
// fields that were left empty omitted rather than printed as gaps.
func customerLines(c store.Customer) []string {
	lines := []string{c.Name}
	if c.ContactName != "" {
		lines = append(lines, c.ContactName)
	}
	lines = append(lines, splitLines(c.Address)...)
	if c.Email != "" {
		lines = append(lines, c.Email)
	}
	return lines
}

// splitLines turns a multi-line value into its lines, dropping blank ones so a
// trailing newline in a setting does not become a gap on the page.
func splitLines(s string) []string {
	var out []string
	for _, line := range strings.Split(s, "\n") {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}

func taxLabel(name string, rate domain.Rate) string {
	if name == "" {
		return absent
	}
	return name + " (" + rate.String() + ")"
}

func withCurrency(currency string, amount domain.Money) string {
	return currency + " " + amount.String()
}

func orAbsent(s string) string {
	if strings.TrimSpace(s) == "" {
		return absent
	}
	return s
}
