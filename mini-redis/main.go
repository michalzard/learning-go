package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/michalzard/learning-go/mini-redis/internal/database"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	// Definetly need to use the controller/router approach next project
	// What i did here is plain awful

	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	queries := database.New(conn)
	server := newAPIServer("127.0.0.1:" + port)
	server.Run(queries)

}
