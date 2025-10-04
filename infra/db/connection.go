package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewConnection() (*sqlx.DB, error) {
	dsn := ""
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	return db, nil
}
