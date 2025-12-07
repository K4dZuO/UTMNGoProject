package kafka

import (
    "log"
    "github.com/IBM/sarama"
)

func NewSyncProducer(brokers []string) (sarama.SyncProducer, error) {
    cfg := sarama.NewConfig()
    cfg.Producer.Return.Successes = true
    cfg.Producer.RequiredAcks = sarama.WaitForAll
    cfg.Producer.Retry.Max = 5

    return sarama.NewSyncProducer(brokers, cfg)
}

func SendMessage(producer sarama.SyncProducer, topic, message string) error {
    msg := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.StringEncoder(message),
    }

    partition, offset, err := producer.SendMessage(msg)
    if err != nil {
        return err
    }

    log.Printf("Message sent => partition %d, offset %d\n", partition, offset)
    return nil
}
