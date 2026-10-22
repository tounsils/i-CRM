package docview

import (
	"strings"
	"testing"
	"time"

	"github.com/tounsils/i-CRM/internal/domain"
	"github.com/tounsils/i-CRM/internal/store"
)

func sampleInput() Input {
	due := time.Date(2026, time.November, 20, 0, 0, 0, 0, time.UTC)
	return Input{
		Document: store.Document{
			ID:        1,
			Type:      domain.DocInvoice,
			Number:    "INV-0001",
			Status:    domain.StatusSent,
			Currency:  "USD",
			IssueDate: time.Date(2026, time.October, 21, 0, 0, 0, 0, time.UTC),
			DueDate:   &due,
			Notes:     "Thank you for your business.",
			Totals: domain.Totals{
				Subtotal: 1500000,
				Tax:      300000,
				Total:    1800000,
			},
			Lines: []store.Line{
				{Line: domain.Line{
					Description: "Implementation",
					Quantity:    10 * domain.QuantityScale,
					UnitPrice:   150000,
					TaxName:     "VAT",
					TaxRate:     2000,
				}},
				{Line: domain.Line{
					Description: "Zero rated book",
					Quantity:    2 * domain.QuantityScale,
					UnitPrice:   1250,
				}},
			},
		},
		Customer: store.Customer{
			Name:        "Northwind Trading",
			ContactName: "Ada Byron",
			Address:     "4 Dock Road\nPortsmouth PO1 3AX",
			Email:       "ap@northwind.example",
		},
		CompanyName:    "Harbour Software",
		CompanyAddress: "12 Harbour Way\nBristol BS1 4TR",
	}
}

func TestBuildCarriesTheDocumentsOwnFigures(t *testing.T) {
	// The point of this package: whatever a rendering shows must be the
	// document's own totals, not a recomputation that could differ.
	in := sampleInput()
	got := Build(in)

	if got.Kind != "Invoice" || got.Number != "INV-0001" {
		t.Fatalf("Kind/Number = %q/%q", got.Kind, got.Number)
	}

	wantTotals := map[string]string{
		"Subtotal": "USD " + in.Document.Totals.Subtotal.String(),
		"Tax":      "USD " + in.Document.Totals.Tax.String(),
		"Total":    "USD " + in.Document.Totals.Total.String(),
	}
	for _, field := range got.Totals {
		want, ok := wantTotals[field.Label]
		if !ok {
			t.Errorf("unexpected totals row %q", field.Label)
			continue
		}
		if field.Value != want {
			t.Errorf("%s = %q, want %q", field.Label, field.Value, want)
		}
		delete(wantTotals, field.Label)
	}
	if len(wantTotals) != 0 {
		t.Errorf("missing totals rows: %v", wantTotals)
	}
}

func TestBuildFormatsLinesOnce(t *testing.T) {
	got := Build(sampleInput())

	if len(got.Lines) != 2 {
		t.Fatalf("got %d lines, want 2", len(got.Lines))
	}

	first := got.Lines[0]
	if first.Quantity != "10" {
		t.Errorf("quantity = %q, want %q; trailing zeros are noise on a document",
			first.Quantity, "10")
	}
	if first.UnitPrice != "1500.00" {
		t.Errorf("unit price = %q, want 1500.00", first.UnitPrice)
	}
	if first.Tax != "VAT (20%)" {
		t.Errorf("tax = %q, want %q", first.Tax, "VAT (20%)")
	}
	if first.Net != "15000.00" || first.TaxAmount != "3000.00" {
		t.Errorf("net/tax = %q/%q, want 15000.00/3000.00", first.Net, first.TaxAmount)
	}

	// An untaxed line says so rather than showing a bare 0%.
	if got.Lines[1].Tax != absent {
		t.Errorf("untaxed line tax = %q, want %q", got.Lines[1].Tax, absent)
	}
}

