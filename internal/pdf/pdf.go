// Package pdf renders a document view as a PDF.
//
// It takes a docview.Document, the same value the HTML print view renders, so
// the figures in a downloaded file are the figures on the screen by
// construction rather than by two layouts being kept in step by hand.
//
// The library is pure Go, so CGO stays off and the binary stays
// self-contained. That constraint is the reason a headless browser was not used
// here, even though it would have given one layout instead of two.
package pdf

import (
	"bytes"
	"fmt"

	"github.com/go-pdf/fpdf"

	"github.com/tounsils/i-CRM/internal/docview"
)

// Page geometry in millimetres, for A4 portrait.
//
// The six column widths sum to exactly the 174mm of content width that A4
// leaves between 18mm margins. They are written as named constants rather than
// computed so that a change which no longer fits is a visible edit instead of a
// silently overflowing table.
const (
	margin      = 18.0
	contentWide = 174.0

	colDescription = 64.0
	colQuantity    = 16.0
	colUnitPrice   = 24.0
	colTax         = 26.0
	colNet         = 22.0
	colTaxAmount   = 22.0
)

// Typography.
const (
	fontFamily = "Liberation"

	sizeTitle = 20.0
	sizeName  = 12.0
	sizeBody  = 9.0
	sizeSmall = 8.0

	lineHeight = 5.0
	rowPadding = 1.5
)

// Render returns the document as a PDF file, drawn with the given fonts.
//
// It returns ErrUnsupportedScript, before drawing anything, if the document
// contains text these fonts or this renderer cannot draw correctly. Refusing is
// deliberate: the alternative, which this replaced, was a document that
// rendered with the text missing and no indication that anything was wrong.
func Render(doc docview.Document, fonts Fonts) ([]byte, error) {
	if bad := unsupportedRunes(doc, fonts); len(bad) > 0 {
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedScript, describeRunes(bad))
	}

	out := fpdf.New("P", "mm", "A4", "")

	// Registered as UTF-8 fonts, which is what makes the text above Latin-1
	// render rather than vanish. Both styles are needed: fpdf will not
	// synthesise a bold face, it reports an undefined font instead.
	out.AddUTF8FontFromBytes(fontFamily, "", fonts.regular)
	out.AddUTF8FontFromBytes(fontFamily, "B", fonts.bold)

	out.SetMargins(margin, margin, margin)
	out.SetAutoPageBreak(true, margin)

	out.SetFooterFunc(func() {
		out.SetY(-margin + 6)
		out.SetFont(fontFamily, "", sizeSmall)
		out.SetTextColor(110, 110, 110)
		footer := fmt.Sprintf("%s %s", doc.Kind, doc.Number)
		out.CellFormat(contentWide/2, lineHeight, footer, "", 0, "L", false, 0, "")
		out.CellFormat(contentWide/2, lineHeight,
			fmt.Sprintf("Page %d", out.PageNo()), "", 0, "R", false, 0, "")
	})

	out.AddPage()

	renderHeading(out, doc)
	renderParties(out, doc)
	renderLines(out, doc)
	renderTotals(out, doc)
	renderPayments(out, doc)
	renderNotes(out, doc)

	var buf bytes.Buffer
	if err := out.Output(&buf); err != nil {
		return nil, fmt.Errorf("pdf: writing %s %s: %w", doc.Kind, doc.Number, err)
	}
	// fpdf collects errors internally and keeps going, so a failure can be
	// reported after a successful-looking Output. Checked explicitly rather
	// than trusted.
	if err := out.Error(); err != nil {
		return nil, fmt.Errorf("pdf: rendering %s %s: %w", doc.Kind, doc.Number, err)
	}
	return buf.Bytes(), nil
}

// renderHeading is the company on the left and what this document is on the
// right.
func renderHeading(out *fpdf.Fpdf, doc docview.Document) {
	top := out.GetY()

	out.SetFont(fontFamily, "B", sizeName)
	out.SetTextColor(0, 0, 0)
	out.CellFormat(contentWide/2, lineHeight+1, doc.CompanyName, "", 1, "L", false, 0, "")

	out.SetFont(fontFamily, "", sizeBody)
	out.SetTextColor(90, 90, 90)
	for _, line := range doc.CompanyAddress {
		out.CellFormat(contentWide/2, lineHeight, line, "", 1, "L", false, 0, "")
	}
	leftBottom := out.GetY()

	out.SetXY(margin+contentWide/2, top)
	out.SetFont(fontFamily, "B", sizeTitle)
	out.SetTextColor(0, 0, 0)
	out.CellFormat(contentWide/2, 9, doc.Kind, "", 2, "R", false, 0, "")

	out.SetFont(fontFamily, "", sizeBody)
	out.SetTextColor(90, 90, 90)
	out.CellFormat(contentWide/2, lineHeight, doc.Number, "", 2, "R", false, 0, "")

	if rightBottom := out.GetY(); rightBottom > leftBottom {
		leftBottom = rightBottom
	}
	out.SetXY(margin, leftBottom+6)
}

