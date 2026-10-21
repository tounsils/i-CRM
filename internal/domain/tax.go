package domain

import (
	"errors"
	"fmt"
)

// Rate is a tax rate in basis points, so 2000 is 20% and 825 is 8.25%.
// Basis points express every rate in real-world use (a quarter of a percent is
// the finest granularity anyone levies) while keeping the rate an exact
// integer, which a percentage held as a float would not be.
type Rate int32

// BasisPointScale is the divisor that turns basis points into a proportion.
const BasisPointScale = 10000

// ErrRateOutOfRange is returned for a percentage below zero or above 100.
var ErrRateOutOfRange = errors.New("domain: a tax rate must be between 0% and 100%")

// ParseRate reads a percentage with at most two decimal places ("20", "8.25")
// into basis points.
//
// The conversion is exactly the one ParseMoney already performs, since a
// percentage with two decimals maps to basis points the same way an amount with
// two decimals maps to minor units. Reusing it means there is one decimal parser
// in this package rather than two that can drift apart.
func ParseRate(s string) (Rate, error) {
	parsed, err := ParseMoney(s)
	if err != nil {
		return 0, err
	}
	if parsed < 0 || parsed > BasisPointScale {
		return 0, fmt.Errorf("%w: got %q", ErrRateOutOfRange, s)
	}
	return Rate(parsed), nil
}

// Apply returns the tax owed on base at this rate, rounded half away from zero
// to the minor unit.
//
// Tax is computed per line and the rounded per-line amounts are summed, rather
// than applying the rate once to the document subtotal. The two results differ
// by up to one minor unit per line. Per-line is the rule here because the tax
// printed beside a line has to be the tax that was added to the total: a
// customer who recomputes the invoice by hand must land on the same figure, and
// an auditor reading one line must not find it inconsistent with the footer.
// TestTotalsAreInternallyConsistent pins this to stop a later change to
// ComputeTotals from quietly switching to subtotal-level tax.
func (r Rate) Apply(base Money) Money {
	return Money(divRoundHalfAway(int64(base)*int64(r), BasisPointScale))
}

// String renders the rate as a percentage with no trailing zeros, so 2000 is
// "20%" and 825 is "8.25%".
func (r Rate) String() string {
	v := int64(r)
	sign := ""
	if v < 0 {
		sign = "-"
		v = -v
	}
	whole, frac := v/100, v%100
	switch {
	case frac == 0:
		return fmt.Sprintf("%s%d%%", sign, whole)
	case frac%10 == 0:
		return fmt.Sprintf("%s%d.%d%%", sign, whole, frac/10)
	default:
		return fmt.Sprintf("%s%d.%02d%%", sign, whole, frac)
	}
}
