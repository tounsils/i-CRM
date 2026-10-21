package web

import (
	"net/http"

	"github.com/tounsils/i-CRM/internal/domain"
	"github.com/tounsils/i-CRM/internal/store"
)

// dashboardPage is the landing screen: what is owed, and what moved recently.
type dashboardPage struct {
	CustomerCounts map[domain.CustomerStatus]int
	Statuses       []domain.CustomerStatus
	Outstanding    domain.Money
	Currency       string
	Recent         []store.Document
	Unpaid         []store.Document
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
	unpaid, err := s.store.ListDocuments(ctx, store.DocumentFilter{
		Type:   domain.DocInvoice,
		Status: domain.StatusSent,
	})
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}

	s.render(w, r, "dashboard.html", pageData{
		Title: "Dashboard",
		Data: dashboardPage{
			CustomerCounts: counts,
			Statuses:       domain.CustomerStatuses(),
			Outstanding:    outstanding,
			Currency:       s.cfg.DefaultCurrency,
			Recent:         capDocuments(recent, recentDocumentLimit),
			Unpaid:         capDocuments(unpaid, recentDocumentLimit),
		},
	})
}

func capDocuments(documents []store.Document, limit int) []store.Document {
	if len(documents) <= limit {
		return documents
	}
	return documents[:limit]
}
