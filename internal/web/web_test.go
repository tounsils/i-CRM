package web

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/tounsils/i-CRM/internal/config"
	"github.com/tounsils/i-CRM/internal/store"
)

// testSessionKey is a synthetic fixture, not a credential. It signs nothing
// outside this test binary and is deliberately obvious about what it is.
var testSessionKey = []byte("synthetic-test-key-not-a-secret-0")

const testPassword = "correct horse battery staple"

func testSettings() config.Settings {
	return config.Settings{
		Addr:             ":0",
		SessionKey:       testSessionKey,
		PaymentTermsDays: 30,
		CompanyName:      "Test Company",
		DefaultCurrency:  "USD",
	}
}

// harness is a running server with a browser-like client in front of it.
type harness struct {
	t      *testing.T
	store  *store.Store
	server *httptest.Server
	client *http.Client
	jar    *cookiejar.Jar
	base   *url.URL
}

func newHarness(t *testing.T) *harness {
	t.Helper()

	db, err := store.Open(filepath.Join(t.TempDir(), "web_test.db"))
	if err != nil {
		t.Fatalf("store.Open: %v", err)
	}
	t.Cleanup(func() { db.Close() })

	srv, err := New(testSettings(), db)
	if err != nil {
		t.Fatalf("web.New: %v", err)
	}

	ts := httptest.NewServer(srv.Routes())
	t.Cleanup(ts.Close)

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatalf("cookiejar.New: %v", err)
	}
	base, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("parsing test server URL: %v", err)
	}

	return &harness{
		t:      t,
		store:  db,
		server: ts,
		// Redirects are not followed, so a test can assert on the 303 and its
		// Location rather than on whatever page it lands on.
		client: &http.Client{
			Jar:           jar,
			CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
		},
		jar:  jar,
		base: base,
	}
}

func (h *harness) get(path string) *http.Response {
	h.t.Helper()
	res, err := h.client.Get(h.server.URL + path)
	if err != nil {
		h.t.Fatalf("GET %s: %v", path, err)
	}
	return res
}

// post submits a form, adding the CSRF token the browser would send.
func (h *harness) post(path string, form url.Values) *http.Response {
	h.t.Helper()
	if form == nil {
		form = url.Values{}
	}
	if token := h.csrfToken(); token != "" {
		form.Set(csrfFieldName, token)
	}
	res, err := h.client.PostForm(h.server.URL+path, form)
	if err != nil {
		h.t.Fatalf("POST %s: %v", path, err)
	}
	return res
}

// postWithoutCSRF submits a form with no token, as a cross-site forgery would.
func (h *harness) postWithoutCSRF(path string, form url.Values) *http.Response {
	h.t.Helper()
	if form == nil {
		form = url.Values{}
	}
	res, err := h.client.PostForm(h.server.URL+path, form)
	if err != nil {
		h.t.Fatalf("POST %s: %v", path, err)
	}
	return res
}

func (h *harness) csrfToken() string {
	for _, cookie := range h.jar.Cookies(h.base) {
		if cookie.Name == csrfCookieName {
			return cookie.Value
		}
	}
	return ""
}

func (h *harness) hasSession() bool {
	for _, cookie := range h.jar.Cookies(h.base) {
		if cookie.Name == sessionCookieName && cookie.Value != "" {
			return true
		}
	}
	return false
}

func (h *harness) body(res *http.Response) string {
	h.t.Helper()
	defer res.Body.Close()
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		h.t.Fatalf("reading body: %v", err)
	}
	return string(raw)
}

// signIn runs first-run setup, which leaves the client holding a session.
func (h *harness) signIn() {
	h.t.Helper()
	h.get("/setup").Body.Close()
	res := h.post("/setup", url.Values{
		"name":     {"Test Admin"},
		"email":    {"admin@example.com"},
		"password": {testPassword},
	})
	res.Body.Close()
	if res.StatusCode != http.StatusSeeOther {
		h.t.Fatalf("POST /setup returned %d, want 303", res.StatusCode)
	}
	if !h.hasSession() {
		h.t.Fatal("setup did not leave a session cookie")
	}
}

