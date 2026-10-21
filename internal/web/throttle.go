package web

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Limits on repeated failed sign-ins.
//
// Derivation rather than round numbers. bcrypt at the default cost takes
// roughly 60ms per verification, so without a limit a single connection manages
// something like 50,000 guesses an hour. These figures cut that to 32 an hour,
// about 280,000 a year, which is hopeless against the 12-character minimum this
// application enforces.
//
// maxFailures is 8 because a person who has mistyped a password, then tried it
// with caps lock on, then tried their other password, is at five or six
// attempts and has done nothing wrong. Eight leaves room for that without a
// support call.
//
// The window and the lock are the same length so the behaviour is easy to state
// to a user: eight wrong attempts in a quarter of an hour, and you wait a
// quarter of an hour.
const (
	maxFailures     = 8
	failureWindow   = 15 * time.Minute
	lockoutPeriod   = 15 * time.Minute
	maxThrottleKeys = 10000
)

// loginThrottle counts recent failed sign-ins per key and locks a key that
// exceeds the limit.
//
// State is in process and is lost on restart. For a single binary serving one
// business that is the right trade: the alternative is a table to write, prune
// and migrate, to defend against an attacker who would have to also be timing
// their guessing around deploys.
type loginThrottle struct {
	mu      sync.Mutex
	entries map[string]*throttleEntry

	// now is injectable so the window and the lock can be tested without
	// sleeping through a quarter of an hour.
	now func() time.Time
}

type throttleEntry struct {
	failures    int
	windowStart time.Time
	lockedUntil time.Time
}

func newLoginThrottle() *loginThrottle {
	return &loginThrottle{
		entries: make(map[string]*throttleEntry),
		now:     time.Now,
	}
}

// retryAfter reports how long a key must wait, and whether it is locked at all.
func (t *loginThrottle) retryAfter(key string) (time.Duration, bool) {
	t.mu.Lock()
	defer t.mu.Unlock()

	entry, ok := t.entries[key]
	if !ok {
		return 0, false
	}
	now := t.now()
	if now.Before(entry.lockedUntil) {
		return entry.lockedUntil.Sub(now), true
	}
	return 0, false
}

// locked reports whether any of the supplied keys is locked, and for how long.
// A sign-in is refused if either the account or the source address has
// exhausted its attempts.
func (t *loginThrottle) locked(keys ...string) (time.Duration, bool) {
	var longest time.Duration
	anyLocked := false
	for _, key := range keys {
		if wait, isLocked := t.retryAfter(key); isLocked {
			anyLocked = true
			if wait > longest {
				longest = wait
			}
		}
	}
	return longest, anyLocked
}

// recordFailure counts a failed attempt against every supplied key.
func (t *loginThrottle) recordFailure(keys ...string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	now := t.now()
	t.pruneLocked(now)

	for _, key := range keys {
		entry, ok := t.entries[key]
		if !ok {
			// Once the table is full, a new key is not tracked rather than
			// evicting a live one. This fails open deliberately: failing closed
			// would let anyone lock every account out by filling the table with
			// invented addresses, turning a brute-force defence into a denial
			// of service. The per-account keys already present keep working.
			if len(t.entries) >= maxThrottleKeys {
				continue
			}
			entry = &throttleEntry{windowStart: now}
			t.entries[key] = entry
		}

		// A window that has expired starts again rather than accumulating
		// failures from last week into today's count.
		if now.Sub(entry.windowStart) > failureWindow {
			entry.failures = 0
			entry.windowStart = now
		}

		entry.failures++
		if entry.failures >= maxFailures {
			entry.lockedUntil = now.Add(lockoutPeriod)
			entry.failures = 0
			entry.windowStart = now
		}
	}
}

// reset clears the counters for the supplied keys, which happens on a
// successful sign-in so a user who eventually remembers is not still counting
// towards a lock.
func (t *loginThrottle) reset(keys ...string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, key := range keys {
		delete(t.entries, key)
	}
}

// size is how many keys are currently tracked. It exists so the bound on this
// table is asserted by a test rather than assumed.
func (t *loginThrottle) size() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return len(t.entries)
}

// pruneLocked drops entries that are neither locked nor inside their window.
// Called while the mutex is held.
func (t *loginThrottle) pruneLocked(now time.Time) {
	for key, entry := range t.entries {
		if now.Before(entry.lockedUntil) {
			continue
		}
		if now.Sub(entry.windowStart) <= failureWindow {
			continue
		}
		delete(t.entries, key)
	}
}

// humanWait rounds a wait up to whole minutes. Telling someone to wait
// "14 minutes 3 seconds" invites a retry at fourteen minutes that then fails.
func humanWait(d time.Duration) string {
	minutes := int((d + time.Minute - 1) / time.Minute)
	if minutes <= 1 {
		return "a minute"
	}
	return fmt.Sprintf("%d minutes", minutes)
}

// accountKey identifies the account a sign-in was attempted against.
//
// It is built from the submitted address whether or not an account exists, so
// an unknown address is throttled exactly like a known one. Keying only real
// accounts would turn the lockout into an oracle for which addresses are
// registered.
func accountKey(email string) string {
	return "account:" + strings.ToLower(strings.TrimSpace(email))
}

// sourceKey identifies where a sign-in came from, so one address cannot work
// through a list of accounts while staying under each account's own limit.
func sourceKey(r *http.Request) string {
	return "source:" + clientIP(r)
}

// clientIP is the address the connection came from.
//
// X-Forwarded-For is deliberately not consulted. The header is set by the
// client unless a trusted proxy overwrites it, so believing it would let an
// attacker defeat the per-source limit by varying one header. The cost is that
// behind a reverse proxy every request appears to come from the proxy, and the
// per-source limit then applies to the proxy as a whole; this is why the
// per-account limit above is the one that matters in that deployment.
func clientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
