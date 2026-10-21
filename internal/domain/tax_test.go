package domain

import (
	"math/rand"
	"testing"
)

func TestRateApply(t *testing.T) {
	cases := []struct {
		name string
		rate Rate
		base Money
		want Money
	}{
		{"twenty percent of ten", 2000, 1000, 200},
		{"zero rated", 0, 123456, 0},
		{"eight and a quarter percent", 825, 10000, 825},
		{"rounds half away from zero", 2000, 7, 1}, // 1.4 minor units
		{"rounds up past half", 2000, 8, 2},        // 1.6 minor units
		{"exact half rounds away", 5000, 1, 1},     // 0.5 minor units
		{"negative base for a credit", 2000, -1000, -200},
		{"negative base rounds away", 2000, -7, -1},
	}
	for _, c := range cases {
		if got := c.rate.Apply(c.base); got != c.want {
			t.Errorf("%s: Rate(%d).Apply(%d) = %d, want %d", c.name, c.rate, c.base, got, c.want)
		}
	}
}

func TestRateApplyNeverExceedsTheBaseForSaneRates(t *testing.T) {
	// Invariant rather than example: for any rate at or below 100% the tax can
	// never come out larger than the amount it is levied on. A sign error or a
	// misplaced scale constant breaks this immediately.
	r := rand.New(rand.NewSource(3))
	for i := 0; i < 5000; i++ {
		base := Money(r.Int63n(10_000_00))
		rate := Rate(r.Int31n(BasisPointScale + 1))
		if tax := rate.Apply(base); tax > base {
			t.Fatalf("rate %s on base %s produced tax %s, which exceeds the base", rate, base, tax)
		}
	}
}

func TestRateString(t *testing.T) {
	cases := []struct {
		rate Rate
		want string
	}{
		{0, "0%"},
		{2000, "20%"},
		{825, "8.25%"},
		{750, "7.5%"},
		{10000, "100%"},
		{5, "0.05%"},
		{-2000, "-20%"},
	}
	for _, c := range cases {
		if got := c.rate.String(); got != c.want {
			t.Errorf("Rate(%d).String() = %q, want %q", c.rate, got, c.want)
		}
	}
}
