package rating

import (
    "context"
    "encoding/json"
    "log"

    "github.com/IBM/sarama"
)

type RecalculationEvent struct {
    ProductID int64  `json:"product_id"`
    ReviewID  string `json:"review_id"`
}

type Consumer struct {
    service *Service
    topic   string
}

func NewConsumer(service *Service, topic string) *Consumer {
    return &Consumer{
        service: service,
        topic:   topic,
    }
}


func (c *Consumer) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (c *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (c *Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

    for msg := range claim.Messages() {

        log.Printf("[Kafka] Message received: %s", string(msg.Value))

        var event RecalculationEvent
        if err := json.Unmarshal(msg.Value, &event); err != nil {
            log.Printf("[Kafka] JSON decode error: %v", err)
            continue
        }

        // Обработка события
        err := c.service.RecalculateAndUpdate(context.Background(), int(event.ProductID))
        if err != nil {
            log.Printf("[Kafka] Failed to recalc rating for product %d: %v", event.ProductID, err)
        }

        sess.MarkMessage(msg, "")
    }
    return nil
}


func RunKafkaConsumer(brokers []string, groupID string, topic string, service *Service) {

    config := sarama.NewConfig()
    config.Consumer.Return.Errors = true

    consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
    if err != nil {
        log.Fatalf("Failed to create consumer group: %v", err)
    }

    consumer := NewConsumer(service, topic)

    go func() {
        for {
            if err := consumerGroup.Consume(context.Background(), []string{topic}, consumer); err != nil {
                log.Printf("Kafka consume error: %v", err)
            }
        }
    }()

    log.Printf("Kafka consumer started: group=%s topic=%s", groupID, topic)
}
