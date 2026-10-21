package domain

import (
	"errors"
	"testing"
)

func allStatuses() []DocStatus {
	return []DocStatus{
		StatusDraft, StatusSent, StatusAccepted, StatusDeclined,
		StatusConverted, StatusPartiallyPaid, StatusPaid, StatusVoid,
	}
}

func allTypes() []DocType {
	return []DocType{DocProposal, DocEstimate, DocInvoice}
}

func TestParseDocTypeRejectsAnythingOutsideTheSet(t *testing.T) {
	for _, s := range []string{"", "Invoice", "receipt", "quote", "invoices"} {
		if got, err := ParseDocType(s); err == nil {
			t.Errorf("ParseDocType(%q) = %q, want an error", s, got)
		} else if !errors.Is(err, ErrUnknownDocType) {
			t.Errorf("ParseDocType(%q) error = %v, want ErrUnknownDocType", s, err)
		}
	}
	for _, want := range allTypes() {
		if got, err := ParseDocType(string(want)); err != nil || got != want {
			t.Errorf("ParseDocType(%q) = %q, %v", want, got, err)
		}
	}
}

func TestNextStatusesAgreesWithCanTransition(t *testing.T) {
	// The UI builds its buttons from NextStatuses and the handler enforces with
	// CanTransition. If those two ever disagree the app offers a button that
	// then fails, so the invariant is that they are the same rule.
	for _, docType := range allTypes() {
		for _, from := range allStatuses() {
			offered := map[DocStatus]bool{}
			for _, to := range NextStatuses(docType, from) {
				offered[to] = true
			}
			for _, to := range allStatuses() {
				if offered[to] != CanTransition(docType, from, to) {
					t.Errorf("%s from %s to %s: NextStatuses says %v, CanTransition says %v",
						docType, from, to, offered[to], CanTransition(docType, from, to))
				}
			}
		}
	}
}

func TestTerminalStatusesAreDeadEnds(t *testing.T) {
	// Paid, declined, converted and void must have no way out. Without this, a
	// paid invoice could be walked back to draft and edited after the money
	// arrived, which is the one thing an invoice must never allow.
	terminal := []DocStatus{StatusPaid, StatusDeclined, StatusConverted, StatusVoid}
	for _, docType := range allTypes() {
		for _, from := range terminal {
			if next := NextStatuses(docType, from); len(next) != 0 {
				t.Errorf("%s in terminal status %s offers moves to %v", docType, from, next)
			}
		}
	}
}

func TestNoTypeEverAcceptsAnotherTypesStatus(t *testing.T) {
	// An invoice has no accepted/declined, and a proposal has no paid. Mixing
	// them is how a document ends up in a status its own screens cannot render.
	forbidden := map[DocType][]DocStatus{
		DocProposal: {StatusPartiallyPaid, StatusPaid},
		DocEstimate: {StatusPartiallyPaid, StatusPaid},
		DocInvoice:  {StatusAccepted, StatusDeclined, StatusConverted},
	}
	for docType, statuses := range forbidden {
		for _, bad := range statuses {
			for _, from := range allStatuses() {
				if CanTransition(docType, from, bad) {
					t.Errorf("%s must never reach %s, but %s allows it", docType, bad, from)
				}
			}
		}
	}
}

func TestTransitionReportsWhatItRefused(t *testing.T) {
	if _, err := Transition(DocInvoice, StatusPaid, StatusDraft); err == nil {
		t.Fatal("reviving a paid invoice to draft was allowed")
	} else if !errors.Is(err, ErrIllegalTransition) {
		t.Fatalf("error = %v, want ErrIllegalTransition", err)
	}

	got, err := Transition(DocInvoice, StatusDraft, StatusSent)
	if err != nil {
		t.Fatalf("sending a draft invoice returned error %v", err)
	}
	if got != StatusSent {
		t.Fatalf("status = %s, want %s", got, StatusSent)
	}
}

func TestConversionChainEndsAtInvoice(t *testing.T) {
	estimate, err := ConversionTarget(DocProposal)
	if err != nil || estimate != DocEstimate {
		t.Fatalf("proposal converts to %q, %v; want estimate", estimate, err)
	}
	invoice, err := ConversionTarget(DocEstimate)
	if err != nil || invoice != DocInvoice {
		t.Fatalf("estimate converts to %q, %v; want invoice", invoice, err)
	}
	if _, err := ConversionTarget(DocInvoice); !errors.Is(err, ErrNotConvertible) {
		t.Fatalf("invoice conversion error = %v, want ErrNotConvertible", err)
	}
}

func TestStatusForPayments(t *testing.T) {
	cases := []struct {
		name    string
		total   Money
		paid    Money
		current DocStatus
		want    DocStatus
	}{
		{"nothing paid stays put", 10000, 0, StatusSent, StatusSent},
		{"part paid", 10000, 2500, StatusSent, StatusPartiallyPaid},
		{"exactly paid", 10000, 10000, StatusSent, StatusPaid},
		{"overpaid is still paid", 10000, 12000, StatusPartiallyPaid, StatusPaid},
		{"a cent short is not paid", 10000, 9999, StatusSent, StatusPartiallyPaid},
		{"void stays void", 10000, 10000, StatusVoid, StatusVoid},
		{"refund past zero does not revive", 10000, -500, StatusSent, StatusSent},
		{"zero-total document is paid by any payment", 0, 1, StatusSent, StatusPaid},
	}
	for _, c := range cases {
		if got := StatusForPayments(c.total, c.paid, c.current); got != c.want {
			t.Errorf("%s: StatusForPayments(%s, %s, %s) = %s, want %s",
				c.name, c.total, c.paid, c.current, got, c.want)
		}
	}
}
