package pdf

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/tounsils/i-CRM/internal/docview"
)

// ErrUnsupportedScript is returned when a document contains text this renderer
// cannot draw correctly.
//
// It exists because the behaviour it replaced was to drop such text silently:
// the document rendered with an empty address block and went out addressed to
// nobody, while looking correct on every screen. Refusing loudly is worse for
// the one person who wanted a Japanese invoice and far better for everyone who
// would otherwise have sent a blank one.
var ErrUnsupportedScript = errors.New("pdf: the document contains characters this renderer cannot draw")

// runeRange is an inclusive range of code points.
type runeRange struct{ lo, hi rune }

// shapingRequired lists the scripts this renderer cannot lay out, whatever
// font is loaded.
//
// Two things make a script belong here. It may read right to left, and the
// renderer places runs of text left to right with no concept of direction. Or
// each letter's drawn form may depend on its neighbours, which needs a shaping
// engine this has none of. Arabic needs both.
//
// This is deliberately separate from whether the font has the glyphs, because
// the two failures are different. A missing glyph is fixed by supplying a
// better font. These are not: Liberation Sans carries Hebrew, which makes
// Hebrew the tempting mistake, since it would draw something and that
// something would be backwards.
//
// See issue #40 for the obstacle: doing this properly needs a shaping library,
// and the mature one is C, which would end CGO_ENABLED=0 and the single static
// binary with it.
var shapingRequired = []runeRange{
	{0x0590, 0x05FF}, // Hebrew
	{0x0600, 0x06FF}, // Arabic
	{0x0700, 0x074F}, // Syriac
	{0x0750, 0x077F}, // Arabic Supplement
	{0x0780, 0x07BF}, // Thaana
	{0x07C0, 0x07FF}, // NKo
	{0x0800, 0x085F}, // Samaritan and Mandaic
	{0x08A0, 0x08FF}, // Arabic Extended-A
	{0x0900, 0x0DFF}, // Devanagari through Sinhala: the Indic scripts
	{0x0E00, 0x0EFF}, // Thai and Lao
	{0x0F00, 0x0FFF}, // Tibetan
	{0x1000, 0x109F}, // Myanmar
	{0x1780, 0x17FF}, // Khmer
	{0x200E, 0x200F}, // left-to-right and right-to-left marks
	{0x202A, 0x202E}, // bidirectional embedding controls
	{0x2066, 0x2069}, // bidirectional isolates
	{0xFB1D, 0xFDFF}, // Hebrew and Arabic presentation forms
	{0xFE70, 0xFEFF}, // Arabic presentation forms B
}

func inRanges(r rune, ranges []runeRange) bool {
	for _, span := range ranges {
		if r >= span.lo && r <= span.hi {
			return true
		}
	}
	return false
}

// drawableWhitespace is the whitespace a document legitimately contains.
// Addresses and notes carry newlines, and tabs occasionally survive a paste.
func drawableWhitespace(r rune) bool {
	return r == '\n' || r == '\t'
}

// isControl matches the C0 and C1 control characters, which have no glyph and
// would silently draw nothing.
func isControl(r rune) bool {
	return r < 0x20 || (r >= 0x7F && r <= 0x9F)
}

// unsupportedRunes returns the distinct code points in a document that the
// given fonts and this renderer cannot draw, in order, so the message names
// the actual problem rather than saying something went wrong.
func unsupportedRunes(doc docview.Document, fonts Fonts) []rune {
	var found []rune
	seen := map[rune]bool{}

	for _, text := range doc.Strings() {
		for _, r := range text {
			if seen[r] || drawableWhitespace(r) {
				continue
			}
			seen[r] = true

			if isControl(r) || inRanges(r, shapingRequired) || !fonts.hasGlyph(r) {
				found = append(found, r)
			}
		}
	}

	slices.Sort(found)
	return found
}

// maxNamedRunes caps how many offending characters a message quotes. Six is
// enough to recognise the script; a full list of every character in a Japanese
// company name is not a more useful error.
const maxNamedRunes = 6

// describeRunes renders the offending characters for a message.
func describeRunes(runes []rune) string {
	quoted := make([]string, 0, maxNamedRunes)
	for i, r := range runes {
		if i == maxNamedRunes {
			break
		}
		quoted = append(quoted, fmt.Sprintf("%q", r))
	}
	out := strings.Join(quoted, ", ")
	if extra := len(runes) - maxNamedRunes; extra > 0 {
		out += fmt.Sprintf(" and %d more", extra)
	}
	return out
}
