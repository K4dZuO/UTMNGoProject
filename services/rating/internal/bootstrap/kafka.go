package bootstrap

import (
	"log"

	"rating_service/internal/config"
	"rating_service/internal/kafka"
	"rating_service/internal/services/ratingService"
)

func initKafkaConsumer(cfg config.KafkaConfig, service *ratingService.Service) {
	if err := kafka.RunConsumer(cfg.Brokers, cfg.GroupID, cfg.Topic, service); err != nil {
		log.Fatalf("kafka consumer error: %v", err)
	}
}
