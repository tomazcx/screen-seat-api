package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/tomazcx/screen-seat-api/config"
	"github.com/tomazcx/screen-seat-api/internal/infra/database"
)

func main(){

	conf, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading the configuration file: %v", err)
	}

	db, err := database.ConnectToDB(conf)
	if err != nil { 
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()
	uuid.EnableRandPool()
}
