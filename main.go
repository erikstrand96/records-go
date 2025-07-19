package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"records-go/api"

	"github.com/joho/godotenv"
)

func main() {

	_, err := os.Stat(".env")
	if err == nil {

		err = godotenv.Load()
		if err != nil {
			log.Printf("Error loading .env file. %s", err)
		}
	}

	srv := api.CreateServer()
	fmt.Printf("Server address: %s\n", srv.Addr)

	http.HandleFunc("/", api.HandleBasePath)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Could not start http server: ", err)
	}

}