// renderParties is who the document is for, beside its dates.
func renderParties(out *fpdf.Fpdf, doc docview.Document) {
	top := out.GetY()

	out.SetFont(fontFamily, "", sizeSmall)
	out.SetTextColor(120, 120, 120)
	out.CellFormat(contentWide/2, lineHeight, "TO", "", 1, "L", false, 0, "")

	out.SetTextColor(0, 0, 0)
	for i, line := range doc.CustomerLines {
		style := ""
		if i == 0 {
			style = "B"
		}
		out.SetFont(fontFamily, style, sizeBody)
		out.CellFormat(contentWide/2, lineHeight, line, "", 1, "L", false, 0, "")
	}
	leftBottom := out.GetY()

	out.SetXY(margin+contentWide/2, top)
	out.SetFont(fontFamily, "", sizeBody)
	for _, field := range doc.Meta {
		out.SetTextColor(120, 120, 120)
		out.CellFormat(contentWide/4, lineHeight, field.Label, "", 0, "R", false, 0, "")
		out.SetTextColor(0, 0, 0)
		out.CellFormat(contentWide/4, lineHeight, field.Value, "", 2, "R", false, 0, "")
		out.SetX(margin + contentWide/2)
	}

	if rightBottom := out.GetY(); rightBottom > leftBottom {
		leftBottom = rightBottom
	}
	out.SetXY(margin, leftBottom+8)
}

// renderLines is the priced table.
func renderLines(out *fpdf.Fpdf, doc docview.Document) {
	out.SetFont(fontFamily, "", sizeSmall)
	out.SetTextColor(120, 120, 120)
	header := []struct {
		width float64
		label string
		align string
	}{
		{colDescription, "DESCRIPTION", "L"},
		{colQuantity, "QTY", "R"},
		{colUnitPrice, "UNIT PRICE", "R"},
		{colTax, "TAX", "L"},
		{colNet, "NET", "R"},
		{colTaxAmount, "TAX", "R"},
	}
	for _, cell := range header {
		out.CellFormat(cell.width, lineHeight, cell.label, "", 0, cell.align, false, 0, "")
	}
	out.Ln(lineHeight)
	rule(out)

	out.SetTextColor(0, 0, 0)
	for _, line := range doc.Lines {
		out.SetFont(fontFamily, "", sizeBody)

		// A long description wraps rather than running under the next column,
		// and the row grows to fit it.
		wrapped := out.SplitLines([]byte(line.Description), colDescription-2)
		if len(wrapped) == 0 {
			wrapped = [][]byte{{}}
		}
		rowHeight := float64(len(wrapped))*lineHeight + rowPadding

		top := out.GetY()
		for i, part := range wrapped {
			out.SetXY(margin, top+float64(i)*lineHeight)
			out.CellFormat(colDescription, lineHeight, string(part), "", 0, "L", false, 0, "")
		}

		out.SetXY(margin+colDescription, top)
		for _, cell := range []struct {
			width float64
			value string
		}{
			{colQuantity, line.Quantity},
			{colUnitPrice, line.UnitPrice},
			{colTax, line.Tax},
			{colNet, line.Net},
			{colTaxAmount, line.TaxAmount},
		} {
			align := "R"
			if cell.width == colTax {
				align = "L"
			}
			out.CellFormat(cell.width, lineHeight, cell.value, "", 0, align, false, 0, "")
		}

		out.SetXY(margin, top+rowHeight)
		rule(out)
	}
}

// renderTotals is the right-aligned footer block.
func renderTotals(out *fpdf.Fpdf, doc docview.Document) {
	out.SetY(out.GetY() + 4)
	const labelWidth = 40.0
	const valueWidth = 44.0
	indent := margin + contentWide - labelWidth - valueWidth

	for _, field := range doc.Totals {
		style := ""
		if field.Emphasis {
			style = "B"
		}
		out.SetX(indent)
		out.SetFont(fontFamily, style, sizeBody)
		out.SetTextColor(90, 90, 90)
		out.CellFormat(labelWidth, lineHeight+1, field.Label, "", 0, "R", false, 0, "")
		out.SetTextColor(0, 0, 0)
		out.CellFormat(valueWidth, lineHeight+1, field.Value, "", 1, "R", false, 0, "")
	}
}

// renderPayments lists what has already been received, if anything has.
func renderPayments(out *fpdf.Fpdf, doc docview.Document) {
	if len(doc.Payments) == 0 {
		return
	}
	section(out, "PAYMENTS RECEIVED")

	widths := []float64{30, 50, 54, 40}
	out.SetFont(fontFamily, "", sizeBody)
	out.SetTextColor(0, 0, 0)
	for _, payment := range doc.Payments {
		cells := []string{payment.Date, payment.Method, payment.Reference, payment.Amount}
		for i, value := range cells {
			align := "L"
			if i == len(cells)-1 {
				align = "R"
			}
			out.CellFormat(widths[i], lineHeight, value, "", 0, align, false, 0, "")
		}
		out.Ln(lineHeight)
	}
}

// renderNotes is whatever the sender wanted the reader to see.
func renderNotes(out *fpdf.Fpdf, doc docview.Document) {
	if doc.Notes == "" {
		return
	}
	section(out, "NOTES")
	out.SetFont(fontFamily, "", sizeBody)
	out.SetTextColor(0, 0, 0)
	out.MultiCell(contentWide, lineHeight, doc.Notes, "", "L", false)
}

// section writes a small grey caption above a block.
func section(out *fpdf.Fpdf, label string) {
	out.SetY(out.GetY() + 8)
	out.SetX(margin)
	out.SetFont(fontFamily, "", sizeSmall)
	out.SetTextColor(120, 120, 120)
	out.CellFormat(contentWide, lineHeight, label, "", 1, "L", false, 0, "")
	out.SetTextColor(0, 0, 0)
}

// rule draws the hairline that separates table rows.
func rule(out *fpdf.Fpdf) {
	y := out.GetY()
	out.SetDrawColor(200, 200, 200)
	out.Line(margin, y, margin+contentWide, y)
	out.SetY(y + rowPadding)
}
