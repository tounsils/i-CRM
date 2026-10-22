// Package web serves the CRM over HTTP: routing, sessions, CSRF and the
// server-rendered screens. It holds no business rules; those are in the domain
// package, and persistence is in the store package.
package web

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/tounsils/i-CRM/internal/config"
	"github.com/tounsils/i-CRM/internal/domain"
	"github.com/tounsils/i-CRM/internal/store"
)

//go:embed templates
var templateFS embed.FS

// Server carries everything a request needs. One instance per process.
type Server struct {
	cfg              config.Settings
	store            *store.Store
	pages            map[string]*template.Template
	standalone       map[string]*template.Template
	throttle         *loginThrottle
	enumerationGuard []byte
}

// New parses the templates up front, so a broken template fails at boot rather
// than on the first request that happens to reach it.
func New(cfg config.Settings, st *store.Store) (*Server, error) {
	pages, err := parsePages()
	if err != nil {
		return nil, err
	}
	standalone, err := parseStandalone()
	if err != nil {
		return nil, err
	}
	guard, err := newEnumerationGuard()
	if err != nil {
		return nil, err
	}
	return &Server{
		cfg:              cfg,
		store:            st,
		pages:            pages,
		standalone:       standalone,
		throttle:         newLoginThrottle(),
		enumerationGuard: guard,
	}, nil
}

// Routes builds the handler.
//
// The router denies by default. The public surface is enumerated explicitly and
// everything else is reached through the catch-all, which is wrapped in
// requireAuth. Adding a route therefore makes it private unless someone goes
// out of their way to list it as public, so the cost of forgetting is a locked
// door rather than an open one.
func (s *Server) Routes() http.Handler {
	private := http.NewServeMux()
	s.registerPrivateRoutes(private)

	root := http.NewServeMux()
	root.HandleFunc("GET /healthz", s.handleHealth)
	root.HandleFunc("GET /static/app.css", s.handleStylesheet)
	root.HandleFunc("GET /login", s.handleLoginForm)
	root.HandleFunc("POST /login", s.handleLogin)
	root.HandleFunc("GET /setup", s.handleSetupForm)
	root.HandleFunc("POST /setup", s.handleSetup)
	root.Handle("/", s.requireAuth(private))

	return s.recoverPanic(s.secureHeaders(s.limitBody(root)))
}

func (s *Server) registerPrivateRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", s.handleDashboard)
	mux.HandleFunc("POST /logout", s.handleLogout)

	mux.HandleFunc("GET /customers", s.handleCustomerList)
	mux.HandleFunc("GET /customers/new", s.handleCustomerNew)
	mux.HandleFunc("POST /customers", s.handleCustomerCreate)
	mux.HandleFunc("GET /customers/{id}", s.handleCustomerShow)
	mux.HandleFunc("GET /customers/{id}/edit", s.handleCustomerEdit)
	mux.HandleFunc("POST /customers/{id}", s.handleCustomerUpdate)
	mux.HandleFunc("POST /customers/{id}/messages", s.handleMessageCreate)

	mux.HandleFunc("GET /documents", s.handleDocumentList)
	mux.HandleFunc("GET /documents/new", s.handleDocumentNew)
	mux.HandleFunc("POST /documents", s.handleDocumentCreate)
	mux.HandleFunc("GET /documents/{id}", s.handleDocumentShow)
	mux.HandleFunc("GET /documents/{id}/print", s.handleDocumentPrint)
	mux.HandleFunc("GET /documents/{id}/pdf", s.handleDocumentPDF)
	mux.HandleFunc("POST /documents/{id}/lines", s.handleDocumentLines)
	mux.HandleFunc("POST /documents/{id}/status", s.handleDocumentStatus)
	mux.HandleFunc("POST /documents/{id}/convert", s.handleDocumentConvert)
	mux.HandleFunc("POST /documents/{id}/payments", s.handlePaymentCreate)

	mux.HandleFunc("GET /catalog", s.handleCatalog)
	mux.HandleFunc("POST /catalog/taxes", s.handleTaxCreate)
	mux.HandleFunc("POST /catalog/items", s.handleItemCreate)
}

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// A failed write to a healthcheck probe is not actionable: the probe has
	// already gone away, and there is nothing to retry or report.
	_, _ = fmt.Fprintln(w, "ok")
}

// pageData is what every template receives. Keeping one shape means the layout
// can rely on User, CSRFToken and the banners being present on every page.
type pageData struct {
	Title       string
	CompanyName string
	User        store.User
	CSRFToken   string
	Notice      string
	Error       string
	Data        any
}

// documentPill is the badge shown for a document's state.
type documentPill struct {
	Label string
	Tone  string
}

var templateFuncs = template.FuncMap{
	// pill is the state badge for a document. An unsettled invoice past its due
	// date reads as Overdue rather than as its stored status, because overdue
	// is the fact the reader needs; the stored status is still Sent, and
	// nothing was swept into a new state to achieve this.
	//
	// The clock is read here rather than threaded in from the handler because
	// this is presentation only. The rule itself is domain.IsOverdue and is
	// tested there against fixed times.
	"pill": func(d store.Document) documentPill {
		if d.Overdue(time.Now()) {
			return documentPill{Label: "Overdue", Tone: "tone-bad"}
		}
		return documentPill{Label: d.Status.Label(), Tone: statusTone(d.Status)}
	},

	"date": func(t time.Time) string { return t.Format("2 Jan 2006") },
	"optionalDate": func(t *time.Time) string {
		if t == nil {
			return ""
		}
		return t.Format("2 Jan 2006")
	},
	"inputDate":  func(t time.Time) string { return t.Format("2006-01-02") },
	"statusTone": statusTone,

	// payable decides whether to offer the record-a-payment form. It lives here
	// rather than as a chain of comparisons in the template so that the rule
	// matches store.RecordPayment exactly and can be read in one place.
	"payable": func(d store.Document) bool {
		return d.Type == domain.DocInvoice &&
			d.Status != domain.StatusDraft &&
			d.Status != domain.StatusVoid
	},

	// editable decides whether to offer the line editor.
	"editable": func(d store.Document) bool {
		return d.Status == domain.StatusDraft
	},

	// rows gives range something to walk when a form needs N empty rows.
	// Templates cannot iterate over an integer.
	"rows": func(n int) []int {
		out := make([]int, n)
		for i := range out {
			out[i] = i
		}
		return out
	},
}

