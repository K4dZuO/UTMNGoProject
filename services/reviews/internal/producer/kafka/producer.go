package kafka

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
)

type Producer struct {
    producer sarama.SyncProducer
    topic string
}


func New(brokers []string, topic string) (*Producer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	p, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return nil, err
	}

	return &Producer{
		producer: p,
		topic:    topic,
	}, nil
}

func (p *Producer) Close() error {
	return p.producer.Close()
}

func (p *Producer) SendReviewCreated(ctx context.Context, product_id int) error {
    event := ReviewCreatedEvent{ProductID: product_id}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg)
	return err
}
