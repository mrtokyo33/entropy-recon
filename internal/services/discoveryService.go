package services

import (
	"context"
	"log"
	"strings"

	"entropy-recon/internal/models"
	"entropy-recon/internal/models/primitives"
	"entropy-recon/internal/store"
	"entropy-recon/internal/tools"
)

type DiscoveryService struct {
	store store.AssetStore
	tool  tools.SubfinderTool
}

func NewDiscoveryService(
	store store.AssetStore,
	tool tools.SubfinderTool,
) *DiscoveryService {
	return &DiscoveryService{
		store: store,
		tool:  tool,
	}
}

func splitSubdomain(raw string, domain primitives.Domain) (string, primitives.Domain, bool) {
	if raw == domain.String() {
		return "", domain, true
	}

	if strings.HasSuffix(raw, "."+domain.String()) {
		sub := strings.TrimSuffix(raw, "."+domain.String())
		return sub, domain, true
	}

	return "", primitives.Domain{}, false
}

func (s *DiscoveryService) Run(ctx context.Context, domain primitives.Domain) error {
	results, err := s.tool.Run(ctx, domain.String())
	if err != nil {
		return err
	}

	for _, result := range results {
		rawHost := strings.TrimSpace(result.Host)
		if rawHost == "" {
			continue
		}

		log.Printf(
			"[discovery] raw=%s source=%s",
			rawHost,
			result.Source,
		)

		if ip, err := primitives.NewIPAddress(rawHost); err == nil {
			log.Printf("[ip] detected ip=%s", ip.String())

			host := primitives.NewHostFromIP(ip)

			asset := models.NewAsset(
				models.AssetHost,
				host.String(),
				models.AssetSource{
					Tool:  "subfinder",
					Stage: "passive",
					Metadata: map[string]string{
						"source": result.Source,
					},
				},
			)

			if exists, _ := s.store.Exists(asset); !exists {
				_ = s.store.Save(asset)
			}

			continue
		}

		sub, dom, ok := splitSubdomain(rawHost, domain)
		if !ok {
			log.Printf("[discard] out-of-scope fqdn=%s", rawHost)
			continue
		}

		subPart, err := primitives.NewSubdomainPart(sub)
		if err != nil {
			log.Printf("[discard] invalid subdomain part=%s", sub)
			continue
		}

		log.Printf("[subdomain] part=%s", subPart.String())
		log.Printf("[domain] base=%s", dom.String())

		fqdn, err := primitives.NewFQDN(subPart, dom)
		if err != nil {
			log.Printf("[discard] invalid fqdn=%s", rawHost)
			continue
		}

		log.Printf(
			"[fqdn] full=%s sub=%s domain=%s source=%s",
			fqdn.String(),
			subPart.String(),
			dom.String(),
			result.Source,
		)

		host := primitives.NewHostFromFQDN(fqdn)

		asset := models.NewAsset(
			models.AssetHost,
			host.String(),
			models.AssetSource{
				Tool:  "subfinder",
				Stage: "passive",
				Metadata: map[string]string{
					"source": result.Source,
				},
			},
		)

		if exists, _ := s.store.Exists(asset); !exists {
			_ = s.store.Save(asset)
		}
	}

	return nil
}
