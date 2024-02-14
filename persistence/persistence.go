package persistence

import (
	"context"
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gofiber/fiber/v2/log"
)

//go:embed schema.sql
var ddl string

//go:embed mockData.sql
var mockData string

func createNewDB(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Error("Could not open db connection", err)
		return nil, err
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Error("Could not execute DDL", err)
	}
	return db, nil
}

func fillWithMockData(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, mockData); err != nil {
		log.Error("Could not insert mock data")
	}
	return nil
}

func SetupStore() (*Queries, error) {
	ctx := context.Background()

	db, err := createNewDB(ctx)
	if err != nil {
		return nil, err
	}

	fillWithMockData(ctx, db)

	store := New(db)

	return store, nil
}
