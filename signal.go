package commandline

import (
	"bytes"
	"errors"
	"fmt"
)

// SignalType is a signal
type SignalType int8

const (
	// StopSignal means stop the progress
	StopSignal SignalType = iota
	// KillSignal means kill the progress
	KillSignal
)

var errUnmarshalNilSignal = errors.New("Can't unmarshal a nil *SignalType")

func (s SignalType) String() string {
	switch s {
	case StopSignal:
		return "stop"
	case KillSignal:
		return "kill"
	default:
		return fmt.Sprintf("Signal(%d)", s)
	}
}

// UnmarshalText unmarshals text to a signal.
func (s *SignalType) UnmarshalText(text []byte) error {
	if s == nil {
		return errUnmarshalNilSignal
	}
	if !s.unmarshalText(text) && !s.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("Unrecognized signal: %q", text)
	}
	return nil
}

func (s *SignalType) unmarshalText(text []byte) bool {
	switch string(text) {
	case "stop", "STOP":
		*s = StopSignal
	case "kill", "KILL":
		*s = KillSignal
	default:
		return false
	}
	return true
}
