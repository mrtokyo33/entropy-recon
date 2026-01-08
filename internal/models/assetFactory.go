package models

import (
	"time"

	"github.com/google/uuid"
)

func NewAsset(
	assetType AssetType,
	value string,
	source AssetSource,
) Asset {
	return Asset{
		ID:        uuid.NewString(),
		Type:      assetType,
		Value:     value,
		State:     StateDiscovered,
		Source:    source,
		Metadata:  make(map[string]string),
		CreatedAt: time.Now(),
	}
}