var idInPath = regexp.MustCompile(`/documents/(\d+)`)

// createdDocumentID reads the new document's id out of a 303 Location.
func (h *harness) createdDocumentID(res *http.Response) int64 {
	h.t.Helper()
	defer res.Body.Close()
	if res.StatusCode != http.StatusSeeOther {
		h.t.Fatalf("expected a 303, got %d", res.StatusCode)
	}
	location := res.Header.Get("Location")
	match := idInPath.FindStringSubmatch(location)
	if match == nil {
		h.t.Fatalf("no document id in Location %q", location)
	}
	id, err := strconv.ParseInt(match[1], 10, 64)
	if err != nil {
		h.t.Fatalf("parsing id from %q: %v", location, err)
	}
	return id
}

// privateRoutes is every route behind the auth gate.
//
// Adding a route to registerPrivateRoutes means adding it here. The catch-all
// case below covers anything that is missed, but an explicit list is what makes
// a regression on a specific endpoint legible.
func privateRoutes() []struct {
	method string
	path   string
} {
	return []struct {
		method string
		path   string
	}{
		{http.MethodGet, "/"},
		{http.MethodPost, "/logout"},
		{http.MethodGet, "/customers"},
		{http.MethodGet, "/customers/new"},
		{http.MethodPost, "/customers"},
		{http.MethodGet, "/customers/1"},
		{http.MethodGet, "/customers/1/edit"},
		{http.MethodPost, "/customers/1"},
		{http.MethodPost, "/customers/1/messages"},
		{http.MethodGet, "/documents"},
		{http.MethodGet, "/documents/new"},
		{http.MethodPost, "/documents"},
		{http.MethodGet, "/documents/1"},
		{http.MethodPost, "/documents/1/lines"},
		{http.MethodPost, "/documents/1/status"},
		{http.MethodPost, "/documents/1/convert"},
		{http.MethodPost, "/documents/1/payments"},
		{http.MethodGet, "/catalog"},
		{http.MethodPost, "/catalog/taxes"},
		{http.MethodPost, "/catalog/items"},
	}
}

func TestEveryPrivateRouteRefusesAnUnauthenticatedRequest(t *testing.T) {
	// A write with no session must come back 401 and a read must be redirected
	// to the sign-in page. Neither may return 200.
	//
	// This runs over every private route rather than a sample, because the
	// endpoint most likely to be left open is the one nobody thinks to check by
	// hand.
	h := newHarness(t)

	for _, route := range privateRoutes() {
		req, err := http.NewRequest(route.method, h.server.URL+route.path, nil)
		if err != nil {
			t.Fatalf("building %s %s: %v", route.method, route.path, err)
		}
		res, err := h.client.Do(req)
		if err != nil {
			t.Fatalf("%s %s: %v", route.method, route.path, err)
		}
		res.Body.Close()

		if route.method == http.MethodGet {
			if res.StatusCode != http.StatusSeeOther {
				t.Errorf("GET %s with no session returned %d, want 303 to /login", route.path, res.StatusCode)
				continue
			}
			if to := res.Header.Get("Location"); !strings.HasPrefix(to, "/login") {
				t.Errorf("GET %s with no session redirected to %q, want /login", route.path, to)
			}
			continue
		}
		if res.StatusCode != http.StatusUnauthorized {
			t.Errorf("%s %s with no session returned %d, want 401",
				route.method, route.path, res.StatusCode)
		}
	}
}

func TestUnknownPathsArePrivateByDefault(t *testing.T) {
	// The router denies by default: the public surface is enumerated and
	// everything else falls through the auth gate. A new route is therefore
	// protected unless someone deliberately makes it public, which is the
	// property that keeps the previous defect from coming back.
	h := newHarness(t)

	for _, path := range []string{"/anything", "/admin", "/customers/1/delete", "/api/v1/users"} {
		res := h.get(path)
		res.Body.Close()
		if res.StatusCode != http.StatusSeeOther {
			t.Errorf("GET %s returned %d, want a 303 to /login", path, res.StatusCode)
		}
	}
}

