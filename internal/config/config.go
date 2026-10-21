// Package config is the single owner of every tunable setting. Defaults live
// here and nowhere else: no handler, template or store function carries its own
// opinion about a value this package publishes.
package config

import (
	"crypto/rand"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Settings is the resolved configuration for one process.
type Settings struct {
	Addr             string
	DatabasePath     string
	SessionKey       []byte
	PaymentTermsDays int
	CompanyName      string
	DefaultCurrency  string
	Production       bool

	// SessionKeyGenerated is true when no key was supplied and an ephemeral one
	// was made up, which means sessions do not survive a restart. main logs
	// this; it is a development convenience and an error in production.
	SessionKeyGenerated bool
}

// Defaults, stated once. Every one of these is overridable by the matching
// environment variable in Load.
const (
	defaultAddr             = ":8080"
	defaultDatabasePath     = "data/icrm.db"
	defaultPaymentTermsDays = 30
	defaultCompanyName      = "i-CRM"
	defaultCurrency         = "USD"

	// sessionKeyBytes is 32 because the session cookie is signed with
	// HMAC-SHA256, whose block-level security is 256 bits; a shorter key would
	// be the weakest part of the signature and a longer one buys nothing.
	sessionKeyBytes = 32
)

// ErrMissingSessionKey is returned when a production process has no signing key.
var ErrMissingSessionKey = errors.New("config: ICRM_SESSION_KEY is required when ICRM_ENV=production")

// Load resolves settings from the environment.
//
// Precedence is simply: environment variable if set and non-empty, otherwise the
// default above. There is no config file, because a self-hosted single binary
// that reads only the environment is one fewer thing for an operator to get
// wrong, and it matches how the container is run.
func Load() (Settings, error) {
	s := Settings{
		Addr:             envOr("ICRM_ADDR", defaultAddr),
		DatabasePath:     envOr("ICRM_DB", defaultDatabasePath),
		CompanyName:      envOr("ICRM_COMPANY_NAME", defaultCompanyName),
		DefaultCurrency:  strings.ToUpper(envOr("ICRM_CURRENCY", defaultCurrency)),
		PaymentTermsDays: defaultPaymentTermsDays,
		Production:       strings.EqualFold(os.Getenv("ICRM_ENV"), "production"),
	}

	if raw := os.Getenv("ICRM_PAYMENT_TERMS_DAYS"); raw != "" {
		days, err := strconv.Atoi(raw)
		if err != nil || days < 0 {
			return Settings{}, fmt.Errorf("config: ICRM_PAYMENT_TERMS_DAYS must be a non-negative whole number, got %q", raw)
		}
		s.PaymentTermsDays = days
	}

	key, generated, err := loadSessionKey(s.Production)
	if err != nil {
		return Settings{}, err
	}
	s.SessionKey = key
	s.SessionKeyGenerated = generated

	return s, nil
}

// loadSessionKey reads the signing key, or makes an ephemeral one outside
// production. Refusing to invent a key in production is deliberate: a generated
// key silently signs out every user on each deploy, and the failure looks like
// a session bug rather than a missing setting.
func loadSessionKey(production bool) (key []byte, generated bool, err error) {
	if raw := os.Getenv("ICRM_SESSION_KEY"); raw != "" {
		if len(raw) < sessionKeyBytes {
			return nil, false, fmt.Errorf("config: ICRM_SESSION_KEY must be at least %d characters", sessionKeyBytes)
		}
		return []byte(raw), false, nil
	}
	if production {
		return nil, false, ErrMissingSessionKey
	}
	buf := make([]byte, sessionKeyBytes)
	if _, err := rand.Read(buf); err != nil {
		return nil, false, fmt.Errorf("config: generating a development session key: %w", err)
	}
	return buf, true, nil
}

// DueDateFor applies the configured payment terms to an issue date. This is the
// only place payment terms are turned into a date, which is why store.Convert
// takes a due date rather than computing one.
func (s Settings) DueDateFor(issued time.Time) time.Time {
	return issued.AddDate(0, 0, s.PaymentTermsDays)
}

func envOr(name, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(name)); v != "" {
		return v
	}
	return fallback
}
