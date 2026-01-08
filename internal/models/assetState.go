package models

type AssetState string

const (
	StateDiscovered AssetState = "discovered"
	StateProcessed  AssetState = "processed"
	StateError      AssetState = "error"
)
