package domain

import (
	"errors"
	"fmt"
	"strings"
)

// CustomerStatus carries the lead-to-customer pipeline. A lead and a customer
// are the same record at different stages rather than two tables, because the
// alternative is copying a record on conversion and then owning two versions of
// one company's details.
type CustomerStatus string

const (
	CustomerLead     CustomerStatus = "lead"
	CustomerActive   CustomerStatus = "active"
	CustomerArchived CustomerStatus = "archived"
)

var (
	ErrUnknownCustomerStatus = errors.New("domain: unknown customer status")
	ErrMissingName           = errors.New("domain: name is required")
	ErrMalformedCurrency     = errors.New("domain: currency must be a three-letter ISO 4217 code")
)

// ParseCustomerStatus validates an untrusted string against the closed set.
func ParseCustomerStatus(s string) (CustomerStatus, error) {
	status := CustomerStatus(s)
	switch status {
	case CustomerLead, CustomerActive, CustomerArchived:
		return status, nil
	}
	return "", fmt.Errorf("%w: %q", ErrUnknownCustomerStatus, s)
}

// Label is the human name for the status.
func (s CustomerStatus) Label() string {
	switch s {
	case CustomerLead:
		return "Lead"
	case CustomerActive:
		return "Customer"
	case CustomerArchived:
		return "Archived"
	}
	return string(s)
}

// CustomerStatuses lists the pipeline in order, for filters and dropdowns.
func CustomerStatuses() []CustomerStatus {
	return []CustomerStatus{CustomerLead, CustomerActive, CustomerArchived}
}

// ValidateName rejects a blank or whitespace-only name and returns the trimmed
// value, so the store never holds a record that cannot be identified in a list.
func ValidateName(s string) (string, error) {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return "", ErrMissingName
	}
	return trimmed, nil
}

// ValidateCurrency accepts a three-letter ISO 4217 code and upper-cases it.
// The code is not checked against the published list: a self-hosted install may
// legitimately bill in something the list does not carry yet, and rejecting it
// would be a worse failure than accepting a typo the user can see on screen.
func ValidateCurrency(s string) (string, error) {
	code := strings.ToUpper(strings.TrimSpace(s))
	if len(code) != 3 {
		return "", fmt.Errorf("%w: %q", ErrMalformedCurrency, s)
	}
	for i := 0; i < len(code); i++ {
		if code[i] < 'A' || code[i] > 'Z' {
			return "", fmt.Errorf("%w: %q", ErrMalformedCurrency, s)
		}
	}
	return code, nil
}
