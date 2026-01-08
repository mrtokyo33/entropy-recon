package models

import "time"

type Asset struct {
	ID    string
	Type  AssetType
	Value string

	State  AssetState
	Source AssetSource

	Metadata  map[string]string
	CreatedAt time.Time
}
