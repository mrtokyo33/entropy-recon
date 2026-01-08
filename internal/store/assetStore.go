package store

import "entropy-recon/internal/models"

type AssetFilter struct {
	Type  *models.AssetType
	State *models.AssetState
	Value *string
}

type AssetStore interface {
	Save(asset models.Asset) error
	Exists(asset models.Asset) (bool, error)
	List(filter AssetFilter) ([]models.Asset, error)
}
