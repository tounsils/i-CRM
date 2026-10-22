package pdf

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
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
	file, err := Render(sampleView(), testFonts(t))
	if err != nil {
		t.Fatalf("Render: %v", err)
	}

	if !bytes.HasPrefix(file, []byte("%PDF-")) {
		t.Errorf("file does not begin with the PDF header, starts with %q", first(file, 16))
	}
	if !bytes.Contains(file, []byte("%%EOF")) {
		t.Error("file has no end-of-file trailer")
	}
	// An embedded font subset puts the floor well above an empty document.
	if len(file) < 10000 {
		t.Errorf("file is %d bytes, too small to carry a page and a font subset", len(file))
	}
	// The font has to actually be embedded, or a reader without Liberation
	// installed falls back to something else and the metrics shift.
	if !bytes.Contains(file, []byte("FontFile2")) {
		t.Error("no embedded font programme in the file")
	}
	if !bytes.Contains(file, []byte("Identity-H")) {
		t.Error("the font is not CID-encoded, so text above Latin-1 cannot be addressed")
	}
	if !bytes.Contains(file, []byte("ToUnicode")) {
		t.Error("no ToUnicode map, so a reader cannot recover the text from the file")
	}
}

// TestRenderEmbedsGlyphsForTextAboveLatin1 is the regression test for the bug
// this renderer shipped with: text outside Latin-1 was silently dropped, so a
// document went out with an empty address block.
//
// The text itself cannot be searched for in the bytes. A CID-keyed font writes
// glyph indices into the content stream rather than characters, and the
// ToUnicode map that reverses that is a compressed stream. So the assertion is
// indirect but sound: a glyph that is drawn gets added to the embedded font
// subset, and a glyph that is dropped does not. Compared against an ASCII name
// of the same length, non-Latin-1 text must therefore make the file bigger.
//
// Reading the text back out properly needs a PDF parser, which is why the
// container end-to-end check runs the file through poppler.
func TestRenderEmbedsGlyphsForTextAboveLatin1(t *testing.T) {
	render := func(name string) int {
		t.Helper()
		view := sampleView()
		view.CustomerLines = []string{name}
		file, err := Render(view, testFonts(t))
		if err != nil {
			t.Fatalf("Render(%q): %v", name, err)
		}
		return len(file)
	}

	// Same number of characters, so only the glyph subset differs.
	baseline := render("Firma Lodz Zazolc")

	cases := []struct {
		name string
		text string
	}{
		{"Polish", "Firma Łódź Zażółć"},
		{"Cyrillic", "ООО Северный Ветер"},
		{"Greek", "Εταιρεία Βορρά ΑΕ"},
		{"Turkish", "Şirket İstanbul AŞ"},
		{"Vietnamese", "Công ty Thương mại"},
	}
	for _, c := range cases {
		size := render(c.text)
		if size <= baseline {
			t.Errorf("%s: %d bytes against an ASCII baseline of %d. No glyphs were added, "+
				"so the text was dropped rather than drawn", c.name, size, baseline)
		}
	}
}

func TestRenderAcceptsEveryScriptItClaimsToSupport(t *testing.T) {
	// If one of these starts failing, either the font changed or a range was
	// removed from the allowlist, and a customer who could be invoiced
	// yesterday cannot be today.
	for _, text := range []string{
		"Plain ASCII Ltd",
		"Café Léger Ützel",       // Latin-1
		"Firma Łódź Zażółć",      // Latin Extended-A
		"Şirket İstanbul",        // Turkish
		"ООО Северный Ветер",     // Cyrillic
		"Εταιρεία Βορρά",         // Greek
		"Công ty Thương mại",     // Latin Extended Additional
		"Dash – quote ‘x’ €1.00", // punctuation and currency
	} {
		view := sampleView()
		view.CustomerLines = []string{text}
		if _, err := Render(view, testFonts(t)); err != nil {
			t.Errorf("Render(%q) returned %v, but this script is supposed to be supported", text, err)
		}
	}
}

func TestRenderRefusesWhatItCannotDrawCorrectly(t *testing.T) {
	// Refusing is the fix. The previous behaviour drew the document with the
	// text missing and reported success, which is how an invoice went out
	// addressed to nobody.
	cases := []struct {
		name string
		text string
	}{
		{"Arabic", "شركة الشمال للتجارة"},
		{"Hebrew", "חברת הצפון"},
		{"Japanese", "北風商事株式会社"},
		{"Korean", "북풍상사"},
		{"Devanagari", "उत्तर व्यापार"},
		{"Thai", "บริษัท เหนือ"},
		{"emoji", "Northwind 🚀 Trading"},
	}
	for _, c := range cases {
		view := sampleView()
		view.CustomerLines = []string{c.text}

		file, err := Render(view, testFonts(t))
		if !errors.Is(err, ErrUnsupportedScript) {
			t.Errorf("%s: Render returned %v, want ErrUnsupportedScript", c.name, err)
			continue
		}
		if file != nil {
			t.Errorf("%s: a refused render still returned %d bytes", c.name, len(file))
		}
		// The message has to name the problem, or an operator cannot tell
		// which customer to fix.
		if !strings.Contains(err.Error(), "cannot draw") {
			t.Errorf("%s: unhelpful message %q", c.name, err)
		}
	}
}

