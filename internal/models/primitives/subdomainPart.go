package primitives

import (
	"errors"
	"strings"
)

var ErrInvalidSubdomainPart = errors.New("invalid subdomain part")

type SubdomainPart struct {
	value string
}

func NewSubdomainPart(raw string) (SubdomainPart, error) {
	s := strings.ToLower(strings.TrimSpace(raw))

	if s == "" {
		return SubdomainPart{}, ErrInvalidSubdomainPart
	}

	if strings.ContainsAny(s, " */@") {
		return SubdomainPart{}, ErrInvalidSubdomainPart
	}

	if strings.Contains(s, "..") {
		return SubdomainPart{}, ErrInvalidSubdomainPart
	}

	return SubdomainPart{value: s}, nil
}

func (s SubdomainPart) String() string {
	return s.value
}

func (s SubdomainPart) IsRoot() bool {
	return s.value == ""
}
