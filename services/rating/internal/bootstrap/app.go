package bootstrap

import (
	"context"
	"log"

	"rating_service/internal/config"
	"rating_service/internal/services/ratingService"
)

func Run(cfg config.Config) error {
	ctx := context.Background()

	// === Postgres ===
	storage, err := initStorage(ctx, cfg.Postgres.DSN)
	if err != nil {
		return err
	}

	// === Redis ===
	cache, err := initRedis(cfg.Redis.Addr)
	if err != nil {
		return err
	}

	// === Service ===
	service := ratingService.New(storage, cache)

	// === Kafka consumer ===
	go initKafkaConsumer(cfg.Kafka, service)

	// === HTTP ===
	router := initHTTP(service)

	log.Printf("rating service started on %s", cfg.HTTP.Addr)
	return router.Run(cfg.HTTP.Addr)
}
