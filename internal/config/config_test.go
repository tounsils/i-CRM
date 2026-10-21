package config

import (
	"errors"
	"testing"
	"time"
)

// clearEnv blanks every variable Load reads, so a test starts from a known
// state regardless of what the developer has exported in their shell.
func clearEnv(t *testing.T) {
	t.Helper()
	for _, name := range []string{
		"ICRM_ADDR", "ICRM_DB", "ICRM_SESSION_KEY", "ICRM_ENV",
		"ICRM_COMPANY_NAME", "ICRM_CURRENCY", "ICRM_PAYMENT_TERMS_DAYS",
	} {
		t.Setenv(name, "")
	}
}

func TestLoadFallsBackToTheDocumentedDefaults(t *testing.T) {
	clearEnv(t)

	got, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if got.Addr != defaultAddr {
		t.Errorf("Addr = %q, want %q", got.Addr, defaultAddr)
	}
	if got.DatabasePath != defaultDatabasePath {
		t.Errorf("DatabasePath = %q, want %q", got.DatabasePath, defaultDatabasePath)
	}
	if got.PaymentTermsDays != defaultPaymentTermsDays {
		t.Errorf("PaymentTermsDays = %d, want %d", got.PaymentTermsDays, defaultPaymentTermsDays)
	}
	if got.DefaultCurrency != defaultCurrency {
		t.Errorf("DefaultCurrency = %q, want %q", got.DefaultCurrency, defaultCurrency)
	}
	if got.Production {
		t.Error("Production is true with no ICRM_ENV set")
	}
	if !got.SessionKeyGenerated {
		t.Error("SessionKeyGenerated is false, but no key was supplied")
	}
	if len(got.SessionKey) != sessionKeyBytes {
		t.Errorf("generated key is %d bytes, want %d", len(got.SessionKey), sessionKeyBytes)
	}
}

func TestLoadTakesTheEnvironmentOverTheDefaults(t *testing.T) {
	clearEnv(t)
	t.Setenv("ICRM_ADDR", "127.0.0.1:9000")
	t.Setenv("ICRM_DB", "/srv/crm.db")
	t.Setenv("ICRM_COMPANY_NAME", "Northwind")
	t.Setenv("ICRM_CURRENCY", "eur")
	t.Setenv("ICRM_PAYMENT_TERMS_DAYS", "14")

	got, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if got.Addr != "127.0.0.1:9000" {
		t.Errorf("Addr = %q", got.Addr)
	}
	if got.DatabasePath != "/srv/crm.db" {
		t.Errorf("DatabasePath = %q", got.DatabasePath)
	}
	if got.CompanyName != "Northwind" {
		t.Errorf("CompanyName = %q", got.CompanyName)
	}
	// A currency is upper-cased on the way in, so the rest of the application
	// never has to wonder which case it is holding.
	if got.DefaultCurrency != "EUR" {
		t.Errorf("DefaultCurrency = %q, want EUR", got.DefaultCurrency)
	}
	if got.PaymentTermsDays != 14 {
		t.Errorf("PaymentTermsDays = %d, want 14", got.PaymentTermsDays)
	}
}

func TestProductionRefusesToStartWithoutASessionKey(t *testing.T) {
	// Generating a key in production would sign every user out on each deploy,
	// and the resulting report reads as a session bug rather than as a missing
	// setting, so this must be a hard failure.
	clearEnv(t)
	t.Setenv("ICRM_ENV", "production")

	if _, err := Load(); !errors.Is(err, ErrMissingSessionKey) {
		t.Fatalf("Load in production with no key returned %v, want ErrMissingSessionKey", err)
	}

	t.Setenv("ICRM_SESSION_KEY", "a-key-that-is-long-enough-to-pass-32")
	got, err := Load()
	if err != nil {
		t.Fatalf("Load with a key: %v", err)
	}
	if !got.Production {
		t.Error("Production is false with ICRM_ENV=production")
	}
	if got.SessionKeyGenerated {
		t.Error("SessionKeyGenerated is true, but a key was supplied")
	}
}

func TestLoadRejectsSettingsThatWouldFailSilentlyLater(t *testing.T) {
	cases := []struct {
		name     string
		variable string
		value    string
	}{
		{"a session key too short to sign with", "ICRM_SESSION_KEY", "too-short"},
		{"payment terms that are not a number", "ICRM_PAYMENT_TERMS_DAYS", "thirty"},
		{"negative payment terms", "ICRM_PAYMENT_TERMS_DAYS", "-5"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			clearEnv(t)
			t.Setenv(c.variable, c.value)
			if _, err := Load(); err == nil {
				t.Fatalf("Load with %s=%q succeeded, want an error", c.variable, c.value)
			}
		})
	}
}

func TestDueDateForAppliesTheConfiguredTerms(t *testing.T) {
	issued := time.Date(2026, time.October, 21, 9, 0, 0, 0, time.UTC)

	cases := []struct {
		days int
		want string
	}{
		{0, "2026-10-21"},
		{30, "2026-11-20"},
		{14, "2026-11-04"},
		// Across a month boundary of a different length, and a leap year.
		{134, "2027-03-04"},
	}
	for _, c := range cases {
		s := Settings{PaymentTermsDays: c.days}
		if got := s.DueDateFor(issued).Format("2006-01-02"); got != c.want {
			t.Errorf("%d day terms on %s gave %s, want %s",
				c.days, issued.Format("2006-01-02"), got, c.want)
		}
	}
}
