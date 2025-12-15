package bootstrap

import (
	"reviews_service/internal/producer/kafka"
)


func initKafka(brokers []string, topic string) (*kafka.Producer, error) {
	return kafka.New(brokers, topic)
}
