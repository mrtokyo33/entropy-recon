package store

import (
	"sync"

	"entropy-recon/internal/models"
)

type MemoryAssetStore struct {
	mu     sync.RWMutex
	assets map[string]models.Asset
}

func NewMemoryAssetStore() *MemoryAssetStore {
	return &MemoryAssetStore{
		assets: make(map[string]models.Asset),
	}
}

func (s *MemoryAssetStore) Save(asset models.Asset) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	key := string(asset.Type) + ":" + asset.Value
	s.assets[key] = asset
	return nil
}

func (s *MemoryAssetStore) Exists(asset models.Asset) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	key := string(asset.Type) + ":" + asset.Value
	_, ok := s.assets[key]
	return ok, nil
}

func (s *MemoryAssetStore) List(filter AssetFilter) ([]models.Asset, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []models.Asset

	for _, asset := range s.assets {
		if filter.Type != nil && asset.Type != *filter.Type {
			continue
		}
		if filter.State != nil && asset.State != *filter.State {
			continue
		}
		if filter.Value != nil && asset.Value != *filter.Value {
			continue
		}
		result = append(result, asset)
	}

	return result, nil
}
