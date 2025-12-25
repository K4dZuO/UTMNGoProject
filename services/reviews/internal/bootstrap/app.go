package bootstrap

import (
	"context"

	httpapi "reviews_service/internal/api/http"
	"reviews_service/internal/services/reviewService"
)

type Config struct {
	HTTPAddr string
	PostgresDSN string
	KafkaBrokers []string
	KafkaTopic string
}

func Run(cfg Config) error {
	ctx := context.Background()

	storage, err := initStorage(ctx, cfg.PostgresDSN)
	if err != nil {
		return err
	}

	producer, err := initKafka(cfg.KafkaBrokers, cfg.KafkaTopic)
	if err != nil {
		return err
	}
	defer producer.Close()

	service := reviewService.New(storage, producer)

	handler := httpapi.New(service)
	router := initHTTP(handler)

	return router.Run(cfg.HTTPAddr)
}
