package domain

import (
	"testing"
	"time"
)

func date(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

func ptr(t time.Time) *time.Time { return &t }

func TestIsOverdue(t *testing.T) {
	due := date(2026, time.October, 21)

	cases := []struct {
		name    string
		docType DocType
		status  DocStatus
		dueDate *time.Time
		asOf    time.Time
		want    bool
	}{
		// The boundary, which is the whole reason this is not a naive After().
		// Payment is due by the close of the due date.
		{"the day before", DocInvoice, StatusSent, ptr(due), date(2026, time.October, 20), false},
		{"on the due date at midnight", DocInvoice, StatusSent, ptr(due), due, false},
		{
			"late on the due date", DocInvoice, StatusSent, ptr(due),
			time.Date(2026, time.October, 21, 23, 59, 59, 0, time.UTC), false,
		},
		{"the day after", DocInvoice, StatusSent, ptr(due), date(2026, time.October, 22), true},
		{"months later", DocInvoice, StatusSent, ptr(due), date(2027, time.February, 1), true},

		// Part paid is still money owed.
		{"partially paid and late", DocInvoice, StatusPartiallyPaid, ptr(due), date(2026, time.November, 1), true},

		// Settled, withdrawn or never issued is never overdue, however old.
		{"paid", DocInvoice, StatusPaid, ptr(due), date(2027, time.January, 1), false},
		{"void", DocInvoice, StatusVoid, ptr(due), date(2027, time.January, 1), false},
		{"draft", DocInvoice, StatusDraft, ptr(due), date(2027, time.January, 1), false},

		// Only an invoice is money owed. An accepted estimate is not a debt.
		{"estimate", DocEstimate, StatusAccepted, ptr(due), date(2027, time.January, 1), false},
		{"proposal", DocProposal, StatusSent, ptr(due), date(2027, time.January, 1), false},

		// No due date means nothing to be late for.
		{"no due date", DocInvoice, StatusSent, nil, date(2027, time.January, 1), false},
	}

	for _, c := range cases {
		got := IsOverdue(c.docType, c.status, c.dueDate, c.asOf)
		if got != c.want {
			t.Errorf("%s: IsOverdue(%s, %s, %v, %s) = %v, want %v",
				c.name, c.docType, c.status, c.dueDate, c.asOf.Format(time.RFC3339), got, c.want)
		}
	}
}

func TestIsOverdueIgnoresTheTimeOfDayAndZone(t *testing.T) {
	// An invoice must not become overdue because the server happens to be
	// reading the clock in a zone ahead of UTC, nor stay current because it is
	// behind. The comparison is on calendar dates in UTC, so every instant
	// within the due date behaves the same.
	due := date(2026, time.October, 21)
	ahead := time.FixedZone("UTC+13", 13*60*60)
	behind := time.FixedZone("UTC-11", -11*60*60)

	for hour := 0; hour < 24; hour++ {
		asOf := time.Date(2026, time.October, 21, hour, 30, 0, 0, time.UTC)
		if IsOverdue(DocInvoice, StatusSent, &due, asOf) {
			t.Fatalf("overdue at %s, which is still the due date", asOf.Format(time.RFC3339))
		}
		if IsOverdue(DocInvoice, StatusSent, &due, asOf.In(ahead)) {
			t.Fatalf("overdue at %s read in UTC+13", asOf.Format(time.RFC3339))
		}
		if IsOverdue(DocInvoice, StatusSent, &due, asOf.In(behind)) {
			t.Fatalf("overdue at %s read in UTC-11", asOf.Format(time.RFC3339))
		}
	}

	firstMomentLate := date(2026, time.October, 22)
	if !IsOverdue(DocInvoice, StatusSent, &due, firstMomentLate) {
		t.Fatal("not overdue at the first instant of the following day")
	}
}

func TestUnsettledInvoiceStatusesAreTheOnesThatCanStillBePaid(t *testing.T) {
	// The invariant tying this set to the lifecycle: every status in it must be
	// one an invoice can legally move out of towards paid. If someone adds a
	// status here that cannot reach paid, the outstanding total starts counting
	// money nobody will ever pay.
	for _, status := range UnsettledInvoiceStatuses() {
		if status != StatusPaid && !CanTransition(DocInvoice, status, StatusPaid) {
			t.Errorf("%s is listed as unsettled but an invoice cannot move from it to paid", status)
		}
	}

	// And the settled ones must not be in it.
	for _, settled := range []DocStatus{StatusDraft, StatusPaid, StatusVoid} {
		for _, unsettled := range UnsettledInvoiceStatuses() {
			if settled == unsettled {
				t.Errorf("%s is settled or unissued but appears in UnsettledInvoiceStatuses", settled)
			}
		}
	}
}
