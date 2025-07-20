package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"records-go/api"
	"records-go/database"
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

	dbPool, err := database.ManageDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	defer dbPool.Close()

	connStr := dbPool.Config().ConnString()
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	defer db.Close()

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
