package web

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/tounsils/i-CRM/internal/domain"
	"github.com/tounsils/i-CRM/internal/store"
)

// customerPage is what the customer screens need beyond the common page data.
type customerPage struct {
	Customer   store.Customer
	Customers  []store.Customer
	Documents  []store.Document
	Messages   []store.Message
	Statuses   []domain.CustomerStatus
	Directions []domain.MessageDirection
	Filter     domain.CustomerStatus
}

func (s *Server) handleCustomerList(w http.ResponseWriter, r *http.Request) {
	var filter domain.CustomerStatus
	if raw := r.URL.Query().Get("status"); raw != "" {
		parsed, err := domain.ParseCustomerStatus(raw)
		if err != nil {
			redirectError(w, r, "/customers", "That is not a customer status.")
			return
		}
		filter = parsed
	}

	customers, err := s.store.ListCustomers(r.Context(), filter)
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	s.render(w, r, "customers.html", pageData{
		Title: "Customers",
		Data: customerPage{
			Customers: customers,
			Statuses:  domain.CustomerStatuses(),
			Filter:    filter,
		},
	})
}

func (s *Server) handleCustomerNew(w http.ResponseWriter, r *http.Request) {
	s.render(w, r, "customer_form.html", pageData{
		Title: "New customer",
		Data: customerPage{
			Customer: store.Customer{Status: domain.CustomerLead, Currency: s.cfg.DefaultCurrency},
			Statuses: domain.CustomerStatuses(),
		},
	})
}

func (s *Server) handleCustomerCreate(w http.ResponseWriter, r *http.Request) {
	customer, problem := s.customerFromForm(r)
	if problem != "" {
		s.render(w, r, "customer_form.html", pageData{
			Title: "New customer",
			Error: problem,
			Data:  customerPage{Customer: customer, Statuses: domain.CustomerStatuses()},
		})
		return
	}

	id, err := s.store.CreateCustomer(r.Context(), customer)
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	redirect(w, r, fmt.Sprintf("/customers/%d", id), "Customer saved.")
}

func (s *Server) handleCustomerShow(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(r)
	if !ok {
		http.NotFound(w, r)
		return
	}

	customer, err := s.store.CustomerByID(r.Context(), id)
	if errors.Is(err, store.ErrNotFound) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}

	documents, err := s.store.ListDocuments(r.Context(), store.DocumentFilter{CustomerID: id})
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	messages, err := s.store.MessagesForCustomer(r.Context(), id)
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}

	s.render(w, r, "customer_show.html", pageData{
		Title: customer.Name,
		Data: customerPage{
			Customer:   customer,
			Documents:  documents,
			Messages:   messages,
			Directions: domain.MessageDirections(),
		},
	})
}

func (s *Server) handleCustomerEdit(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(r)
	if !ok {
		http.NotFound(w, r)
		return
	}
	customer, err := s.store.CustomerByID(r.Context(), id)
	if errors.Is(err, store.ErrNotFound) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	s.render(w, r, "customer_form.html", pageData{
		Title: "Edit " + customer.Name,
		Data:  customerPage{Customer: customer, Statuses: domain.CustomerStatuses()},
	})
}

func (s *Server) handleCustomerUpdate(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(r)
	if !ok {
		http.NotFound(w, r)
		return
	}

	customer, problem := s.customerFromForm(r)
	customer.ID = id
	if problem != "" {
		s.render(w, r, "customer_form.html", pageData{
			Title: "Edit customer",
			Error: problem,
			Data:  customerPage{Customer: customer, Statuses: domain.CustomerStatuses()},
		})
		return
	}

	if err := s.store.UpdateCustomer(r.Context(), customer); errors.Is(err, store.ErrNotFound) {
		http.NotFound(w, r)
		return
	} else if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	redirect(w, r, fmt.Sprintf("/customers/%d", id), "Customer updated.")
}

// customerFromForm validates the submitted fields.
//
// The second return is a message for the user, empty when the record is valid.
// It is deliberately not an error: these are sentences a person reads on a
// form, they are never logged, and typing them as errors would invite a caller
// to treat a mistyped currency as a server fault.
//
// The partially built record comes back alongside it so the form can be
// re-rendered with what the user typed rather than emptied.
func (s *Server) customerFromForm(r *http.Request) (store.Customer, string) {
	c := store.Customer{
		ContactName: strings.TrimSpace(r.PostFormValue("contact_name")),
		Email:       strings.TrimSpace(r.PostFormValue("email")),
		Phone:       strings.TrimSpace(r.PostFormValue("phone")),
		Address:     strings.TrimSpace(r.PostFormValue("address")),
		Notes:       strings.TrimSpace(r.PostFormValue("notes")),
		Name:        strings.TrimSpace(r.PostFormValue("name")),
		Status:      domain.CustomerLead,
		Currency:    s.cfg.DefaultCurrency,
	}

	name, err := domain.ValidateName(r.PostFormValue("name"))
	if err != nil {
		return c, "A customer needs a name."
	}
	c.Name = name

	status, err := domain.ParseCustomerStatus(r.PostFormValue("status"))
	if err != nil {
		return c, "Choose a status of lead, customer or archived."
	}
	c.Status = status

	currency, err := domain.ValidateCurrency(r.PostFormValue("currency"))
	if err != nil {
		return c, "The currency must be a three-letter code such as USD."
	}
	c.Currency = currency

	return c, ""
}

func (s *Server) handleMessageCreate(w http.ResponseWriter, r *http.Request) {
	customerID, ok := pathID(r)
	if !ok {
		http.NotFound(w, r)
		return
	}
	user, ok := userFrom(r.Context())
	if !ok {
		s.fail(w, r, http.StatusInternalServerError, errors.New("no user on an authenticated request"))
		return
	}
	target := fmt.Sprintf("/customers/%d", customerID)

	direction, err := domain.ParseMessageDirection(r.PostFormValue("direction"))
	if err != nil {
		redirectError(w, r, target, "Choose whether this was sent, received or an internal note.")
		return
	}
	body := strings.TrimSpace(r.PostFormValue("body"))
	if body == "" {
		redirectError(w, r, target, "A message needs a body.")
		return
	}

	_, err = s.store.CreateMessage(r.Context(), store.Message{
		CustomerID: customerID,
		AuthorID:   user.ID,
		Direction:  direction,
		Subject:    strings.TrimSpace(r.PostFormValue("subject")),
		Body:       body,
	})
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	redirect(w, r, target, "Message added to the timeline.")
}

// pathID reads the {id} path segment.
func pathID(r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		return 0, false
	}
	return id, true
}
