package web

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

// fixedClock lets a test move time without sleeping, which is the whole reason
// loginThrottle takes its clock as a field.
type fixedClock struct{ at time.Time }

func (c *fixedClock) now() time.Time          { return c.at }
func (c *fixedClock) advance(d time.Duration) { c.at = c.at.Add(d) }

func newTestThrottle() (*loginThrottle, *fixedClock) {
	clock := &fixedClock{at: time.Date(2026, time.October, 21, 9, 0, 0, 0, time.UTC)}
	t := newLoginThrottle()
	t.now = clock.now
	return t, clock
}

func TestThrottleAllowsAttemptsUpToTheLimit(t *testing.T) {
	throttle, _ := newTestThrottle()
	key := accountKey("someone@example.com")

	// One short of the limit must still be allowed: the limit is the point at
	// which it locks, not one before.
	for i := 1; i < maxFailures; i++ {
		throttle.recordFailure(key)
		if _, locked := throttle.locked(key); locked {
			t.Fatalf("locked after %d failures, the limit is %d", i, maxFailures)
		}
	}

	throttle.recordFailure(key)
	wait, locked := throttle.locked(key)
	if !locked {
		t.Fatalf("not locked after %d failures", maxFailures)
	}
	if wait <= 0 || wait > lockoutPeriod {
		t.Fatalf("wait is %s, want something in (0, %s]", wait, lockoutPeriod)
	}
}

func TestThrottleReleasesAfterTheLockout(t *testing.T) {
	throttle, clock := newTestThrottle()
	key := accountKey("someone@example.com")

	for i := 0; i < maxFailures; i++ {
		throttle.recordFailure(key)
	}
	if _, locked := throttle.locked(key); !locked {
		t.Fatal("not locked after reaching the limit")
	}

	// One second before release it is still locked.
	clock.advance(lockoutPeriod - time.Second)
	if _, locked := throttle.locked(key); !locked {
		t.Fatal("released a second early")
	}

	clock.advance(2 * time.Second)
	if _, locked := throttle.locked(key); locked {
		t.Fatal("still locked after the lockout period elapsed")
	}
}

func TestThrottleForgetsFailuresOnceTheWindowPasses(t *testing.T) {
	// Failures from an hour ago must not add to failures from now, or a user
	// who mistypes twice a day eventually locks themselves out.
	throttle, clock := newTestThrottle()
	key := accountKey("someone@example.com")

	for i := 0; i < maxFailures-1; i++ {
		throttle.recordFailure(key)
	}
	clock.advance(failureWindow + time.Second)

	for i := 0; i < maxFailures-1; i++ {
		throttle.recordFailure(key)
		if _, locked := throttle.locked(key); locked {
			t.Fatalf("locked after %d failures in a fresh window", i+1)
		}
	}
}

func TestThrottleResetClearsTheCount(t *testing.T) {
	throttle, _ := newTestThrottle()
	key := accountKey("someone@example.com")

	for i := 0; i < maxFailures-1; i++ {
		throttle.recordFailure(key)
	}
	throttle.reset(key)

	// After a success the count starts again, so the next mistake is the first.
	for i := 1; i < maxFailures; i++ {
		throttle.recordFailure(key)
		if _, locked := throttle.locked(key); locked {
			t.Fatalf("locked after %d failures following a reset", i)
		}
	}
}

func TestThrottleTreatsUnknownAndKnownAddressesIdentically(t *testing.T) {
	// The counter is keyed on what was submitted, not on what is registered,
	// so a lockout cannot be used to discover which addresses have accounts.
	throttle, _ := newTestThrottle()
	known := accountKey("admin@example.com")
	unknown := accountKey("nobody@example.com")

	for i := 0; i < maxFailures; i++ {
		throttle.recordFailure(known)
		throttle.recordFailure(unknown)
	}

	knownWait, knownLocked := throttle.locked(known)
	unknownWait, unknownLocked := throttle.locked(unknown)
	if knownLocked != unknownLocked {
		t.Fatalf("known locked=%v but unknown locked=%v", knownLocked, unknownLocked)
	}
	if knownWait != unknownWait {
		t.Fatalf("known waits %s but unknown waits %s", knownWait, unknownWait)
	}
}

func TestThrottleLocksASourceSprayingManyAccounts(t *testing.T) {
	// Each account sees one failure, which is under its own limit, but the
	// source exhausts its allowance. Without this, a list of addresses defeats
	// a per-account limit entirely.
	throttle, _ := newTestThrottle()
	source := "source:198.51.100.7"

	for i := 0; i < maxFailures; i++ {
		account := accountKey("user" + strconv.Itoa(i) + "@example.com")
		throttle.recordFailure(account, source)
		if _, locked := throttle.locked(account); locked {
			t.Fatalf("account %d locked after a single failure", i)
		}
	}

	if _, locked := throttle.locked(source); !locked {
		t.Fatal("the source is not locked after spraying the whole allowance")
	}
}

func TestThrottleDoesNotGrowWithoutBound(t *testing.T) {
	// An attacker inventing addresses must not be able to grow this map until
	// the process dies.
	throttle, clock := newTestThrottle()

	for i := 0; i < maxThrottleKeys+500; i++ {
		throttle.recordFailure(accountKey("invented" + strconv.Itoa(i) + "@example.com"))
	}
	if size := throttle.size(); size > maxThrottleKeys {
		t.Fatalf("the table holds %d keys, above the cap of %d", size, maxThrottleKeys)
	}

	// And expired entries are released rather than held forever.
	clock.advance(failureWindow + lockoutPeriod + time.Second)
	throttle.recordFailure(accountKey("one-more@example.com"))
	if size := throttle.size(); size > 2 {
		t.Fatalf("the table still holds %d keys after everything expired", size)
	}
}

func TestThrottleFailsOpenRatherThanLockingEveryoneOut(t *testing.T) {
	// Filling the table must not lock out an account that was never attempted.
	// Failing closed here would turn a brute-force defence into a trivial
	// denial of service against every user.
	throttle, _ := newTestThrottle()
	for i := 0; i < maxThrottleKeys+100; i++ {
		throttle.recordFailure(accountKey("invented" + strconv.Itoa(i) + "@example.com"))
	}

	if _, locked := throttle.locked(accountKey("innocent@example.com")); locked {
		t.Fatal("an untouched account is locked because the table filled up")
	}
}

func TestClientIPIgnoresAForwardedForHeader(t *testing.T) {
	// The header is attacker-controlled unless a trusted proxy overwrites it,
	// so believing it would let one source defeat the per-source limit by
	// varying a header.
	r := httptest.NewRequest(http.MethodPost, "/login", nil)
	r.RemoteAddr = "198.51.100.7:54321"
	r.Header.Set("X-Forwarded-For", "203.0.113.1")

	if got := clientIP(r); got != "198.51.100.7" {
		t.Fatalf("clientIP = %q, want the connection address 198.51.100.7", got)
	}
}

func TestHumanWaitRoundsUp(t *testing.T) {
	cases := []struct {
		in   time.Duration
		want string
	}{
		{time.Second, "a minute"},
		{time.Minute, "a minute"},
		{time.Minute + time.Second, "2 minutes"},
		{14*time.Minute + 3*time.Second, "15 minutes"},
		{15 * time.Minute, "15 minutes"},
	}
	for _, c := range cases {
		if got := humanWait(c.in); got != c.want {
			t.Errorf("humanWait(%s) = %q, want %q", c.in, got, c.want)
		}
	}
}
