package main

import (
	"log"

	"recommend_service/internal/bootstrap"
	"recommend_service/internal/config"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	app, err := bootstrap.Build(cfg)
	if err != nil {
		log.Fatalf("bootstrap error: %v", err)
	}

	if err := app.HTTP.Run(); err != nil {
		log.Fatalf("http server error: %v", err)
	}
}
