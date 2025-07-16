package api

import (
	"fmt"
	"log"
	"net/http"
)

func HandleBasePath(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello World!")
	if err != nil {
		log.Fatalf("Could not write response: %s", err)
	}
}
