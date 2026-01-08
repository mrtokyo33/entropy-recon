package sql

const (
	insertAssetQuery = `
		INSERT OR IGNORE INTO assets
		(id, type, value, state, source_tool, source_stage, source_metadata, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	existsAssetQuery = `
		SELECT 1 FROM assets
		WHERE type = ? AND value = ?
		LIMIT 1
	`

	listAssetsBaseQuery = `
		SELECT
			id,
			type,
			value,
			state,
			source_tool,
			source_stage,
			source_metadata,
			created_at
		FROM assets
		WHERE 1 = 1
	`
)
