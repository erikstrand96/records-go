package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func ManageDatabase() *pgxpool.Pool {

	connStr := os.Getenv("RECORDS_DB_CONNECTION")

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}

	return pool

}
