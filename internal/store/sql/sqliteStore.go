package sql

import (
	"database/sql"
	_ "embed"
	"encoding/json"
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"entropy-recon/internal/models"
	"entropy-recon/internal/store"
)

//go:embed schema.sql
var schema string

type SQLiteAssetStore struct {
	db *sql.DB
}

func NewSQLiteAssetStore(dbPath string) (*SQLiteAssetStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := applySchema(db); err != nil {
		return nil, err
	}

	return &SQLiteAssetStore{db: db}, nil
}

func applySchema(db *sql.DB) error {
	if schema == "" {
		return errors.New("empty database schema")
	}

	_, err := db.Exec(schema)
	return err
}

func (s *SQLiteAssetStore) Save(asset models.Asset) error {
	var metadataJSON []byte
	if asset.Source.Metadata != nil {
		metadataJSON, _ = json.Marshal(asset.Source.Metadata)
	}

	_, err := s.db.Exec(
		insertAssetQuery,
		asset.ID,
		asset.Type,
		asset.Value,
		asset.State,
		asset.Source.Tool,
		asset.Source.Stage,
		string(metadataJSON),
		asset.CreatedAt,
	)

	return err
}

func (s *SQLiteAssetStore) Exists(asset models.Asset) (bool, error) {
	row := s.db.QueryRow(
		existsAssetQuery,
		asset.Type,
		asset.Value,
	)

	var tmp int
	err := row.Scan(&tmp)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *SQLiteAssetStore) List(filter store.AssetFilter) ([]models.Asset, error) {
	query := listAssetsBaseQuery
	args := []any{}

	if filter.Type != nil {
		query += " AND type = ?"
		args = append(args, *filter.Type)
	}

	if filter.State != nil {
		query += " AND state = ?"
		args = append(args, *filter.State)
	}

	if filter.Value != nil {
		query += " AND value = ?"
		args = append(args, *filter.Value)
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []models.Asset

	for rows.Next() {
		var (
			a           models.Asset
			metadataRaw sql.NullString
		)

		err := rows.Scan(
			&a.ID,
			&a.Type,
			&a.Value,
			&a.State,
			&a.Source.Tool,
			&a.Source.Stage,
			&metadataRaw,
			&a.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if metadataRaw.Valid {
			metadata := map[string]string{}
			_ = json.Unmarshal([]byte(metadataRaw.String), &metadata)
			a.Source.Metadata = metadata
		}

		assets = append(assets, a)
	}

	return assets, nil
}