func TestPublicRoutesAreReachableWithoutASession(t *testing.T) {
	h := newHarness(t)

	// /login redirects to /setup while no account exists, which is itself a
	// public route, so both are checked for "not a redirect to /login".
	for _, path := range []string{"/healthz", "/static/app.css", "/setup"} {
		res := h.get(path)
		res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("GET %s returned %d, want 200", path, res.StatusCode)
		}
	}
}

func TestWritesRequireACSRFToken(t *testing.T) {
	h := newHarness(t)
	h.signIn()

	// Signed in, so these fail only on the missing token.
	for _, path := range []string{"/customers", "/catalog/taxes", "/documents"} {
		res := h.postWithoutCSRF(path, url.Values{"name": {"Anything"}})
		res.Body.Close()
		if res.StatusCode != http.StatusForbidden {
			t.Errorf("POST %s without a CSRF token returned %d, want 403", path, res.StatusCode)
		}
	}
}

func TestLoginRejectsTheWrongPasswordAndAcceptsTheRight(t *testing.T) {
	h := newHarness(t)
	h.signIn()

	// Start from a clean jar so the sign-in is genuinely exercised.
	fresh := newHarnessSharingStore(t, h.store)

	fresh.get("/login").Body.Close()
	res := fresh.post("/login", url.Values{
		"email":    {"admin@example.com"},
		"password": {"not the password"},
	})
	page := fresh.body(res)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("a failed sign-in returned %d, want 200 with the form again", res.StatusCode)
	}
	if !strings.Contains(page, "do not match an account") {
		t.Error("a failed sign-in did not show the rejection message")
	}
	if fresh.hasSession() {
		t.Fatal("a failed sign-in issued a session")
	}

	res = fresh.post("/login", url.Values{
		"email":    {"admin@example.com"},
		"password": {testPassword},
	})
	res.Body.Close()
	if res.StatusCode != http.StatusSeeOther {
		t.Fatalf("a successful sign-in returned %d, want 303", res.StatusCode)
	}
	if !fresh.hasSession() {
		t.Fatal("a successful sign-in did not issue a session")
	}
}

// newHarnessSharingStore gives a second client against the same database, for
// tests that need a different cookie jar.
func newHarnessSharingStore(t *testing.T, db *store.Store) *harness {
	t.Helper()
	srv, err := New(testSettings(), db)
	if err != nil {
		t.Fatalf("web.New: %v", err)
	}
	ts := httptest.NewServer(srv.Routes())
	t.Cleanup(ts.Close)

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatalf("cookiejar.New: %v", err)
	}
	base, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("parsing URL: %v", err)
	}
	return &harness{
		t: t, store: db, server: ts, jar: jar, base: base,
		client: &http.Client{
			Jar:           jar,
			CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
		},
	}
}

func TestSetupClosesOnceAnAccountExists(t *testing.T) {
	h := newHarness(t)
	h.signIn()

	res := h.get("/setup")
	res.Body.Close()
	if res.StatusCode != http.StatusSeeOther {
		t.Errorf("GET /setup with an account present returned %d, want a 303 to /login", res.StatusCode)
	}

	// And it cannot be used to mint a second administrator.
	res = h.post("/setup", url.Values{
		"name":     {"Second Admin"},
		"email":    {"second@example.com"},
		"password": {testPassword},
	})
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Errorf("POST /setup with an account present returned %d, want 403", res.StatusCode)
	}
}

func TestSecurityHeadersAreSet(t *testing.T) {
	h := newHarness(t)
	res := h.get("/healthz")
	res.Body.Close()

	want := map[string]string{
		"Content-Security-Policy": "default-src 'self'",
		"X-Content-Type-Options":  "nosniff",
		"X-Frame-Options":         "DENY",
		"Referrer-Policy":         "same-origin",
	}
	for header, expected := range want {
		if got := res.Header.Get(header); !strings.Contains(got, expected) {
			t.Errorf("%s = %q, want it to contain %q", header, got, expected)
		}
	}
}

