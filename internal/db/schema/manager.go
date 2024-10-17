package schema

import (
	"database/sql"
	"embed"
)

//go:embed migrations/*.sql
var schema embed.FS

//go:embed seed.sql
var seed string

type SchemaManager struct {
	db *sql.DB
}

func NewSchemaManager(dbPath string) (*SchemaManager, error) {
	// Create the connection to the database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return &SchemaManager{
		db: db,
	}, nil
}

func (sm *SchemaManager) Seed() error {
	_, err := sm.db.Exec(seed)
	return err
}
