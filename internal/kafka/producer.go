package kafka

import (
	"encoding/json"

	"github.com/IBM/sarama"
)

func NewSyncProducer(brokers []string) (Producer, error) {
    cfg := sarama.NewConfig()
    cfg.Producer.Return.Successes = true

    p, err := sarama.NewSyncProducer(brokers, cfg)
    if err != nil {
        return nil, err
    }

    return &SyncProducer{p: p}, nil
}

type Producer interface {
    Send(topic string, event interface{}) error
    Close() error
}

func (s *SyncProducer) Close() error {
    return s.p.Close()
}


type SyncProducer struct {
    p sarama.SyncProducer
}

func (s *SyncProducer) Send(topic string, event interface{}) error {
    data, _ := json.Marshal(event)
    msg := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.StringEncoder(data),
    }
    _, _, err := s.p.SendMessage(msg)
    return err
}

type RecalculationEvent struct {
    ProductID int `json:"product_id"`
    ReviewID  string `json:"review_id"`
}
