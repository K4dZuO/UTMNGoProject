package kafka

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

func RunConsumer(
	brokers []string,
	groupID string,
	topic string,
	service RatingService,
) error {

	cfg := sarama.NewConfig()
	cfg.Consumer.Return.Errors = true

	group, err := sarama.NewConsumerGroup(brokers, groupID, cfg)
	if err != nil {
		return err
	}

	consumer := NewConsumer(service)

	go func() {
		for {
			if err := group.Consume(
				context.Background(),
				[]string{topic},
				consumer,
			); err != nil {
				log.Printf("[Kafka] consume error: %v", err)
			}
		}
	}()

	log.Printf(
		"[Kafka] consumer started (group=%s, topic=%s)",
		groupID,
		topic,
	)

	return nil
}
