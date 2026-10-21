package web

import (
	"crypto/hmac"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
)

const (
	csrfCookieName = "icrm_csrf"
	csrfFieldName  = "csrf_token"

	// csrfTokenBytes gives a 256-bit token, matching the session signature.
	csrfTokenBytes = 32
)

var errCSRF = errors.New("csrf token missing or mismatched")

// ensureCSRFToken returns the request's token, minting and setting one if the
// browser does not have it yet.
//
// This is a double-submit token: the same value travels in a cookie and in a
// hidden form field, and a write is accepted only when they match. The primary
// defence is actually SameSite=Lax on both cookies, which stops a cross-site
// form post from carrying them at all; the double submit is what remains if a
// future browser or proxy weakens that.
func (s *Server) ensureCSRFToken(w http.ResponseWriter, r *http.Request) string {
	if cookie, err := r.Cookie(csrfCookieName); err == nil && cookie.Value != "" {
		return cookie.Value
	}

	buf := make([]byte, csrfTokenBytes)
	if _, err := rand.Read(buf); err != nil {
		// No token means no form on the page can be submitted. Log it rather
		// than silently rendering a page whose buttons all fail.
		log.Printf("web: generating a CSRF token: %v", err)
		return ""
	}
	token := base64.RawURLEncoding.EncodeToString(buf)

	http.SetCookie(w, &http.Cookie{
		Name:     csrfCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   s.cfg.Production,
		SameSite: http.SameSiteLaxMode,
	})
	return token
}

// checkCSRF verifies a write request. It is called centrally in requireAuth, so
// no private handler can forget it, and explicitly by the public login handler.
func (s *Server) checkCSRF(r *http.Request) error {
	cookie, err := r.Cookie(csrfCookieName)
	if err != nil || cookie.Value == "" {
		return errCSRF
	}
	submitted := r.PostFormValue(csrfFieldName)
	if submitted == "" || !hmac.Equal([]byte(cookie.Value), []byte(submitted)) {
		return errCSRF
	}
	return nil
}
