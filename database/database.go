package database

import (
	"context"
	"database/sql"
	"embed"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

func ManageDatabase() (*pgxpool.Pool, error) {

	connStr := os.Getenv("RECORDS_DB_CONNECTION")

	pool, err := pgxpool.New(context.Background(), connStr)

	return pool, err

}

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
