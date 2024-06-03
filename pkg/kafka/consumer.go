package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer struct {
	Consumer *kafka.Consumer
}

func NewKafkaConsumer(host string, port string) *Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", host, port),
		"group.id":          "ecommerce",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}
	return &Consumer{Consumer: c}
}

func (kc *Consumer) Subscribe(topics []string) error {
	return kc.Consumer.SubscribeTopics(topics, nil)
}