func TestSafeNextRefusesAnOffsiteRedirect(t *testing.T) {
	// A next parameter is an open redirect if it is not constrained: the user
	// has just typed a password into a page that really was this application,
	// and is then sent somewhere else wearing a fresh session.
	offsite := []string{
		"https://evil.example/steal",
		"//evil.example",
		`/\evil.example`,
		"http://evil.example",
		"javascript:alert(1)",
		"/path\r\nSet-Cookie: x=y",
		"",
	}
	for _, candidate := range offsite {
		if got := safeNext(candidate); got != "/" {
			t.Errorf("safeNext(%q) = %q, want \"/\"", candidate, got)
		}
	}

	onsite := []string{"/customers", "/documents/4", "/catalog?x=1"}
	for _, candidate := range onsite {
		if got := safeNext(candidate); got != candidate {
			t.Errorf("safeNext(%q) = %q, want it unchanged", candidate, got)
		}
	}
}

func TestProposalThroughToAPaidInvoiceOverHTTP(t *testing.T) {
	// The advertised product, exercised the way a user would: sign in, set up a
	// tax, add a customer, raise a proposal, walk it to an invoice and pay it.
	h := newHarness(t)
	h.signIn()

	res := h.post("/catalog/taxes", url.Values{"name": {"VAT"}, "rate": {"20"}})
	res.Body.Close()
	if res.StatusCode != http.StatusSeeOther {
		t.Fatalf("creating a tax returned %d", res.StatusCode)
	}
	taxes, err := h.store.ListTaxes(t.Context())
	if err != nil || len(taxes) != 1 {
		t.Fatalf("ListTaxes = %d taxes, %v; want 1", len(taxes), err)
	}
	if taxes[0].Rate.String() != "20%" {
		t.Fatalf("tax rate stored as %s, want 20%%", taxes[0].Rate)
	}

	res = h.post("/customers", url.Values{
		"name":     {"Northwind Trading"},
		"status":   {"lead"},
		"currency": {"USD"},
		"email":    {"ap@northwind.example"},
	})
	res.Body.Close()
	if res.StatusCode != http.StatusSeeOther {
		t.Fatalf("creating a customer returned %d", res.StatusCode)
	}
	customers, err := h.store.ListCustomers(t.Context(), "")
	if err != nil || len(customers) != 1 {
		t.Fatalf("ListCustomers = %d, %v; want 1", len(customers), err)
	}
	customerID := customers[0].ID

	// A proposal for 10 days at 1,500.00 plus 20% VAT: 15,000.00 net,
	// 3,000.00 tax, 18,000.00 total.
	proposalID := h.createdDocumentID(h.post("/documents", url.Values{
		"type":             {"proposal"},
		"customer_id":      {strconv.FormatInt(customerID, 10)},
		"currency":         {"USD"},
		"issue_date":       {"2026-10-21"},
		"line_description": {"Implementation", ""},
		"line_quantity":    {"10", "1"},
		"line_unit_price":  {"1500.00", ""},
		"line_tax_id":      {strconv.FormatInt(taxes[0].ID, 10), ""},
	}))

	proposal, err := h.store.DocumentByID(t.Context(), proposalID)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}
	if len(proposal.Lines) != 1 {
		t.Fatalf("proposal has %d lines, want 1; the blank row should have been skipped", len(proposal.Lines))
	}
	if proposal.Totals.Total.String() != "18000.00" {
		t.Fatalf("proposal total is %s, want 18000.00", proposal.Totals.Total)
	}

	// Walk proposal to estimate, then estimate to invoice.
	invoiceID := walkToConversion(h, proposalID)
	invoiceID = walkToConversion(h, invoiceID)

	invoice, err := h.store.DocumentByID(t.Context(), invoiceID)
	if err != nil {
		t.Fatalf("DocumentByID invoice: %v", err)
	}
	if invoice.Type != "invoice" {
		t.Fatalf("the chain ended at a %s, want an invoice", invoice.Type)
	}
	if invoice.DueDate == nil {
		t.Fatal("the invoice has no due date, so the configured payment terms were not applied")
	}
	if invoice.Totals.Total != proposal.Totals.Total {
		t.Fatalf("the total changed across the chain: %s became %s",
			proposal.Totals.Total, invoice.Totals.Total)
	}

	// A payment is refused until the invoice has been sent.
	res = h.post("/documents/"+strconv.FormatInt(invoiceID, 10)+"/payments", url.Values{
		"amount":  {"18000.00"},
		"paid_on": {"2026-10-21"},
	})
	res.Body.Close()
	if to := res.Header.Get("Location"); !strings.Contains(to, "error=") {
		t.Fatalf("paying a draft invoice redirected to %q, want an error banner", to)
	}

	res = h.post("/documents/"+strconv.FormatInt(invoiceID, 10)+"/status", url.Values{"status": {"sent"}})
	res.Body.Close()

	// Part payment, then the rest.
	res = h.post("/documents/"+strconv.FormatInt(invoiceID, 10)+"/payments", url.Values{
		"amount":  {"8000.00"},
		"paid_on": {"2026-10-21"},
		"method":  {"Bank transfer"},
	})
	res.Body.Close()
	partial, err := h.store.DocumentByID(t.Context(), invoiceID)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}
	if partial.Status != "partially_paid" {
		t.Fatalf("after a part payment the invoice is %s, want partially_paid", partial.Status)
	}
	if partial.Outstanding().String() != "10000.00" {
		t.Fatalf("outstanding is %s, want 10000.00", partial.Outstanding())
	}

	res = h.post("/documents/"+strconv.FormatInt(invoiceID, 10)+"/payments", url.Values{
		"amount":  {"10000.00"},
		"paid_on": {"2026-10-22"},
	})
	res.Body.Close()
	paid, err := h.store.DocumentByID(t.Context(), invoiceID)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}
	if paid.Status != "paid" {
		t.Fatalf("after full payment the invoice is %s, want paid", paid.Status)
	}

	// And the screen says so.
	page := h.body(h.get("/documents/" + strconv.FormatInt(invoiceID, 10)))
	if !strings.Contains(page, "Paid") {
		t.Error("the invoice page does not show that it is paid")
	}
	if !strings.Contains(page, "18000.00") {
		t.Error("the invoice page does not show the total")
	}
}

