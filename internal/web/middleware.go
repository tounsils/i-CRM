package web

import (
	"errors"
	"log"
	"net/http"
	"net/url"
)

// requireAuth is the gate every private route sits behind.
//
// It does three things in one place so that none of them can be forgotten per
// route: it rejects requests without a valid session, it loads the user for the
// handler, and it enforces CSRF on every method that is not a read. A gate in
// middleware covers the routes nobody thought to audit, which is the one thing
// a check written per handler cannot do.
func (s *Server) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := s.sessionUserID(r)
		if !ok {
			s.denyUnauthenticated(w, r)
			return
		}

		user, err := s.store.UserByID(r.Context(), userID)
		if err != nil {
			// A correctly signed cookie for a user who has since been deleted.
			// Clear it so the browser stops presenting it.
			s.clearSession(w)
			s.denyUnauthenticated(w, r)
			return
		}

		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			if err := s.checkCSRF(r); err != nil {
				s.fail(w, r, http.StatusForbidden, err)
				return
			}
		}

		next.ServeHTTP(w, r.WithContext(withUser(r.Context(), user)))
	})
}

// denyUnauthenticated refuses a request that has no session.
//
// A read is redirected to the sign-in page carrying where it was going. A write
// gets 401 rather than a redirect, because a redirect to an HTML page is a 200
// as far as a non-browser client is concerned, and that client needs to see
// that its write was refused.
func (s *Server) denyUnauthenticated(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet || r.Method == http.MethodHead {
		to := "/login?next=" + url.QueryEscape(r.URL.RequestURI())
		http.Redirect(w, r, to, http.StatusSeeOther)
		return
	}
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

// secureHeaders sets the response headers that do not depend on the route.
//
// The content security policy can be this strict because the application ships
// no JavaScript and loads nothing from another origin: every asset is embedded
// in the binary. A self-hosted install therefore works with no internet access
// at all, which is part of what self-hostable has to mean.
func (s *Server) secureHeaders(next http.Handler) http.Handler {
	const policy = "default-src 'self'; form-action 'self'; frame-ancestors 'none'; base-uri 'none'"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("Content-Security-Policy", policy)
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("X-Frame-Options", "DENY")
		h.Set("Referrer-Policy", "same-origin")
		next.ServeHTTP(w, r)
	})
}

// recoverPanic keeps one bad request from taking the process down, which on a
// self-hosted single binary would take the whole CRM with it.
func (s *Server) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recovered := recover(); recovered != nil {
				// http.ErrAbortHandler is how a handler asks to abort without
				// a log entry, so it is re-panicked rather than swallowed.
				if err, ok := recovered.(error); ok && errors.Is(err, http.ErrAbortHandler) {
					panic(recovered)
				}
				log.Printf("web: panic serving %s %s: %v", r.Method, r.URL.Path, recovered)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// maxRequestBody caps a submission at 1 MiB.
//
// Derived from the largest legitimate form: a document with a hundred lines
// submits roughly 20 KB. A megabyte leaves two orders of magnitude of headroom
// while stopping ParseForm from reading an unbounded body into memory, which on
// a single-binary deployment is the whole process's memory.
const maxRequestBody = 1 << 20

func (s *Server) limitBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			r.Body = http.MaxBytesReader(w, r.Body, maxRequestBody)
		}
		next.ServeHTTP(w, r)
	})
}

// handleStylesheet serves the one embedded asset. It is public because the
// sign-in page needs it before a session exists.
func (s *Server) handleStylesheet(w http.ResponseWriter, r *http.Request) {
	body, err := templateFS.ReadFile("templates/app.css")
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	if _, err := w.Write(body); err != nil {
		log.Printf("web: writing stylesheet: %v", err)
	}
}
