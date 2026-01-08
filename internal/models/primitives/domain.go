package primitives

import (
	"errors"
	"strings"
)

var ErrInvalidDomain = errors.New("invalid domain")

type Domain struct {
	value string
}

func NewDomain(raw string) (Domain, error) {
	d := strings.ToLower(strings.TrimSpace(raw))

	if d == "" || strings.Contains(d, " ") {
		return Domain{}, ErrInvalidDomain
	}

	if !strings.Contains(d, ".") {
		return Domain{}, ErrInvalidDomain
	}

	return Domain{value: d}, nil
}

func (d Domain) String() string {
	return d.value
}

func (d Domain) Equals(other Domain) bool {
	return d.value == other.value
}
