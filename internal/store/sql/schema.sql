CREATE TABLE IF NOT EXISTS assets (
    id TEXT PRIMARY KEY,
    type TEXT NOT NULL,
    value TEXT NOT NULL,
    state TEXT NOT NULL,
    source_tool TEXT NOT NULL,
    source_stage TEXT NOT NULL,
    created_at DATETIME NOT NULL,

    UNIQUE(type, value)
);

CREATE INDEX IF NOT EXISTS idx_assets_type ON assets(type);
CREATE INDEX IF NOT EXISTS idx_assets_state ON assets(state);
