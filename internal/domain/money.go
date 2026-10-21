// Package domain holds the CRM's business rules: money arithmetic, tax
// computation and the document lifecycle. Nothing in here touches a database,
// an HTTP request or a template, so the rules that cost money when they are
// wrong can be tested without standing anything up.
package domain

import (
	"errors"
	"fmt"
	"strings"
)

// Money is an amount in the minor unit of its currency, so cents for USD and
// EUR. Invoice totals are legal figures and are never float64: binary floating
// point cannot hold 0.10 exactly, and the error compounds the moment a rate
// and a quantity are applied on top of it.
type Money int64

// Quantity is a line quantity in thousandths of a unit, so 1.5 hours is 1500.
// Three decimal places covers hourly billing and part-unit goods, which is the
// granularity invoicing actually needs, without reintroducing a float.
type Quantity int64

// QuantityScale turns a Quantity back into whole units.
const QuantityScale = 1000

// maxWholeDigits bounds the integer part accepted by ParseMoney. int64 minor
// units overflow above 9.2e16, so 15 digits of major units (1e17 minor) cannot
// be reached and the multiply in MulQuantity still has headroom.
const maxWholeDigits = 15

var ErrMalformedAmount = errors.New("domain: malformed amount")

// ParseMoney reads a decimal amount with at most two fractional digits, as a
// user types it into a line-item field, into minor units.
func ParseMoney(s string) (Money, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("%w: empty", ErrMalformedAmount)
	}

	negative := false
	switch s[0] {
	case '+':
		s = s[1:]
	case '-':
		negative = true
		s = s[1:]
	}

	whole, frac := s, ""
	if i := strings.IndexByte(s, '.'); i >= 0 {
		whole, frac = s[:i], s[i+1:]
	}
	if whole == "" && frac == "" {
		return 0, fmt.Errorf("%w: %q has no digits", ErrMalformedAmount, s)
	}
	if len(whole) > maxWholeDigits {
		return 0, fmt.Errorf("%w: %q is too large", ErrMalformedAmount, s)
	}
	if len(frac) > 2 {
		return 0, fmt.Errorf("%w: %q has more than two decimal places", ErrMalformedAmount, s)
	}
	for len(frac) < 2 {
		frac += "0"
	}

	units, err := parseDigits(whole)
	if err != nil {
		return 0, fmt.Errorf("%w: %q", ErrMalformedAmount, s)
	}
	cents, err := parseDigits(frac)
	if err != nil {
		return 0, fmt.Errorf("%w: %q", ErrMalformedAmount, s)
	}

	total := units*100 + cents
	if negative {
		total = -total
	}
	return Money(total), nil
}

func parseDigits(s string) (int64, error) {
	var n int64
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, ErrMalformedAmount
		}
		n = n*10 + int64(c-'0')
	}
	return n, nil
}

// String renders the amount with two decimal places and no currency symbol.
// The currency lives on the document, not on the amount, because a CRM holds
// documents in more than one currency and the symbol is a display concern.
func (m Money) String() string {
	v := int64(m)
	sign := ""
	if v < 0 {
		sign = "-"
		v = -v
	}
	return fmt.Sprintf("%s%d.%02d", sign, v/100, v%100)
}

// ParseQuantity reads a decimal quantity with at most three fractional digits,
// as typed into a line's quantity field.
func ParseQuantity(s string) (Quantity, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("%w: empty quantity", ErrMalformedAmount)
	}

	negative := false
	switch s[0] {
	case '+':
		s = s[1:]
	case '-':
		negative = true
		s = s[1:]
	}

	whole, frac := s, ""
	if i := strings.IndexByte(s, '.'); i >= 0 {
		whole, frac = s[:i], s[i+1:]
	}
	if whole == "" && frac == "" {
		return 0, fmt.Errorf("%w: %q has no digits", ErrMalformedAmount, s)
	}
	if len(whole) > maxWholeDigits {
		return 0, fmt.Errorf("%w: %q is too large", ErrMalformedAmount, s)
	}
	if len(frac) > 3 {
		return 0, fmt.Errorf("%w: %q has more than three decimal places", ErrMalformedAmount, s)
	}
	for len(frac) < 3 {
		frac += "0"
	}

	units, err := parseDigits(whole)
	if err != nil {
		return 0, fmt.Errorf("%w: %q", ErrMalformedAmount, s)
	}
	thousandths, err := parseDigits(frac)
	if err != nil {
		return 0, fmt.Errorf("%w: %q", ErrMalformedAmount, s)
	}

	total := units*QuantityScale + thousandths
	if negative {
		total = -total
	}
	return Quantity(total), nil
}

// String renders a quantity without trailing zeros, so 1500 is "1.5" and 2000
// is "2". A quantity column full of "2.000" is noise on an invoice.
func (q Quantity) String() string {
	v := int64(q)
	sign := ""
	if v < 0 {
		sign = "-"
		v = -v
	}
	whole, frac := v/QuantityScale, v%QuantityScale
	if frac == 0 {
		return fmt.Sprintf("%s%d", sign, whole)
	}
	text := fmt.Sprintf("%s%d.%03d", sign, whole, frac)
	return strings.TrimRight(text, "0")
}

// MulQuantity multiplies a unit price by a quantity and rounds half away from
// zero to the nearest minor unit. Half away from zero is what a person doing
// this on paper applies, and unlike banker's rounding it is symmetric, so a
// credit note for a returned line is the exact negative of the line.
func (m Money) MulQuantity(q Quantity) Money {
	return Money(divRoundHalfAway(int64(m)*int64(q), QuantityScale))
}

// divRoundHalfAway divides num by den, rounding a remainder of exactly half
// away from zero. den must be positive.
func divRoundHalfAway(num, den int64) int64 {
	if num < 0 {
		return -((-num*2 + den) / (den * 2))
	}
	return (num*2 + den) / (den * 2)
}
