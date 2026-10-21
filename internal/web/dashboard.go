package web

import (
	"net/http"
	"time"

	"github.com/tounsils/i-CRM/internal/domain"
	"github.com/tounsils/i-CRM/internal/store"
)

// dashboardPage is the landing screen: what is owed, what is late, and what
// moved recently.
type dashboardPage struct {
	CustomerCounts map[domain.CustomerStatus]int
	Statuses       []domain.CustomerStatus
	Outstanding    domain.Money
	OverdueTotal   domain.Money
	Currency       string
	Recent         []store.Document
	Unsettled      []store.Document
	Overdue        []store.Document
}

// recentDocumentLimit caps the dashboard lists. Ten is a screenful; the full
// lists are one click away on /documents, so a higher number would only make
// the page longer without answering a question the dashboard is asked.
const recentDocumentLimit = 10

func (s *Server) handleDashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	counts, err := s.store.CountCustomersByStatus(ctx)
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	outstanding, err := s.store.OutstandingTotal(ctx)
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	recent, err := s.store.ListDocuments(ctx, store.DocumentFilter{})
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	unsettled, err := s.store.ListUnsettledInvoices(ctx)
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}

	// Overdue is derived here from the invoices that are still owed, rather
	// than fetched by a second query carrying a copy of the date comparison.
	// One clock reading is used for the whole page so the tile and the table
	// cannot disagree about where the boundary fell.
	asOf := time.Now()
	var (
		overdue      []store.Document
		overdueTotal domain.Money
	)
	for _, invoice := range unsettled {
		if invoice.Overdue(asOf) {
			overdue = append(overdue, invoice)
			overdueTotal += invoice.Outstanding()
		}
	}

	s.render(w, r, "dashboard.html", pageData{
		Title: "Dashboard",
		Data: dashboardPage{
			CustomerCounts: counts,
			Statuses:       domain.CustomerStatuses(),
			Outstanding:    outstanding,
			OverdueTotal:   overdueTotal,
			Currency:       s.cfg.DefaultCurrency,
			Recent:         capDocuments(recent, recentDocumentLimit),
			Unsettled:      capDocuments(unsettled, recentDocumentLimit),
			Overdue:        capDocuments(overdue, recentDocumentLimit),
		},
	})
}

func capDocuments(documents []store.Document, limit int) []store.Document {
	if len(documents) <= limit {
		return documents
	}
	return documents[:limit]
}
