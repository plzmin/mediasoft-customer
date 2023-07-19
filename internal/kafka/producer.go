package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"mediasoft-customer/internal/model"
)

type Producer struct {
	producer sarama.SyncProducer
}

func New(broker []string) (*Producer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Retry.Max = 5
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(broker, cfg)
	if err != nil {
		return nil, err
	}
	return &Producer{producer: producer}, err
}

func (p *Producer) SendMessage(topic string, message model.Order) error {
	orderJson, err := json.Marshal(message)
	if err != nil {
		return err
	}
	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(orderJson),
	}
	_, _, err = p.producer.SendMessage(&msg)
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) Close() error {
	return p.producer.Close()
}
