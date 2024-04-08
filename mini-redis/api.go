package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/michalzard/learning-go/mini-redis/internal/database"
)

type APIServer struct {
	address string
}

func newAPIServer(address string) *APIServer {
	return &APIServer{
		address: address,
	}
}

func (api *APIServer) Run(queries *database.Queries) {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to mini redis")
	})
	apiConfig := APIConfig{DB: queries}

	router.HandleFunc("POST /key", func(w http.ResponseWriter, r *http.Request) {
		apiConfig.CreateKV(w, r)
	})

	fmt.Printf("API listening to %v\n", api.address)
	log.Fatal(http.ListenAndServe(api.address, router))
}

type APIConfig struct {
	DB *database.Queries
}

// this is basically controller
func (cfg *APIConfig) CreateKV(w http.ResponseWriter, r *http.Request) {

	type Params struct {
		K string `json:"key"`
		V string `json:"value"`
	}
	decoder := json.NewDecoder(r.Body)

	parameters := Params{}

	err := decoder.Decode(&parameters)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	kv, err := cfg.DB.CreateKeyValue(r.Context(), database.CreateKeyValueParams{
		K:         parameters.K,
		V:         parameters.V,
		CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		fmt.Printf("%v", err)
		http.Error(w, "Error saving key value", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(kv)

	if err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
