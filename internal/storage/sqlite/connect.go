package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func MustConnect(ctx context.Context, dsn string) *sql.DB {
	db, err := Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("db connection failed %v", err)
	}
	return db
}

func Connect(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db %s failed: %w", dsn, err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping db error: %w", err)
	}

	return db, nil
}
