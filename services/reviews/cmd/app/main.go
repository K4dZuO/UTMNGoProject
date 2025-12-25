package main

import (
	"log"
	"fmt"

	"reviews_service/internal/bootstrap"
	"reviews_service/internal/config"
)

func main() {
	cfg, err := config.Load("config.yaml")
	fmt.Println(cfg)
	if err != nil {
		log.Fatal(err)
	} 

	if err := bootstrap.Run(bootstrap.Config{
		HTTPAddr:    cfg.HTTP.Addr,
		PostgresDSN: cfg.Postgres.DSN,
		KafkaBrokers: cfg.Kafka.Brokers,
		KafkaTopic:   cfg.Kafka.Topic,
	}); err != nil {
		log.Fatal(err)
	}
}
