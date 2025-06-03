package main

import (
	"fmt"
	"log"
	"net/http"
	"records-go/api"
)

func main() {

	srv := api.CreateServer()
	fmt.Printf("Server address: %s\n", srv.Addr)

	http.HandleFunc("/", api.HandleBasePath)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Could not start http server: ", err)
	}

}
