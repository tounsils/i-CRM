package pdf

import (
	"embed"
	"fmt"
	"os"

	"golang.org/x/image/font/sfnt"
)

// Liberation Sans, embedded so the binary carries everything it needs to draw
// a document. The core PDF fonts every reader already has are Latin-1 only,
// which silently dropped any name outside it.
//
// Liberation is metric-compatible with Helvetica, so the column widths in
// pdf.go did not change when the font did.
//
// The licence is beside the fonts and must stay there: OFL 1.1 requires it to
// accompany the font software, and the Reserved Font Names in it are why these
// files must not be renamed or modified in place.
//
//go:embed fonts/LiberationSans-Regular.ttf fonts/LiberationSans-Bold.ttf
var fontFS embed.FS

const (
	embeddedRegular = "fonts/LiberationSans-Regular.ttf"
	embeddedBold    = "fonts/LiberationSans-Bold.ttf"
)

// Fonts is the typeface documents are drawn with, together with the parsed
// form used to ask which characters it can actually draw.
//
// Coverage is asked of the font rather than hardcoded. An operator who supplies
// a font with Chinese glyphs gets Chinese documents without this package being
// changed, and a list of ranges in the source cannot drift away from what the
// font really contains.
type Fonts struct {
	regular []byte
	bold    []byte

	// parsed is the regular face. Coverage is taken from it alone: a bold face
	// missing a glyph the regular one has would be an odd font, and checking
	// both would refuse documents that render fine in the body text.
	parsed *sfnt.Font

	// Source names where the fonts came from, for the startup log.
	Source string

	// BoldIsRegular is true when no bold face was supplied and the regular one
	// stands in, so bold text renders at normal weight rather than failing.
	BoldIsRegular bool
}

// Embedded returns the fonts compiled into the binary.
func Embedded() (Fonts, error) {
	regular, err := fontFS.ReadFile(embeddedRegular)
	if err != nil {
		return Fonts{}, fmt.Errorf("pdf: reading the embedded regular font: %w", err)
	}
	bold, err := fontFS.ReadFile(embeddedBold)
	if err != nil {
		return Fonts{}, fmt.Errorf("pdf: reading the embedded bold font: %w", err)
	}
	return newFonts(regular, bold, "embedded Liberation Sans", false)
}

// Load reads a typeface from disk, falling back to the embedded one when no
// path is given.
//
// boldPath may be empty, in which case the regular face is used for bold text.
// That is a deliberate degradation rather than an error: an operator who has
// one font file for a script this binary cannot otherwise draw is better served
// by documents that render in a single weight than by no documents.
func Load(regularPath, boldPath string) (Fonts, error) {
	if regularPath == "" {
		return Embedded()
	}

	regular, err := os.ReadFile(regularPath)
	if err != nil {
		return Fonts{}, fmt.Errorf("pdf: reading the font at %s: %w", regularPath, err)
	}

	bold, boldIsRegular := regular, true
	if boldPath != "" {
		bold, err = os.ReadFile(boldPath)
		if err != nil {
			return Fonts{}, fmt.Errorf("pdf: reading the bold font at %s: %w", boldPath, err)
		}
		boldIsRegular = false
	}

	return newFonts(regular, bold, regularPath, boldIsRegular)
}

func newFonts(regular, bold []byte, source string, boldIsRegular bool) (Fonts, error) {
	parsed, err := sfnt.Parse(regular)
	if err != nil {
		// The two formats people reach for first both fail here, so say so.
		// A .ttc holds several faces and the PDF writer cannot address one of
		// them; an .otf usually carries PostScript outlines rather than the
		// TrueType ones it can read. Both produce an unhelpful "undefined
		// font" much later, at the moment somebody tries to download an
		// invoice, which is the wrong time to find out.
		return Fonts{}, fmt.Errorf(
			"pdf: %s is not a usable font: %w. It must be a single-face TrueType .ttf "+
				"with glyf outlines; a .ttc collection or a .otf with PostScript outlines will not work",
			source, err)
	}

	return Fonts{
		regular:       regular,
		bold:          bold,
		parsed:        parsed,
		Source:        source,
		BoldIsRegular: boldIsRegular,
	}, nil
}

// hasGlyph reports whether the font can draw a code point.
func (f Fonts) hasGlyph(r rune) bool {
	if f.parsed == nil {
		return false
	}
	// A fresh buffer per call keeps this safe to use from concurrent requests.
	// sfnt.Buffer is explicitly not safe to share, and the cost is a small
	// allocation against a request that is about to write a PDF.
	var buf sfnt.Buffer
	index, err := f.parsed.GlyphIndex(&buf, r)
	return err == nil && index != 0
}