// walkToConversion sends a document, accepts it, converts it, and returns the
// id of what it became.
func walkToConversion(h *harness, id int64) int64 {
	h.t.Helper()
	path := "/documents/" + strconv.FormatInt(id, 10)
	for _, status := range []string{"sent", "accepted"} {
		res := h.post(path+"/status", url.Values{"status": {status}})
		res.Body.Close()
		if res.StatusCode != http.StatusSeeOther {
			h.t.Fatalf("marking document %d as %s returned %d", id, status, res.StatusCode)
		}
	}
	return h.createdDocumentID(h.post(path+"/convert", nil))
}

func TestOverdueInvoicesSurfaceOnTheDashboardAndTheList(t *testing.T) {
	h := newHarness(t)
	h.signIn()

	res := h.post("/customers", url.Values{
		"name": {"Northwind Trading"}, "status": {"active"}, "currency": {"USD"},
	})
	res.Body.Close()
	customers, err := h.store.ListCustomers(t.Context(), "")
	if err != nil || len(customers) == 0 {
		t.Fatalf("ListCustomers = %d, %v", len(customers), err)
	}
	customerID := strconv.FormatInt(customers[0].ID, 10)

	// Fixed dates rather than offsets from today: 2020 is in the past and 2099
	// is not, whenever this suite happens to run.
	raise := func(issue, due string) int64 {
		t.Helper()
		return h.createdDocumentID(h.post("/documents", url.Values{
			"type":             {"invoice"},
			"customer_id":      {customerID},
			"currency":         {"USD"},
			"issue_date":       {issue},
			"due_date":         {due},
			"line_description": {"Consulting"},
			"line_quantity":    {"1"},
			"line_unit_price":  {"400.00"},
		}))
	}
	late := raise("2020-01-01", "2020-01-31")
	notYetDue := raise("2020-01-01", "2099-12-31")

	// A draft is not overdue however old, so both are sent.
	for _, id := range []int64{late, notYetDue} {
		res := h.post("/documents/"+strconv.FormatInt(id, 10)+"/status", url.Values{"status": {"sent"}})
		res.Body.Close()
	}

	lateDoc, err := h.store.DocumentByID(t.Context(), late)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}
	notYetDueDoc, err := h.store.DocumentByID(t.Context(), notYetDue)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}

	dashboard := h.body(h.get("/"))
	if !strings.Contains(dashboard, "Overdue") {
		t.Error("the dashboard does not mention Overdue despite an invoice being late")
	}
	if !strings.Contains(dashboard, lateDoc.Number) {
		t.Errorf("the dashboard does not show the overdue invoice %s", lateDoc.Number)
	}

	// The filter shows the late one and withholds the one not yet due.
	filtered := h.body(h.get("/documents?overdue=1"))
	if !strings.Contains(filtered, lateDoc.Number) {
		t.Errorf("the overdue filter omits %s, which is late", lateDoc.Number)
	}
	if strings.Contains(filtered, notYetDueDoc.Number) {
		t.Errorf("the overdue filter includes %s, which is due in 2099", notYetDueDoc.Number)
	}

	// Paying it takes it off both, with no sweep or scheduled job involved.
	res = h.post("/documents/"+strconv.FormatInt(late, 10)+"/payments", url.Values{
		"amount":  {lateDoc.Totals.Total.String()},
		"paid_on": {"2026-10-21"},
	})
	res.Body.Close()

	settled, err := h.store.DocumentByID(t.Context(), late)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}
	if settled.Status != "paid" {
		t.Fatalf("invoice is %s after full payment, want paid", settled.Status)
	}
	if again := h.body(h.get("/documents?overdue=1")); strings.Contains(again, settled.Number) {
		t.Errorf("%s is still listed as overdue after being paid in full", settled.Number)
	}
}

