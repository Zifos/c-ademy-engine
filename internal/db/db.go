package db

import (
	"c-ademy/internal/db/sqlc"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Queries struct {
	*sqlc.Queries
	db *sql.DB
}

func New(dataSourceName string) (*Queries, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	q := sqlc.New(db)
	return &Queries{
		Queries: q,
		db:      db,
	}, nil
}

func (q *Queries) Close() error {
	return q.db.Close()
}
