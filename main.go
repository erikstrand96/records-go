package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"records-go/api"
	"records-go/database"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {

	var db *sql.DB

	_, err := os.Stat(".env")
	if err == nil {

		err = godotenv.Load()
		if err != nil {
			log.Printf("Error loading .env file. %s", err)
		}
	}

	connStr := os.Getenv("RECORDS_DB_CONNECTION")
	dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	defer dbPool.Close()

	db, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	err = database.ApplyMigrations(db)
	if err != nil {
		log.Fatalf("Could not apply migrations: %s", err)
	}

	srv := api.CreateServer()
	fmt.Printf("Server address: %s\n", srv.Addr)

	http.HandleFunc("/", api.HandleBasePath)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Could not start http server: ", err)
	}

}
