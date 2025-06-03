package api

import (
	"fmt"
	"log"
	"net/http"
)

func HandleBasePath(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World!")
	if err != nil {
		log.Fatal("Could not write response!")
	}
}
