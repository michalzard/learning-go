package main

import (
	"fmt"
	"log"
	"net/http"
)

type APIServer struct {
	address string
}

func newAPIServer(address string) *APIServer {
	return &APIServer{
		address: address,
	}
}

func (api *APIServer) Run() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to mini redis")
	})

	// router.HandleFunc("GET /{key}", JSONMiddleware())
	/*

		GET key
		PUT key
		DELETE key
	*/

	fmt.Printf("API listetning to %v§\n", api.address)
	log.Fatal(http.ListenAndServe(api.address, router))
}
