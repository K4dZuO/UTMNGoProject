package main

import (
	"log"

	"rating_service/internal/bootstrap"
	"rating_service/internal/config"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	if err := bootstrap.Run(*cfg); err != nil {
		log.Fatalf("application error: %v", err)
	}
}
