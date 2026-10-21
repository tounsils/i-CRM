package domain

import (
	"errors"
	"fmt"
)

// MessageDirection distinguishes what was said to a customer from what they
// said back, and from a staff-only note. Threading all three into one timeline
// is the point of the messaging feature: the history of a deal is the history
// of its correspondence.
type MessageDirection string

const (
	MessageOutbound MessageDirection = "outbound"
	MessageInbound  MessageDirection = "inbound"
	MessageNote     MessageDirection = "note"
)

var ErrUnknownDirection = errors.New("domain: unknown message direction")

// ParseMessageDirection validates an untrusted string against the closed set.
func ParseMessageDirection(s string) (MessageDirection, error) {
	d := MessageDirection(s)
	switch d {
	case MessageOutbound, MessageInbound, MessageNote:
		return d, nil
	}
	return "", fmt.Errorf("%w: %q", ErrUnknownDirection, s)
}

// Label is the human name for the direction.
func (d MessageDirection) Label() string {
	switch d {
	case MessageOutbound:
		return "Sent to customer"
	case MessageInbound:
		return "From customer"
	case MessageNote:
		return "Internal note"
	}
	return string(d)
}

// MessageDirections lists the set, for dropdowns.
func MessageDirections() []MessageDirection {
	return []MessageDirection{MessageOutbound, MessageInbound, MessageNote}
}
