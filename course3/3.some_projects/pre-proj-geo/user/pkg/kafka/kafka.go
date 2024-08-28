package kafka

import (
	"fmt"

	k "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	Consumer   *k.Consumer
	NotifyChan chan error
}

func New(kafkaServiceAddress, groupID string) (*KafkaConsumer, error) {
	// Create a new Kafka consumer
	c, err := k.NewConsumer(&k.ConfigMap{
		"bootstrap.servers": kafkaServiceAddress,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		return nil, fmt.Errorf("kafka - failed to create consumer: %s", err)
	}

	return &KafkaConsumer{
		Consumer:   c,
		NotifyChan: make(chan error),
	}, nil
}

func (kc *KafkaConsumer) Notify() <-chan error {
	return kc.NotifyChan
}

func (kc *KafkaConsumer) Close() error {
	err := kc.Consumer.Close()
	if err != nil {
		return fmt.Errorf("kafka - failed to close consumer: %s", err)
	}

	return nil
}
