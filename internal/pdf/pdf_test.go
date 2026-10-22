package pdf

import (
	"bytes"
	"strings"
	"testing"

	"github.com/tounsils/i-CRM/internal/docview"
)

func sampleView() docview.Document {
	return docview.Document{
		Kind:           "Invoice",
		Number:         "INV-0001",
		CompanyName:    "Harbour Software",
		CompanyAddress: []string{"12 Harbour Way", "Bristol BS1 4TR"},
		CustomerLines:  []string{"Northwind Trading", "Ada Byron", "4 Dock Road"},
		Meta: []docview.Field{
			{Label: "Issued", Value: "21 Oct 2026"},
			{Label: "Due", Value: "20 Nov 2026"},
			{Label: "Currency", Value: "USD"},
		},
		Lines: []docview.Line{
			{
				Description: "Implementation",
				Quantity:    "10",
				UnitPrice:   "1500.00",
				Tax:         "VAT (20%)",
				Net:         "15000.00",
				TaxAmount:   "3000.00",
			},
		},
		Totals: []docview.Field{
			{Label: "Subtotal", Value: "USD 15000.00"},
			{Label: "Tax", Value: "USD 3000.00"},
			{Label: "Total", Value: "USD 18000.00", Emphasis: true},
		},
		Notes: "Thank you for your business.",
	}
}

func TestRenderProducesAValidPDF(t *testing.T) {
	file, err := Render(sampleView())
	if err != nil {
		t.Fatalf("Render: %v", err)
	}

	// The two markers every reader looks for. Without the header a browser
	// will not open it; without the trailer it is truncated.
	if !bytes.HasPrefix(file, []byte("%PDF-")) {
		t.Errorf("file does not begin with the PDF header, starts with %q", first(file, 16))
	}
	if !bytes.Contains(file, []byte("%%EOF")) {
		t.Error("file has no end-of-file trailer")
	}
	if len(file) < 1000 {
		t.Errorf("file is %d bytes, which is too small to be a rendered page", len(file))
	}
}

func TestRenderWritesTheFiguresIntoTheFile(t *testing.T) {
	// Compression is off precisely so this can be asserted: the figures a
	// customer reads must actually be in the bytes, not merely passed to a
	// library that was trusted to draw them.
	file, err := Render(sampleView())
	if err != nil {
		t.Fatalf("Render: %v", err)
	}

	for _, want := range []string{
		"INV-0001", "Invoice",
		"Harbour Software", "Bristol BS1 4TR",
		"Northwind Trading", "Ada Byron",
		"Implementation", "VAT (20%)",
		"15000.00", "3000.00", "18000.00",
		"Thank you for your business.",
		"21 Oct 2026", "20 Nov 2026",
	} {
		if !bytes.Contains(file, []byte(pdfLiteral(want))) {
			t.Errorf("the rendered file does not contain %q", want)
		}
	}
}

// pdfLiteral escapes a string the way a PDF string literal holds it.
//
// Inside a content stream, text is written as `(…) Tj`, so a bracket in the
// text itself has to be escaped or it would close the literal early. A tax
// label like "VAT (20%)" is therefore stored as `VAT \(20%\)`, which a naive
// byte search for the original would miss and wrongly report as missing text.
func pdfLiteral(s string) string {
	replacer := strings.NewReplacer(`\`, `\\`, `(`, `\(`, `)`, `\)`)
	return replacer.Replace(s)
}

func TestRenderHandlesADocumentWithNothingOptional(t *testing.T) {
	// No address, no notes, no payments, one line. Every optional block is
	// absent, which is the shape most likely to index past the end of a slice.
	minimal := docview.Document{
		Kind:          "Proposal",
		Number:        "PRO-0001",
		CompanyName:   "Harbour Software",
		CustomerLines: []string{"Northwind Trading"},
		Meta:          []docview.Field{{Label: "Currency", Value: "USD"}},
		Lines:         []docview.Line{{Description: "Discovery", Quantity: "1", UnitPrice: "0.00", Tax: "-", Net: "0.00", TaxAmount: "0.00"}},
		Totals:        []docview.Field{{Label: "Total", Value: "USD 0.00", Emphasis: true}},
	}

	file, err := Render(minimal)
	if err != nil {
		t.Fatalf("Render: %v", err)
	}
	if !bytes.Contains(file, []byte("PRO-0001")) {
		t.Error("the minimal document does not carry its own number")
	}
}

func TestRenderWrapsALongDescriptionRatherThanOverrunning(t *testing.T) {
	// A description wider than its column must grow the row. The check is that
	// the file gains pages or length rather than that nothing visibly breaks,
	// which a byte comparison cannot see.
	view := sampleView()
	short, err := Render(view)
	if err != nil {
		t.Fatalf("Render short: %v", err)
	}

	view.Lines[0].Description = strings.Repeat("A very long description of the work performed. ", 12)
	long, err := Render(view)
	if err != nil {
		t.Fatalf("Render long: %v", err)
	}

	if len(long) <= len(short) {
		t.Errorf("a long description produced %d bytes against %d for a short one, "+
			"so it was probably truncated rather than wrapped", len(long), len(short))
	}
}

func TestRenderSurvivesManyLines(t *testing.T) {
	// Enough lines to force a page break, which is where an automatic break
	// interacting with manually positioned cells tends to go wrong.
	view := sampleView()
	for i := 0; i < 80; i++ {
		view.Lines = append(view.Lines, docview.Line{
			Description: "Additional work item",
			Quantity:    "1",
			UnitPrice:   "100.00",
			Tax:         "VAT (20%)",
			Net:         "100.00",
			TaxAmount:   "20.00",
		})
	}

	file, err := Render(view)
	if err != nil {
		t.Fatalf("Render: %v", err)
	}
	// More than one page object means the break happened rather than the rows
	// running off the bottom of a single page.
	if pages := bytes.Count(file, []byte("/Type /Page\n")); pages < 2 {
		t.Errorf("found %d page objects, want at least 2 for 81 lines", pages)
	}
	if !bytes.Contains(file, []byte("Page 2")) {
		t.Error("the footer does not show a second page")
	}
}

func first(b []byte, n int) string {
	if len(b) < n {
		n = len(b)
	}
	return string(b[:n])
}