// statusTone maps a status to a CSS class, so the template does not carry a
// chain of comparisons against status literals.
func statusTone(s domain.DocStatus) string {
	switch s {
	case domain.StatusPaid, domain.StatusAccepted:
		return "tone-good"
	case domain.StatusDeclined, domain.StatusVoid:
		return "tone-bad"
	case domain.StatusSent, domain.StatusPartiallyPaid:
		return "tone-warn"
	default:
		return "tone-neutral"
	}
}

// parsePages pairs each page with the layout. Pages are parsed in separate sets
// rather than all together because every page defines a template named
// "content", which would collide in one set.
func parsePages() (map[string]*template.Template, error) {
	names, err := fs.Glob(templateFS, "templates/pages/*.html")
	if err != nil {
		return nil, fmt.Errorf("web: listing page templates: %w", err)
	}
	if len(names) == 0 {
		return nil, fmt.Errorf("web: no page templates were embedded")
	}

	out := make(map[string]*template.Template, len(names))
	for _, name := range names {
		t, err := template.New("layout.html").Funcs(templateFuncs).
			ParseFS(templateFS, "templates/layout.html", name)
		if err != nil {
			return nil, fmt.Errorf("web: parsing %s: %w", name, err)
		}
		out[path.Base(name)] = t
	}
	return out, nil
}

// parseStandalone parses the pages that carry their own <html> and are not
// wrapped in the layout, because they are documents rather than screens: no
// navigation, no sign-out form, nothing that should appear on paper.
func parseStandalone() (map[string]*template.Template, error) {
	names, err := fs.Glob(templateFS, "templates/standalone/*.html")
	if err != nil {
		return nil, fmt.Errorf("web: listing standalone templates: %w", err)
	}

	out := make(map[string]*template.Template, len(names))
	for _, name := range names {
		base := path.Base(name)
		t, err := template.New(base).Funcs(templateFuncs).ParseFS(templateFS, name)
		if err != nil {
			return nil, fmt.Errorf("web: parsing %s: %w", name, err)
		}
		out[base] = t
	}
	return out, nil
}

// renderStandalone writes a page that is not wrapped in the layout.
//
// Buffered for the same reason as render: a template error must become a 500
// rather than a 200 carrying half a document.
func (s *Server) renderStandalone(w http.ResponseWriter, r *http.Request, page string, data any) {
	t, ok := s.standalone[page]
	if !ok {
		s.fail(w, r, http.StatusInternalServerError, fmt.Errorf("unknown standalone page %q", page))
		return
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		s.fail(w, r, http.StatusInternalServerError, fmt.Errorf("rendering %s: %w", page, err))
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := buf.WriteTo(w); err != nil {
		log.Printf("web: writing response for %s: %v", r.URL.Path, err)
	}
}

// render writes a page.
//
// The template is executed into a buffer first. Executing straight to the
// ResponseWriter would commit a 200 and a half-rendered page before a template
// error could be reported, which turns a template bug into a corrupt response
// instead of a 500.
func (s *Server) render(w http.ResponseWriter, r *http.Request, page string, data pageData) {
	t, ok := s.pages[page]
	if !ok {
		s.fail(w, r, http.StatusInternalServerError, fmt.Errorf("unknown page %q", page))
		return
	}

	data.CompanyName = s.cfg.CompanyName
	if user, ok := userFrom(r.Context()); ok {
		data.User = user
	}
	if data.Notice == "" {
		data.Notice = r.URL.Query().Get("notice")
	}
	if data.Error == "" {
		data.Error = r.URL.Query().Get("error")
	}
	if data.CSRFToken == "" {
		data.CSRFToken = s.ensureCSRFToken(w, r)
	}

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "layout.html", data); err != nil {
		s.fail(w, r, http.StatusInternalServerError, fmt.Errorf("rendering %s: %w", page, err))
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := buf.WriteTo(w); err != nil {
		log.Printf("web: writing response for %s: %v", r.URL.Path, err)
	}
}

// fail logs the real error and returns only a status text to the client, so an
// SQL error or a file path never reaches a browser.
func (s *Server) fail(w http.ResponseWriter, r *http.Request, code int, err error) {
	log.Printf("web: %s %s: %v", r.Method, r.URL.Path, err)
	http.Error(w, http.StatusText(code), code)
}

// redirect sends the browser on with an optional banner message. Using a query
// parameter rather than a server-side flash store keeps the process stateless,
// which matters for a single binary an operator may restart at any time.
func redirect(w http.ResponseWriter, r *http.Request, to, notice string) {
	if notice != "" {
		to += "?notice=" + url.QueryEscape(notice)
	}
	http.Redirect(w, r, to, http.StatusSeeOther)
}

// redirectError sends the browser back with an error banner.
func redirectError(w http.ResponseWriter, r *http.Request, to, message string) {
	http.Redirect(w, r, to+"?error="+url.QueryEscape(message), http.StatusSeeOther)
}
