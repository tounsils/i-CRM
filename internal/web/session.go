package web

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/tounsils/i-CRM/internal/store"
)

const (
	sessionCookieName = "icrm_session"

	// sessionTTL is one working day. Long enough that staff are not signed out
	// mid-task, short enough that a forgotten session on a shared machine
	// expires the same day rather than next week.
	sessionTTL = 12 * time.Hour
)

type contextKey int

const userContextKey contextKey = iota

// issueSession sets a signed cookie carrying the user id and an expiry.
//
// The cookie is signed rather than looked up in a server-side table, so a
// restart of the single binary does not sign everybody out and there is no
// session store to grow or prune. The signature is over both the id and the
// expiry, so neither can be edited without invalidating it.
func (s *Server) issueSession(w http.ResponseWriter, userID int64) {
	expires := time.Now().Add(sessionTTL)
	payload := fmt.Sprintf("%d:%d", userID, expires.Unix())
	signed := payload + "." + sign(s.cfg.SessionKey, payload)

	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    base64.RawURLEncoding.EncodeToString([]byte(signed)),
		Path:     "/",
		Expires:  expires,
		HttpOnly: true,
		// Secure is tied to the environment because a self-hosted install is
		// routinely reached over plain HTTP on a LAN during setup, where a
		// Secure cookie would silently never be sent.
		Secure:   s.cfg.Production,
		SameSite: http.SameSiteLaxMode,
	})
}

// clearSession expires the cookie.
func (s *Server) clearSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   s.cfg.Production,
		SameSite: http.SameSiteLaxMode,
	})
}

// sessionUserID returns the signed-in user id, or false if there is no valid
// unexpired session. Every failure path returns the same false rather than a
// reason, because the caller's response must not distinguish a tampered cookie
// from an absent one.
func (s *Server) sessionUserID(r *http.Request) (int64, bool) {
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		return 0, false
	}
	raw, err := base64.RawURLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return 0, false
	}

	value := string(raw)
	dot := strings.LastIndexByte(value, '.')
	if dot < 0 {
		return 0, false
	}
	payload, mac := value[:dot], value[dot+1:]
	if !hmac.Equal([]byte(mac), []byte(sign(s.cfg.SessionKey, payload))) {
		return 0, false
	}

	id, expires, ok := splitPayload(payload)
	if !ok || time.Now().After(time.Unix(expires, 0)) {
		return 0, false
	}
	return id, true
}

func splitPayload(payload string) (userID, expiresUnix int64, ok bool) {
	parts := strings.SplitN(payload, ":", 2)
	if len(parts) != 2 {
		return 0, 0, false
	}
	userID, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, false
	}
	expiresUnix, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, 0, false
	}
	return userID, expiresUnix, true
}

func sign(key []byte, payload string) string {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(payload))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func withUser(ctx context.Context, u store.User) context.Context {
	return context.WithValue(ctx, userContextKey, u)
}

// userFrom returns the signed-in user. It is only populated behind requireAuth,
// so a handler that finds it missing is a handler that was registered public by
// mistake.
func userFrom(ctx context.Context) (store.User, bool) {
	u, ok := ctx.Value(userContextKey).(store.User)
	return u, ok
}