func TestIllegalStatusMoveIsRefusedOverHTTP(t *testing.T) {
	h := newHarness(t)
	h.signIn()

	res := h.post("/catalog/taxes", url.Values{"name": {"VAT"}, "rate": {"20"}})
	res.Body.Close()
	res = h.post("/customers", url.Values{
		"name": {"Acme"}, "status": {"active"}, "currency": {"USD"},
	})
	res.Body.Close()
	customers, _ := h.store.ListCustomers(t.Context(), "")

	id := h.createdDocumentID(h.post("/documents", url.Values{
		"type":             {"invoice"},
		"customer_id":      {strconv.FormatInt(customers[0].ID, 10)},
		"currency":         {"USD"},
		"issue_date":       {"2026-10-21"},
		"line_description": {"Thing"},
		"line_quantity":    {"1"},
		"line_unit_price":  {"100.00"},
	}))

	// Draft straight to paid skips sending entirely.
	res = h.post("/documents/"+strconv.FormatInt(id, 10)+"/status", url.Values{"status": {"paid"}})
	res.Body.Close()
	if to := res.Header.Get("Location"); !strings.Contains(to, "error=") {
		t.Fatalf("draft to paid redirected to %q, want an error banner", to)
	}
	doc, err := h.store.DocumentByID(t.Context(), id)
	if err != nil {
		t.Fatalf("DocumentByID: %v", err)
	}
	if doc.Status != "draft" {
		t.Fatalf("the refused move still changed the status to %s", doc.Status)
	}
}
