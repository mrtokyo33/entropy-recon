package primitives

import "errors"

var ErrInvalidHost = errors.New("invalid host")

type HostType string

const (
	HostFQDN HostType = "fqdn"
	HostIP   HostType = "ip"
)

type Host struct {
	kind HostType
	fqdn *FQDN
	ip   *IPAddress
}

func NewHostFromFQDN(f FQDN) Host {
	return Host{
		kind: HostFQDN,
		fqdn: &f,
	}
}

func NewHostFromIP(ip IPAddress) Host {
	return Host{
		kind: HostIP,
		ip:   &ip,
	}
}

func (h Host) Type() HostType {
	return h.kind
}

func (h Host) String() string {
	if h.kind == HostIP {
		return h.ip.String()
	}
	return h.fqdn.String()
}

func (h Host) IsInScope(domain Domain) bool {
	if h.kind != HostFQDN {
		return false
	}
	return h.fqdn.Domain().Equals(domain)
}
