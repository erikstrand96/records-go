package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func CreateServer() *http.Server {

	port, err := strconv.Atoi(os.Getenv("RECORDS_API_PORT"))
	if err != nil {
		log.Fatal("Could not find port number!\n")
	}

	s := http.Server{Addr: fmt.Sprintf(":%d", port)}
	return &s
}
