package models

type AssetType string

const (
	AssetDomain AssetType = "domain"
	AssetHost   AssetType = "host"
	AssetURL    AssetType = "url"
)
