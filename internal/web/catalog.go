package web

import (
	"net/http"
	"strings"

	"github.com/tounsils/i-CRM/internal/domain"
	"github.com/tounsils/i-CRM/internal/store"
)

// catalogPage holds the tax rates and the reusable items.
type catalogPage struct {
	Taxes []store.Tax
	Items []store.Item
}

func (s *Server) handleCatalog(w http.ResponseWriter, r *http.Request) {
	taxes, err := s.store.ListTaxes(r.Context())
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	items, err := s.store.ListItems(r.Context())
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	s.render(w, r, "catalog.html", pageData{
		Title: "Taxes and items",
		Data:  catalogPage{Taxes: taxes, Items: items},
	})
}

func (s *Server) handleTaxCreate(w http.ResponseWriter, r *http.Request) {
	name, err := domain.ValidateName(r.PostFormValue("name"))
	if err != nil {
		redirectError(w, r, "/catalog", "A tax rate needs a name, such as VAT.")
		return
	}
	rate, err := domain.ParseRate(r.PostFormValue("rate"))
	if err != nil {
		redirectError(w, r, "/catalog", "The rate must be a percentage between 0 and 100, such as 20 or 8.25.")
		return
	}

	if _, err := s.store.CreateTax(r.Context(), store.Tax{Name: name, Rate: rate}); err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	redirect(w, r, "/catalog", "Tax rate added.")
}

func (s *Server) handleItemCreate(w http.ResponseWriter, r *http.Request) {
	name, err := domain.ValidateName(r.PostFormValue("name"))
	if err != nil {
		redirectError(w, r, "/catalog", "An item needs a name.")
		return
	}
	price, err := domain.ParseMoney(r.PostFormValue("unit_price"))
	if err != nil {
		redirectError(w, r, "/catalog", "The price must be a figure such as 150.00.")
		return
	}

	_, err = s.store.CreateItem(r.Context(), store.Item{
		Name:        name,
		Description: strings.TrimSpace(r.PostFormValue("description")),
		UnitPrice:   price,
	})
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	redirect(w, r, "/catalog", "Item added.")
}
