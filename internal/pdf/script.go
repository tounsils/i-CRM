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
// It exists because the previous behaviour was to drop such text silently: the
// document rendered with an empty address block and went out addressed to
// nobody, while looking correct on every screen. Refusing loudly is worse for
// the one person who wanted a Japanese invoice and far better for everyone who
// would otherwise have sent a blank one.
var ErrUnsupportedScript = errors.New("pdf: the document contains characters this renderer cannot draw")

// runeRange is an inclusive range of code points.
type runeRange struct{ lo, hi rune }

// renderable is what the embedded font can draw and this renderer can lay out
// correctly. It is an allowlist rather than a list of exclusions, because the
// set of scripts in the world is open and the set this can handle is not.
//
// Two separate requirements have to hold for a range to appear here. The font
// must carry the glyphs, and the script must read left to right without
// contextual shaping, because the renderer places runs of text and has no
// concept of either.
//
// Deliberately absent, and why:
//
//   - Arabic, Hebrew, Thaana, Syriac. Right to left, and Arabic additionally
//     needs each letter's form chosen by its neighbours. Liberation Sans has
//     no Arabic glyphs at all; it does have Hebrew, which makes Hebrew the
//     tempting mistake, since it would draw something, in the wrong order.
//   - CJK and Hangul. No font small enough to embed carries them. A font with
//     full CJK coverage is tens of megabytes.
//   - Devanagari, Thai, and the other shaping scripts, for the same reason as
//     Arabic.
var renderable = []runeRange{
	{0x0009, 0x000A}, // tab and newline, which appear in notes and addresses
	{0x0020, 0x007E}, // Basic Latin
	{0x00A0, 0x00FF}, // Latin-1 Supplement: accented French, German, Spanish
	{0x0100, 0x017F}, // Latin Extended-A: Polish, Czech, Turkish, Baltic
	{0x0180, 0x024F}, // Latin Extended-B
	{0x02B0, 0x02FF}, // spacing modifiers
	{0x0300, 0x036F}, // combining diacritics
	{0x0370, 0x03FF}, // Greek and Coptic
	{0x0400, 0x04FF}, // Cyrillic
	{0x0500, 0x052F}, // Cyrillic Supplement
	{0x1E00, 0x1EFF}, // Latin Extended Additional, including Vietnamese
	{0x2010, 0x2027}, // dashes, quotation marks, ellipsis
	{0x20A0, 0x20BF}, // currency symbols
	{0x2122, 0x2122}, // trade mark sign
}

// renderableRune reports whether this renderer can draw a code point.
func renderableRune(r rune) bool {
	for _, span := range renderable {
		if r >= span.lo && r <= span.hi {
			return true
		}
	}
	return false
}

// unsupportedRunes returns the distinct code points in a document that cannot
// be drawn, in order, so the message names the actual problem rather than
// saying something went wrong.
func unsupportedRunes(doc docview.Document) []rune {
	var found []rune
	for _, text := range doc.Strings() {
		for _, r := range text {
			if !renderableRune(r) && !slices.Contains(found, r) {
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
