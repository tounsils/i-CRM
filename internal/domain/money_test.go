package domain

import (
	"errors"
	"math/rand"
	"testing"
)

func TestParseMoneyAcceptsWhatUsersType(t *testing.T) {
	cases := []struct {
		in   string
		want Money
	}{
		{"0", 0},
		{"1", 100},
		{"1.5", 150},
		{"1.05", 105},
		{"1234.56", 123456},
		{".5", 50},
		{"-8", -800},
		{"-0.01", -1},
		{"+3.20", 320},
		{"  12.00  ", 1200},
	}
	for _, c := range cases {
		got, err := ParseMoney(c.in)
		if err != nil {
			t.Errorf("ParseMoney(%q) returned error %v", c.in, err)
			continue
		}
		if got != c.want {
			t.Errorf("ParseMoney(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestParseMoneyRejectsWhatWouldCorruptATotal(t *testing.T) {
	// Each of these has reached a billing system in the wild. Three decimal
	// places is the dangerous one: silently truncating 1.005 loses half a cent
	// per line with no error anywhere.
	cases := []string{"", "   ", "abc", "1.234", "1.2.3", "1,50", "$5", "1e3", "--1", "1234567890123456"}
	for _, in := range cases {
		if got, err := ParseMoney(in); err == nil {
			t.Errorf("ParseMoney(%q) = %d, want an error", in, got)
		} else if !errors.Is(err, ErrMalformedAmount) {
			t.Errorf("ParseMoney(%q) error = %v, want ErrMalformedAmount", in, err)
		}
	}
}

func TestMoneyStringRoundTripsThroughParse(t *testing.T) {
	// The invariant: rendering an amount and reading it back must be lossless,
	// because that round trip happens on every edit of a saved document.
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 2000; i++ {
		want := Money(r.Int63n(2_000_000_00) - 1_000_000_00)
		got, err := ParseMoney(want.String())
		if err != nil {
			t.Fatalf("ParseMoney(%q) returned error %v", want.String(), err)
		}
		if got != want {
			t.Fatalf("round trip of %d rendered %q and parsed back as %d", want, want.String(), got)
		}
	}
}

func TestMulQuantityRoundsHalfAwayFromZero(t *testing.T) {
	cases := []struct {
		name  string
		price Money
		qty   Quantity
		want  Money
	}{
		{"whole units", 1000, 3 * QuantityScale, 3000},
		{"one and a half hours", 1000, 1500, 1500},
		{"zero quantity", 1000, 0, 0},
		{"exact half rounds up", 1, 500, 1},
		{"exact half rounds away when negative", -1, 500, -1},
		{"below half rounds down", 1, 499, 0},
		{"above half rounds up", 1, 501, 1},
		{"third of a cent", 1, 333, 0},
		{"negative price", -1000, 1500, -1500},
		{"negative quantity", 1000, -1500, -1500},
	}
	for _, c := range cases {
		if got := c.price.MulQuantity(c.qty); got != c.want {
			t.Errorf("%s: Money(%d).MulQuantity(%d) = %d, want %d", c.name, c.price, c.qty, got, c.want)
		}
	}
}

func TestMulQuantityIsSymmetricAboutZero(t *testing.T) {
	// The invariant that half-away-from-zero buys us: a credit note for a
	// returned line is the exact negative of the line. Banker's rounding would
	// break this and leave a cent stranded on roughly half of all credits.
	r := rand.New(rand.NewSource(2))
	for i := 0; i < 5000; i++ {
		price := Money(r.Int63n(1_000_00))
		qty := Quantity(r.Int63n(10 * QuantityScale))
		if positive, negative := price.MulQuantity(qty), (-price).MulQuantity(qty); positive != -negative {
			t.Fatalf("price %d qty %d: %d is not the negative of %d", price, qty, positive, negative)
		}
	}
}
