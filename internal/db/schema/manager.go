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

func (sm *SchemaManager) CreateSchema() error {
	// Iterate over the schema files and execute
	// each one of them
	files, err := schema.ReadDir("migrations")
	if err != nil {
		return err
	}

	for _, file := range files {
		content, err := schema.ReadFile("migrations/" + file.Name())
		if err != nil {
			return err
		}

		_, err = sm.db.Exec(string(content))
		if err != nil {
			return err
		}
	}

	return nil
}

func (sm *SchemaManager) Seed() error {
	_, err := sm.db.Exec(seed)
	return err
}
