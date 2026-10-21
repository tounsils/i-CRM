package web

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/tounsils/i-CRM/internal/domain"
	"github.com/tounsils/i-CRM/internal/store"
)

// documentPage is what the document screens need beyond the common page data.
type documentPage struct {
	Document         store.Document
	Documents        []store.Document
	Customers        []store.Customer
	Taxes            []store.Tax
	Items            []store.Item
	Payments         []store.Payment
	NextStatuses     []domain.DocStatus
	ConversionTarget domain.DocType
	Convertible      bool
	Types            []domain.DocType
	Filter           store.DocumentFilter
	Today            string
	BlankLines       int
}

// blankLineRows is how many empty line rows a new document form offers. Five
// covers the great majority of documents without scrolling, and more can be
// added by saving the draft and editing it again.
const blankLineRows = 5

func documentTypes() []domain.DocType {
	return []domain.DocType{domain.DocProposal, domain.DocEstimate, domain.DocInvoice}
}

func (s *Server) handleDocumentList(w http.ResponseWriter, r *http.Request) {
	filter := store.DocumentFilter{}
	query := r.URL.Query()

	if raw := query.Get("type"); raw != "" {
		parsed, err := domain.ParseDocType(raw)
		if err != nil {
			redirectError(w, r, "/documents", "That is not a document type.")
			return
		}
		filter.Type = parsed
	}
	if raw := query.Get("status"); raw != "" {
		filter.Status = domain.DocStatus(raw)
	}

	documents, err := s.store.ListDocuments(r.Context(), filter)
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	s.render(w, r, "documents.html", pageData{
		Title: "Documents",
		Data:  documentPage{Documents: documents, Types: documentTypes(), Filter: filter},
	})
}

func (s *Server) handleDocumentNew(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	docType := domain.DocInvoice
	if raw := query.Get("type"); raw != "" {
		parsed, err := domain.ParseDocType(raw)
		if err != nil {
			redirectError(w, r, "/documents", "That is not a document type.")
			return
		}
		docType = parsed
	}

	page, err := s.documentFormData(r, docType)
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	if raw := query.Get("customer"); raw != "" {
		if id, err := strconv.ParseInt(raw, 10, 64); err == nil {
			page.Document.CustomerID = id
		}
	}

	s.render(w, r, "document_form.html", pageData{
		Title: "New " + strings.ToLower(docType.Label()),
		Data:  page,
	})
}

// documentFormData gathers the pickers a document form needs.
func (s *Server) documentFormData(r *http.Request, docType domain.DocType) (documentPage, error) {
	customers, err := s.store.ListCustomers(r.Context(), "")
	if err != nil {
		return documentPage{}, err
	}
	taxes, err := s.store.ListTaxes(r.Context())
	if err != nil {
		return documentPage{}, err
	}
	items, err := s.store.ListItems(r.Context())
	if err != nil {
		return documentPage{}, err
	}

	now := time.Now()
	doc := store.Document{
		Type:      docType,
		Currency:  s.cfg.DefaultCurrency,
		IssueDate: now,
	}
	if docType == domain.DocInvoice {
		due := s.cfg.DueDateFor(now)
		doc.DueDate = &due
	}

	return documentPage{
		Document:   doc,
		Customers:  customers,
		Taxes:      taxes,
		Items:      items,
		Types:      documentTypes(),
		Today:      now.Format("2006-01-02"),
		BlankLines: blankLineRows,
	}, nil
}

func (s *Server) handleDocumentCreate(w http.ResponseWriter, r *http.Request) {
	docType, err := domain.ParseDocType(r.PostFormValue("type"))
	if err != nil {
		redirectError(w, r, "/documents", "That is not a document type.")
		return
	}

	reject := func(message string) {
		page, err := s.documentFormData(r, docType)
		if err != nil {
			s.fail(w, r, http.StatusInternalServerError, err)
			return
		}
		page.Document.CustomerID, _ = strconv.ParseInt(r.PostFormValue("customer_id"), 10, 64)
		page.Document.Notes = r.PostFormValue("notes")
		s.render(w, r, "document_form.html", pageData{
			Title: "New " + strings.ToLower(docType.Label()),
			Error: message,
			Data:  page,
		})
	}

	customerID, err := strconv.ParseInt(r.PostFormValue("customer_id"), 10, 64)
	if err != nil || customerID <= 0 {
		reject("Choose a customer.")
		return
	}
	if _, err := s.store.CustomerByID(r.Context(), customerID); errors.Is(err, store.ErrNotFound) {
		reject("That customer no longer exists.")
		return
	} else if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}

	issue, err := parseDateField(r.PostFormValue("issue_date"), time.Now())
	if err != nil {
		reject("The issue date must be a real date.")
		return
	}
	due, err := parseOptionalDateField(r.PostFormValue("due_date"))
	if err != nil {
		reject("The due date must be a real date.")
		return
	}
	if due == nil && docType == domain.DocInvoice {
		// Fall back to the configured payment terms rather than leaving an
		// invoice with no date by which it is late.
		fallback := s.cfg.DueDateFor(issue)
		due = &fallback
	}

	currency, err := domain.ValidateCurrency(r.PostFormValue("currency"))
	if err != nil {
		reject("The currency must be a three-letter code such as USD.")
		return
	}

	lines, problem := s.linesFromForm(r)
	if problem != "" {
		reject(problem)
		return
	}

	id, err := s.store.CreateDocument(r.Context(), store.Document{
		Type:       docType,
		CustomerID: customerID,
		Currency:   currency,
		IssueDate:  issue,
		DueDate:    due,
		Notes:      strings.TrimSpace(r.PostFormValue("notes")),
	}, lines)
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	redirect(w, r, fmt.Sprintf("/documents/%d", id), docType.Label()+" created as a draft.")
}

