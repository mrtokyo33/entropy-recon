package primitives

import (
	"errors"
	"strings"
)

var ErrInvalidHost = errors.New("invalid host")

type Host struct {
	value string
}

func NewHost(raw string) (Host, error) {
	h := strings.ToLower(strings.TrimSpace(raw))

	h = strings.TrimSuffix(h, ".")

	if h == "" || strings.Contains(h, " ") {
		return Host{}, ErrInvalidHost
	}

	if !strings.Contains(h, ".") {
		return Host{}, ErrInvalidHost
	}

	return Host{value: h}, nil
}

func (h Host) String() string {
	return h.value
}

func (h Host) Equals(other Host) bool {
	return h.value == other.value
}
