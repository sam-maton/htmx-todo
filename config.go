package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sam-maton/htmx-todo/internal/database"
)

type serverConfig struct {
	db        *database.Queries
	jwtSecret string
}

func setupServerConfig() serverConfig {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	secret := os.Getenv("SECRET")

	if dbURL == "" {
		log.Fatal("DB_URL environment variable must be set")
	}

	if secret == "" {
		log.Fatal("SECRET environment variable must be set")
	}

	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Printf("There was an error connecting to the db: %s", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	return serverConfig{
		db:        dbQueries,
		jwtSecret: secret,
	}
}