func TestRefusalNamesTheOffendingCharacters(t *testing.T) {
	view := sampleView()
	view.Notes = "Delivered to 北風 warehouse"

	_, err := Render(view, testFonts(t))
	if !errors.Is(err, ErrUnsupportedScript) {
		t.Fatalf("Render returned %v, want ErrUnsupportedScript", err)
	}
	for _, want := range []string{"北", "風"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("the message %q does not name the character %q", err, want)
		}
	}
}

func TestRefusalChecksEveryPartOfTheDocument(t *testing.T) {
	// A field that is not checked is a field that renders blank. Each of these
	// puts the offending text somewhere different.
	base := sampleView()

	withKanji := func(mutate func(*docview.Document)) error {
		view := base
		mutate(&view)
		_, err := Render(view, testFonts(t))
		return err
	}

	cases := map[string]func(*docview.Document){
		"company name":    func(d *docview.Document) { d.CompanyName = "北風" },
		"company address": func(d *docview.Document) { d.CompanyAddress = []string{"北風"} },
		"customer":        func(d *docview.Document) { d.CustomerLines = []string{"北風"} },
		"notes":           func(d *docview.Document) { d.Notes = "北風" },
		"line description": func(d *docview.Document) {
			d.Lines = []docview.Line{{Description: "北風", Quantity: "1", UnitPrice: "0.00"}}
		},
		"meta value":   func(d *docview.Document) { d.Meta = []docview.Field{{Label: "Issued", Value: "北風"}} },
		"totals label": func(d *docview.Document) { d.Totals = []docview.Field{{Label: "北風", Value: "USD 0.00"}} },
		"payment method": func(d *docview.Document) {
			d.Payments = []docview.Payment{{Date: "21 Oct 2026", Method: "北風", Reference: "-", Amount: "1.00"}}
		},
	}
	for where, mutate := range cases {
		if err := withKanji(mutate); !errors.Is(err, ErrUnsupportedScript) {
			t.Errorf("unsupported text in the %s was not caught: err = %v", where, err)
		}
	}
}

func TestShapingRangesAreSaneAndOrdered(t *testing.T) {
	// An inverted or overlapping range is a silent bug: the lookup would still
	// work, so nothing would fail, but the table would no longer say what it
	// appears to say.
	for i, span := range shapingRequired {
		if span.lo > span.hi {
			t.Errorf("range %d is inverted: %#x to %#x", i, span.lo, span.hi)
		}
		if i > 0 && span.lo <= shapingRequired[i-1].hi {
			t.Errorf("range %d (%#x) overlaps or precedes range %d (ending %#x)",
				i, span.lo, i-1, shapingRequired[i-1].hi)
		}
	}
}

func TestEmbeddedFontCoverageIsAskedOfTheFontNotHardcoded(t *testing.T) {
	// The reason coverage is read from the font: a list of ranges in the
	// source can drift away from what the file actually contains, and the
	// failure is a document that renders blank.
	fonts := testFonts(t)

	for _, present := range []rune{'A', 'é', 'Ł', 'ż', 'Ş', 'О', 'Σ', '€', 'ế'} {
		if !fonts.hasGlyph(present) {
			t.Errorf("the embedded font reports no glyph for %q, but it should have one", present)
		}
	}
	for _, absent := range []rune{'北', '風', '한', 'ا'} {
		if fonts.hasGlyph(absent) {
			t.Errorf("the embedded font reports a glyph for %q, which it does not carry", absent)
		}
	}
}

func TestHebrewIsRefusedEvenThoughTheFontHasTheGlyphs(t *testing.T) {
	// The tempting mistake. Liberation Sans carries Hebrew, so a check based
	// only on glyph coverage would accept it and draw it left to right, which
	// is backwards. Direction is a renderer limitation, not a font one, and
	// the two are tracked separately for exactly this case.
	fonts := testFonts(t)

	// Asserted rather than skipped over. If the font stops carrying Hebrew
	// this test still passes for the wrong reason, so the premise is checked
	// explicitly and a change to it fails here and makes somebody think.
	if !fonts.hasGlyph('א') {
		t.Fatal("the premise of this test has changed: the embedded font no longer " +
			"carries Hebrew, so glyph coverage alone would now refuse it anyway")
	}

	view := sampleView()
	view.CustomerLines = []string{"חברת הצפון"}
	if _, err := Render(view, fonts); !errors.Is(err, ErrUnsupportedScript) {
		t.Fatalf("Render returned %v, want ErrUnsupportedScript: the glyphs exist but the "+
			"direction cannot be honoured", err)
	}
}

