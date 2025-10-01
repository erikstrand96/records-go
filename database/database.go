package database

import (
	"database/sql"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func ApplyMigrations(db *sql.DB) error {

	err := goose.SetDialect("postgres")
	if err != nil {
		log.Fatalf("Could not set dialect: %s", err)
	}

	goose.SetBaseFS(embedMigrations)

	err = goose.Up(db, "migrations")

	return err

}
