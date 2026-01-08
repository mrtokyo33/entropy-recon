package primitives

import "errors"

var ErrInvalidFQDN = errors.New("invalid fqdn")

type FQDN struct {
	subdomain SubdomainPart
	domain    Domain
}

func NewFQDN(sub SubdomainPart, domain Domain) (FQDN, error) {
	if domain.String() == "" {
		return FQDN{}, ErrInvalidFQDN
	}

	return FQDN{
		subdomain: sub,
		domain:    domain,
	}, nil
}

func (f FQDN) String() string {
	if f.subdomain.String() == "" {
		return f.domain.String()
	}

	return f.subdomain.String() + "." + f.domain.String()
}

func (f FQDN) Domain() Domain {
	return f.domain
}

func (f FQDN) SubdomainPart() SubdomainPart {
	return f.subdomain
}