func TestLoadFallsBackToTheEmbeddedFont(t *testing.T) {
	fonts, err := Load("", "")
	if err != nil {
		t.Fatalf("Load with no paths: %v", err)
	}
	if !fonts.hasGlyph('Ł') {
		t.Error("the fallback font cannot draw Latin Extended")
	}
	if fonts.BoldIsRegular {
		t.Error("the embedded set has a real bold face, but BoldIsRegular is true")
	}
}

func TestLoadRejectsAFileThatIsNotAUsableFont(t *testing.T) {
	// The two formats people reach for first, a .ttc collection and a
	// PostScript-outline .otf, both fail inside the PDF writer with an
	// unhelpful "undefined font" at download time. Catching it at load means
	// the message arrives when the operator sets the variable.
	notAFont := filepath.Join(t.TempDir(), "notafont.ttf")
	if err := os.WriteFile(notAFont, []byte("this is not a font at all"), 0o600); err != nil {
		t.Fatalf("writing the fixture: %v", err)
	}

	_, err := Load(notAFont, "")
	if err == nil {
		t.Fatal("Load accepted a file that is not a font")
	}
	for _, want := range []string{"single-face TrueType", ".ttc", ".otf"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("the error does not mention %q, so it does not say how to fix it: %v", want, err)
		}
	}
}

func TestLoadReportsAMissingFile(t *testing.T) {
	if _, err := Load(filepath.Join(t.TempDir(), "absent.ttf"), ""); err == nil {
		t.Fatal("Load accepted a path that does not exist")
	}
}

func TestLoadUsesTheRegularFaceForBoldWhenNoneIsSupplied(t *testing.T) {
	// A deliberate degradation: an operator with one font file for a script
	// this binary cannot otherwise draw is better served by documents in a
	// single weight than by no documents.
	regular, err := fontFS.ReadFile(embeddedRegular)
	if err != nil {
		t.Fatalf("reading the embedded font: %v", err)
	}
	path := filepath.Join(t.TempDir(), "supplied.ttf")
	if err := os.WriteFile(path, regular, 0o600); err != nil {
		t.Fatalf("writing the fixture: %v", err)
	}

	fonts, err := Load(path, "")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !fonts.BoldIsRegular {
		t.Error("no bold face was supplied, but BoldIsRegular is false")
	}
	if fonts.Source != path {
		t.Errorf("Source = %q, want the supplied path", fonts.Source)
	}
	// And it still renders, which is the point of the fallback.
	if _, err := Render(sampleView(), fonts); err != nil {
		t.Fatalf("Render with a regular-only font: %v", err)
	}
}

func TestRenderHandlesADocumentWithNothingOptional(t *testing.T) {
	minimal := docview.Document{
		Kind:          "Proposal",
		Number:        "PRO-0001",
		CompanyName:   "Harbour Software",
		CustomerLines: []string{"Northwind Trading"},
		Meta:          []docview.Field{{Label: "Currency", Value: "USD"}},
		Lines: []docview.Line{{
			Description: "Discovery", Quantity: "1", UnitPrice: "0.00",
			Tax: "-", Net: "0.00", TaxAmount: "0.00",
		}},
		Totals: []docview.Field{{Label: "Total", Value: "USD 0.00", Emphasis: true}},
	}

	if _, err := Render(minimal, testFonts(t)); err != nil {
		t.Fatalf("Render: %v", err)
	}
}

func TestRenderWrapsALongDescriptionRatherThanOverrunning(t *testing.T) {
	view := sampleView()
	short, err := Render(view, testFonts(t))
	if err != nil {
		t.Fatalf("Render short: %v", err)
	}

	view.Lines[0].Description = strings.Repeat("A very long description of the work performed. ", 12)
	long, err := Render(view, testFonts(t))
	if err != nil {
		t.Fatalf("Render long: %v", err)
	}

	if len(long) <= len(short) {
		t.Errorf("a long description produced %d bytes against %d for a short one, "+
			"so it was probably truncated rather than wrapped", len(long), len(short))
	}
}

func TestRenderSurvivesManyLines(t *testing.T) {
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

	file, err := Render(view, testFonts(t))
	if err != nil {
		t.Fatalf("Render: %v", err)
	}
	if pages := bytes.Count(file, []byte("/Type /Page\n")); pages < 2 {
		t.Errorf("found %d page objects, want at least 2 for 81 lines", pages)
	}
}

func first(b []byte, n int) string {
	if len(b) < n {
		n = len(b)
	}
	return string(b[:n])
}

// testFonts is the embedded typeface, which is what a default deployment uses.
func testFonts(t *testing.T) Fonts {
	t.Helper()
	fonts, err := Embedded()
	if err != nil {
		t.Fatalf("Embedded: %v", err)
	}
	return fonts
}
