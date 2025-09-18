package main

import (
	"fmt"
	"log"
	"stocky/internal/api"
	"stocky/internal/config"
	"stocky/internal/database"
)

func main() {
	cfg := config.LoadConfig()
	database.Connect(cfg)
	
	r := api.SetupRouter()
	
	addr := fmt.Sprintf(":%s", cfg.SERVER_PORT)
	log.Printf("Starting server on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}