func (s *Server) handleDocumentShow(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(r)
	if !ok {
		http.NotFound(w, r)
		return
	}

	doc, err := s.store.DocumentByID(r.Context(), id)
	if errors.Is(err, store.ErrNotFound) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}

	payments, err := s.store.PaymentsForDocument(r.Context(), id)
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	taxes, err := s.store.ListTaxes(r.Context())
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}

	page := documentPage{
		Document:     doc,
		Payments:     payments,
		Taxes:        taxes,
		NextStatuses: domain.NextStatuses(doc.Type, doc.Status),
		Today:        time.Now().Format("2006-01-02"),
	}
	// The convert button is offered only when the rules would actually allow
	// it, so the UI never shows an action that the handler then refuses.
	if target, err := domain.ConversionTarget(doc.Type); err == nil {
		page.ConversionTarget = target
		page.Convertible = doc.Status == domain.StatusAccepted
	}

	s.render(w, r, "document_show.html", pageData{
		Title: doc.Type.Label() + " " + doc.Number,
		Data:  page,
	})
}

func (s *Server) handleDocumentLines(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(r)
	if !ok {
		http.NotFound(w, r)
		return
	}
	target := fmt.Sprintf("/documents/%d", id)

	lines, problem := s.linesFromForm(r)
	if problem != "" {
		redirectError(w, r, target, problem)
		return
	}

	err := s.store.ReplaceLines(r.Context(), id, lines)
	switch {
	case errors.Is(err, store.ErrNotFound):
		http.NotFound(w, r)
	case errors.Is(err, store.ErrNotEditable):
		redirectError(w, r, target, "This document has been issued, so its lines can no longer be changed.")
	case err != nil:
		s.fail(w, r, http.StatusInternalServerError, err)
	default:
		redirect(w, r, target, "Lines updated.")
	}
}

func (s *Server) handleDocumentStatus(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(r)
	if !ok {
		http.NotFound(w, r)
		return
	}
	target := fmt.Sprintf("/documents/%d", id)

	to := domain.DocStatus(r.PostFormValue("status"))
	err := s.store.SetStatus(r.Context(), id, to)
	switch {
	case errors.Is(err, store.ErrNotFound):
		http.NotFound(w, r)
	case errors.Is(err, domain.ErrIllegalTransition):
		redirectError(w, r, target, "That is not a move this document can make from where it is.")
	case err != nil:
		s.fail(w, r, http.StatusInternalServerError, err)
	default:
		redirect(w, r, target, "Marked as "+strings.ToLower(to.Label())+".")
	}
}

func (s *Server) handleDocumentConvert(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(r)
	if !ok {
		http.NotFound(w, r)
		return
	}
	target := fmt.Sprintf("/documents/%d", id)

	doc, err := s.store.DocumentByID(r.Context(), id)
	if errors.Is(err, store.ErrNotFound) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}

	// Payment terms are owned by config, which is why the due date is computed
	// here and passed in rather than decided inside the store.
	var due *time.Time
	if next, err := domain.ConversionTarget(doc.Type); err == nil && next == domain.DocInvoice {
		computed := s.cfg.DueDateFor(time.Now())
		due = &computed
	}

	newID, err := s.store.Convert(r.Context(), id, due)
	switch {
	case errors.Is(err, domain.ErrNotConvertible):
		redirectError(w, r, target, "An invoice is the end of the chain, so there is nothing to convert it into.")
	case errors.Is(err, domain.ErrIllegalTransition):
		redirectError(w, r, target, "Only an accepted document can be converted.")
	case err != nil:
		s.fail(w, r, http.StatusInternalServerError, err)
	default:
		redirect(w, r, fmt.Sprintf("/documents/%d", newID), "Converted. The new document is a draft.")
	}
}

