package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
	*kafka.Producer
}

func NewKafkaProducer(host string, port string) *Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%s", host, port)})
	if err != nil {
		panic(err)
	}
	return &Producer{
		p,
	}

}

func (kp *Producer) ProduceMsg(topics []string, msg any) error {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatalln(err)
	}
	for _, topic := range topics {
		err = kp.Producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          data,
		}, nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return err
}