func TestPaidAndOutstandingAppearOnlyWhenSomethingWasPaid(t *testing.T) {
	// On a document with no payments, an Outstanding row equal to the total
	// invites the reader to think a payment went missing.
	unpaid := Build(sampleInput())
	for _, field := range unpaid.Totals {
		if field.Label == "Paid" || field.Label == "Outstanding" {
			t.Fatalf("%q appears on a document with no payments", field.Label)
		}
	}
	if len(unpaid.Payments) != 0 {
		t.Fatalf("got %d payment rows, want none", len(unpaid.Payments))
	}

	in := sampleInput()
	in.Document.Paid = 800000
	in.Payments = []store.Payment{{
		Amount:    800000,
		Method:    "Bank transfer",
		PaidOn:    time.Date(2026, time.October, 22, 0, 0, 0, 0, time.UTC),
		Reference: "",
	}}
	paid := Build(in)

	labels := map[string]string{}
	for _, field := range paid.Totals {
		labels[field.Label] = field.Value
	}
	if labels["Paid"] != "USD 8000.00" {
		t.Errorf("Paid = %q, want USD 8000.00", labels["Paid"])
	}
	if labels["Outstanding"] != "USD 10000.00" {
		t.Errorf("Outstanding = %q, want USD 10000.00", labels["Outstanding"])
	}

	if len(paid.Payments) != 1 {
		t.Fatalf("got %d payment rows, want 1", len(paid.Payments))
	}
	if paid.Payments[0].Reference != absent {
		t.Errorf("a payment with no reference shows %q, want %q",
			paid.Payments[0].Reference, absent)
	}
	if paid.Payments[0].Date != "22 Oct 2026" {
		t.Errorf("payment date = %q, want 22 Oct 2026", paid.Payments[0].Date)
	}
}

func TestEmphasisMarksTheFiguresThatMatter(t *testing.T) {
	in := sampleInput()
	in.Document.Paid = 800000
	in.Payments = []store.Payment{{Amount: 800000, PaidOn: in.Document.IssueDate}}

	emphasised := map[string]bool{}
	for _, field := range Build(in).Totals {
		emphasised[field.Label] = field.Emphasis
	}
	for _, label := range []string{"Total", "Outstanding"} {
		if !emphasised[label] {
			t.Errorf("%q is not emphasised", label)
		}
	}
	for _, label := range []string{"Subtotal", "Tax", "Paid"} {
		if emphasised[label] {
			t.Errorf("%q is emphasised and should not be", label)
		}
	}
}

func TestAddressBlocksDropEmptyAndBlankLines(t *testing.T) {
	in := sampleInput()
	in.CompanyAddress = "12 Harbour Way\n\n  \nBristol BS1 4TR\n"
	in.Customer.ContactName = ""
	in.Customer.Email = ""

	got := Build(in)

	if len(got.CompanyAddress) != 2 {
		t.Fatalf("company address = %q, want the two real lines", got.CompanyAddress)
	}
	// The customer block is the name plus the two address lines, with the
	// contact and the email omitted because they were not supplied.
	want := []string{"Northwind Trading", "4 Dock Road", "Portsmouth PO1 3AX"}
	if strings.Join(got.CustomerLines, "|") != strings.Join(want, "|") {
		t.Fatalf("customer lines = %q, want %q", got.CustomerLines, want)
	}
}

func TestDueDateIsOmittedWhenThereIsNone(t *testing.T) {
	in := sampleInput()
	in.Document.DueDate = nil

	for _, field := range Build(in).Meta {
		if field.Label == "Due" {
			t.Fatal("a Due row appears on a document with no due date")
		}
	}

	withDue := Build(sampleInput())
	var found string
	for _, field := range withDue.Meta {
		if field.Label == "Due" {
			found = field.Value
		}
	}
	if found != "20 Nov 2026" {
		t.Fatalf("Due = %q, want 20 Nov 2026", found)
	}
}