func (s *Server) handlePaymentCreate(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(r)
	if !ok {
		http.NotFound(w, r)
		return
	}
	target := fmt.Sprintf("/documents/%d", id)

	amount, err := domain.ParseMoney(r.PostFormValue("amount"))
	if err != nil {
		redirectError(w, r, target, "The amount must be a figure such as 250.00.")
		return
	}
	if amount <= 0 {
		redirectError(w, r, target, "A payment has to be more than nothing.")
		return
	}
	paidOn, err := parseDateField(r.PostFormValue("paid_on"), time.Now())
	if err != nil {
		redirectError(w, r, target, "The payment date must be a real date.")
		return
	}

	_, err = s.store.RecordPayment(r.Context(), store.Payment{
		DocumentID: id,
		Amount:     amount,
		Method:     strings.TrimSpace(r.PostFormValue("method")),
		Reference:  strings.TrimSpace(r.PostFormValue("reference")),
		PaidOn:     paidOn,
	})
	switch {
	case errors.Is(err, store.ErrNotFound):
		http.NotFound(w, r)
	case errors.Is(err, store.ErrNotPayable):
		redirectError(w, r, target, "Only an invoice that has been sent can take a payment.")
	case err != nil:
		s.fail(w, r, http.StatusInternalServerError, err)
	default:
		redirect(w, r, target, "Payment recorded.")
	}
}

// linesFromForm reads the repeated line fields.
//
// A row whose description is blank is skipped rather than rejected, because the
// form always renders spare rows and a user who fills three of five has not
// made a mistake. The second return is a message for the user, empty when the
// lines are usable.
func (s *Server) linesFromForm(r *http.Request) ([]domain.Line, string) {
	taxes, err := s.store.ListTaxes(r.Context())
	if err != nil {
		return nil, "The tax rates could not be read."
	}
	byID := make(map[int64]store.Tax, len(taxes))
	for _, t := range taxes {
		byID[t.ID] = t
	}

	descriptions := r.PostForm["line_description"]
	quantities := r.PostForm["line_quantity"]
	prices := r.PostForm["line_unit_price"]
	taxIDs := r.PostForm["line_tax_id"]

	var lines []domain.Line
	for i, rawDescription := range descriptions {
		description := strings.TrimSpace(rawDescription)
		if description == "" {
			continue
		}

		quantity := domain.Quantity(domain.QuantityScale)
		if raw := fieldAt(quantities, i); raw != "" {
			quantity, err = domain.ParseQuantity(raw)
			if err != nil {
				return nil, fmt.Sprintf("Line %d: the quantity must be a number such as 1.5.", i+1)
			}
		}
		if quantity <= 0 {
			return nil, fmt.Sprintf("Line %d: the quantity must be more than nothing.", i+1)
		}

		var price domain.Money
		if raw := fieldAt(prices, i); raw != "" {
			price, err = domain.ParseMoney(raw)
			if err != nil {
				return nil, fmt.Sprintf("Line %d: the unit price must be a figure such as 150.00.", i+1)
			}
		}

		line := domain.Line{Description: description, Quantity: quantity, UnitPrice: price}
		if raw := fieldAt(taxIDs, i); raw != "" && raw != "0" {
			taxID, err := strconv.ParseInt(raw, 10, 64)
			if err != nil {
				return nil, fmt.Sprintf("Line %d: that is not a tax rate.", i+1)
			}
			tax, ok := byID[taxID]
			if !ok {
				return nil, fmt.Sprintf("Line %d: that tax rate no longer exists.", i+1)
			}
			// Copied onto the line, not referenced, so a later rate change
			// cannot restate a document that has already been sent.
			line.TaxName, line.TaxRate = tax.Name, tax.Rate
		}
		lines = append(lines, line)
	}

	if len(lines) == 0 {
		return nil, "A document needs at least one line with a description."
	}
	return lines, ""
}

func fieldAt(values []string, i int) string {
	if i >= len(values) {
		return ""
	}
	return strings.TrimSpace(values[i])
}

// parseDateField reads a date input, falling back when the field is empty.
func parseDateField(value string, fallback time.Time) (time.Time, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return fallback, nil
	}
	return time.Parse("2006-01-02", value)
}

// parseOptionalDateField reads a date input that is allowed to be empty.
func parseOptionalDateField(value string) (*time.Time, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil, nil
	}
	parsed, err := time.Parse("2006-01-02", value)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}
