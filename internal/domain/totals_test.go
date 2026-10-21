package domain

import (
	"math/rand"
	"testing"
)

func TestComputeTotalsOnAnOrdinaryInvoice(t *testing.T) {
	lines := []Line{
		{Description: "Consulting", Quantity: 10 * QuantityScale, UnitPrice: 15000, TaxName: "VAT", TaxRate: 2000},
		{Description: "Hosting", Quantity: 1 * QuantityScale, UnitPrice: 4999, TaxName: "VAT", TaxRate: 2000},
		{Description: "Zero rated book", Quantity: 2 * QuantityScale, UnitPrice: 1250, TaxName: "None", TaxRate: 0},
	}

	got := ComputeTotals(lines)

	// 150000 + 4999 + 2500 subtotal; tax is 30000 + 1000 (999.8 rounded away
	// from zero) + 0.
	want := Totals{Subtotal: 157499, Tax: 31000, Total: 188499}
	if got != want {
		t.Fatalf("ComputeTotals = %+v, want %+v", got, want)
	}
}

func TestTotalsAreInternallyConsistent(t *testing.T) {
	// Two invariants over arbitrary documents, which together are the whole
	// contract of the footer:
	//
	//   1. Total is exactly Subtotal + Tax, so the figure a customer pays is
	//      the figure the two lines above it add up to.
	//   2. Subtotal is exactly the sum of the per-line subtotals, so no line is
	//      ever lost or double counted regardless of quantity or price.
	//
	// Asserting these rather than specific numbers is what catches a change to
	// ComputeTotals that happens to keep one worked example passing.
	r := rand.New(rand.NewSource(4))
	for i := 0; i < 2000; i++ {
		lines := make([]Line, r.Intn(12))
		for j := range lines {
			lines[j] = Line{
				Description: "line",
				Quantity:    Quantity(r.Int63n(20 * QuantityScale)),
				UnitPrice:   Money(r.Int63n(500_00)),
				TaxRate:     Rate(r.Int31n(BasisPointScale)),
			}
		}

		totals := ComputeTotals(lines)

		if totals.Total != totals.Subtotal+totals.Tax {
			t.Fatalf("%d lines: total %s is not subtotal %s plus tax %s",
				len(lines), totals.Total, totals.Subtotal, totals.Tax)
		}

		var sumOfLines Money
		for _, l := range lines {
			sumOfLines += l.Totals().Subtotal
		}
		if totals.Subtotal != sumOfLines {
			t.Fatalf("%d lines: subtotal %s but lines sum to %s", len(lines), totals.Subtotal, sumOfLines)
		}
	}
}

func TestTaxIsRoundedPerLineNotOnTheSubtotal(t *testing.T) {
	// The documented rounding decision, pinned so it cannot drift. Three lines
	// of 0.07 at 20%: each line's tax is 1.4 minor units which rounds to 1, so
	// the document tax is 3. Taxing the 0.21 subtotal once would give 4.2,
	// rounding to 4. The one-unit gap is the whole reason the rule is written
	// down in Rate.Apply, and this test fails if someone "simplifies"
	// ComputeTotals into a single subtotal-level multiply.
	lines := []Line{
		{Quantity: QuantityScale, UnitPrice: 7, TaxRate: 2000},
		{Quantity: QuantityScale, UnitPrice: 7, TaxRate: 2000},
		{Quantity: QuantityScale, UnitPrice: 7, TaxRate: 2000},
	}

	totals := ComputeTotals(lines)

	if totals.Subtotal != 21 {
		t.Fatalf("subtotal = %d, want 21", totals.Subtotal)
	}
	if totals.Tax != 3 {
		t.Fatalf("tax = %d, want 3 from per-line rounding; 4 means it was taken on the subtotal", totals.Tax)
	}
	if onSubtotal := Rate(2000).Apply(totals.Subtotal); onSubtotal != 4 {
		t.Fatalf("subtotal-level tax = %d, want 4; the premise of this test has changed", onSubtotal)
	}
}

func TestEmptyDocumentHasZeroTotals(t *testing.T) {
	if got := ComputeTotals(nil); got != (Totals{}) {
		t.Fatalf("ComputeTotals(nil) = %+v, want all zeroes", got)
	}
}
