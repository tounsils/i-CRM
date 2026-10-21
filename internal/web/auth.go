package web

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/tounsils/i-CRM/internal/store"
)

// minPasswordLength is 12 characters. Rationale rather than a round number:
// this is a self-hosted application whose sign-in page is frequently exposed to
// the internet by an operator with no rate limiting in front of it, so the
// password is the only barrier, and 12 characters is the shortest length that
// stays out of reach of offline guessing against a bcrypt hash.
const minPasswordLength = 12

// bcryptCost is left at the library default of 10, which is roughly 50-100ms
// per verification on current hardware. That is slow enough to make offline
// guessing expensive and fast enough that a sign-in does not feel broken. It is
// stated here so that raising it is a deliberate decision with a measurement
// behind it rather than a guess.
const bcryptCost = bcrypt.DefaultCost

// newEnumerationGuard returns a bcrypt hash of random bytes, generated once per
// process. When a sign-in names an address that does not exist, the password is
// verified against this instead of returning early, so an attempt costs the
// same whether or not the account exists and the page cannot be used to
// discover who has an account.
//
// It is generated rather than written here as a constant so that no
// credential-shaped literal sits in the source at all.
func newEnumerationGuard() ([]byte, error) {
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		return nil, fmt.Errorf("web: seeding the enumeration guard: %w", err)
	}
	hash, err := bcrypt.GenerateFromPassword(secret, bcryptCost)
	if err != nil {
		return nil, fmt.Errorf("web: building the enumeration guard: %w", err)
	}
	return hash, nil
}

func (s *Server) handleLoginForm(w http.ResponseWriter, r *http.Request) {
	count, err := s.store.CountUsers(r.Context())
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	if count == 0 {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}
	if _, ok := s.sessionUserID(r); ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	s.render(w, r, "login.html", pageData{
		Title: "Sign in",
		Data:  map[string]string{"Next": safeNext(r.URL.Query().Get("next"))},
	})
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	if err := s.checkCSRF(r); err != nil {
		s.fail(w, r, http.StatusForbidden, err)
		return
	}

	email := strings.TrimSpace(r.PostFormValue("email"))
	password := r.PostFormValue("password")
	next := safeNext(r.PostFormValue("next"))

	// Both keys are checked and both are counted, so one address cannot work
	// through a list of accounts while staying under each account's own limit.
	keys := []string{accountKey(email), sourceKey(r)}

	if wait, locked := s.throttle.locked(keys...); locked {
		// Deliberately before the password is read against anything. The
		// message is identical whether or not the address exists, because the
		// counter is keyed on what was submitted rather than on what is
		// registered.
		s.render(w, r, "login.html", pageData{
			Title: "Sign in",
			Error: "Too many sign-in attempts. Try again in " + humanWait(wait) + ".",
			Data:  map[string]string{"Next": next, "Email": email},
		})
		return
	}

	user, err := s.store.UserByEmail(r.Context(), email)
	if err != nil && !errors.Is(err, store.ErrNotFound) {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	hash := []byte(user.PasswordHash)
	if errors.Is(err, store.ErrNotFound) {
		hash = s.enumerationGuard
	}

	if bcrypt.CompareHashAndPassword(hash, []byte(password)) != nil || user.ID == 0 {
		s.throttle.recordFailure(keys...)

		// One message for both a wrong password and an unknown address, unless
		// this was the attempt that tripped the lock, in which case say so now.
		// Reporting "wrong password" on the attempt that locks the account and
		// only mentioning the lock on the next one makes the lock look
		// arbitrary.
		problem := "Those details do not match an account."
		if wait, locked := s.throttle.locked(keys...); locked {
			problem = "Too many sign-in attempts. Try again in " + humanWait(wait) + "."
		}

		s.render(w, r, "login.html", pageData{
			Title: "Sign in",
			Error: problem,
			Data:  map[string]string{"Next": next, "Email": email},
		})
		return
	}

	// A user who eventually remembers their password should not still be
	// counting towards a lock.
	s.throttle.reset(keys...)
	s.issueSession(w, user.ID)
	http.Redirect(w, r, next, http.StatusSeeOther)
}

func (s *Server) handleLogout(w http.ResponseWriter, r *http.Request) {
	s.clearSession(w)
	redirect(w, r, "/login", "You have been signed out.")
}

// handleSetupForm offers to create the first administrator, and only ever when
// there are no users at all.
func (s *Server) handleSetupForm(w http.ResponseWriter, r *http.Request) {
	count, err := s.store.CountUsers(r.Context())
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	if count > 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	// Data is an empty map rather than nil because the template repopulates the
	// fields with index, and index of a nil value is an error.
	s.render(w, r, "setup.html", pageData{
		Title: "Create the first administrator",
		Data:  map[string]string{},
	})
}

// handleSetup creates the first administrator.
//
// The zero-users check is the entire authorisation for this route: it is public
// because nobody can sign in before it runs, and it stops being reachable the
// moment it succeeds. The check is repeated here rather than trusted from the
// form, because the form is not where authorisation decisions belong.
func (s *Server) handleSetup(w http.ResponseWriter, r *http.Request) {
	if err := s.checkCSRF(r); err != nil {
		s.fail(w, r, http.StatusForbidden, err)
		return
	}
	count, err := s.store.CountUsers(r.Context())
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	if count > 0 {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	name := strings.TrimSpace(r.PostFormValue("name"))
	email := strings.TrimSpace(r.PostFormValue("email"))
	password := r.PostFormValue("password")

	reject := func(message string) {
		s.render(w, r, "setup.html", pageData{
			Title: "Create the first administrator",
			Error: message,
			Data:  map[string]string{"Name": name, "Email": email},
		})
	}
	switch {
	case name == "" || email == "":
		reject("A name and an email address are both required.")
		return
	case !strings.Contains(email, "@"):
		reject("That does not look like an email address.")
		return
	case len(password) < minPasswordLength:
		reject("The password must be at least 12 characters.")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	id, err := s.store.CreateUser(r.Context(), store.User{
		Email:        email,
		Name:         name,
		PasswordHash: string(hash),
		Role:         "admin",
	})
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}

	s.issueSession(w, id)
	redirect(w, r, "/", "Welcome. Add a tax rate and your first customer to get started.")
}

// safeNext sanitises a post-sign-in redirect target.
//
// Only a same-site absolute path is allowed. Without this the next parameter is
// an open redirect: a link to /login?next=https://evil.example would send the
// user there wearing a fresh session, having just typed their password into a
// page that genuinely was this application.
func safeNext(candidate string) string {
	const fallback = "/"
	if candidate == "" || !strings.HasPrefix(candidate, "/") {
		return fallback
	}
	// "//host" and "/\host" are both read as protocol-relative URLs by browsers.
	if strings.HasPrefix(candidate, "//") || strings.HasPrefix(candidate, `/\`) {
		return fallback
	}
	if strings.ContainsAny(candidate, "\r\n") {
		return fallback
	}
	return candidate
}
