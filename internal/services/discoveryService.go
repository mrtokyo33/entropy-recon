package services

import (
	"context"

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

func (s *DiscoveryService) Run(ctx context.Context, domain primitives.Domain) error {
	results, err := s.tool.Run(ctx, domain.String())
	if err != nil {
		return err
	}

	for _, raw := range results {
		host, err := primitives.NewHost(raw)
		if err != nil {
			continue
		}

		asset := models.NewAsset(
			models.AssetHost,
			host.String(),
			models.AssetSource{
				Tool:  "subfinder",
				Stage: "discovery",
			},
		)

		exists, _ := s.store.Exists(asset)
		if exists {
			continue
		}

		_ = s.store.Save(asset)
	}

	return nil
}
