package primitives

import (
	"errors"
	"net"
	"strings"
)

var ErrInvalidIP = errors.New("invalid ip address")

type IPAddress struct {
	value string
}

func NewIPAddress(raw string) (IPAddress, error) {
	ip := strings.TrimSpace(raw)

	parsed := net.ParseIP(ip)
	if parsed == nil {
		return IPAddress{}, ErrInvalidIP
	}

	return IPAddress{value: parsed.String()}, nil
}

func (ip IPAddress) String() string {
	return ip.value
}

func (ip IPAddress) IsIPv6() bool {
	return strings.Contains(ip.value, ":")
}
