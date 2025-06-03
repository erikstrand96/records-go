package api

import "net/http"

func CreateServer() *http.Server {

	s := http.Server{Addr: "localhost:8181"}
	return &s
}
