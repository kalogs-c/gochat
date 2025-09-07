package migrations

import (
	"database/sql"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
	"github.com/pressly/goose/v3/database"
)

//go:embed *.sql
var Embed embed.FS

func MustProvide(db *sql.DB) *goose.Provider {
	provider, err := goose.NewProvider(database.DialectSQLite3, db, Embed)
	if err != nil {
		log.Fatalf("migrations provider failed %v", err)
	}
	return provider
}
