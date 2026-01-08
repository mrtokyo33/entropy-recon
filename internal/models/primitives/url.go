package primitives

import (
	"errors"
	"net/url"
	"strings"
)

var ErrInvalidURL = errors.New("invalid url")

type URL struct {
	scheme string
	host   Host
	raw    string
}

func NewURL(raw string) (URL, error) {
	u := strings.TrimSpace(raw)

	parsed, err := url.Parse(u)
	if err != nil {
		return URL{}, ErrInvalidURL
	}

	if parsed.Scheme == "" || parsed.Host == "" {
		return URL{}, ErrInvalidURL
	}

	scheme := strings.ToLower(parsed.Scheme)

	host, err := NewHost(parsed.Host)
	if err != nil {
		return URL{}, ErrInvalidURL
	}

	normalized := scheme + "://" + host.String()

	return URL{
		scheme: scheme,
		host:   host,
		raw:    normalized,
	}, nil
}

func (u URL) String() string {
	return u.raw
}

func (u URL) Host() Host {
	return u.host
}

func (u URL) Scheme() string {
	return u.scheme
}
