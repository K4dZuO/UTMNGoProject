package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

type RatingService interface {
	HandleReviewCreated(ctx context.Context, productID int) error
}

type Consumer struct {
	service RatingService
}

func NewConsumer(service RatingService) *Consumer {
	return &Consumer{service: service}
}

func (c *Consumer) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (c *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (c *Consumer) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {

	for msg := range claim.Messages() {

		var event ReviewCreatedEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("[Kafka] invalid message: %v", err)
			session.MarkMessage(msg, "")
			continue
		}

		if err := c.service.HandleReviewCreated(
			context.Background(),
			event.ProductID,
		); err != nil {
			log.Printf(
				"[Kafka] rating recalculation failed (product=%d): %v",
				event.ProductID,
				err,
			)
		}

		session.MarkMessage(msg, "")
	}

	return nil
}
