package kafka

import (
	"auth-service/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func New(kafkaServiceAddress string, attempts int, interval time.Duration, lg logger.Interface) (*KafkaProducer, error) {
	for attempts > 0 {
		p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServiceAddress})
		if err != nil {
			lg.Error(fmt.Errorf("kafka - failed to create Kafka producer: %w, attempts left: %d", err, attempts))
			attempts--
			time.Sleep(interval)
			continue
		}
		return &KafkaProducer{producer: p}, nil
	}

	return nil, fmt.Errorf("failed to create Kafka producer")
}

func (kp *KafkaProducer) SerializeAndProduce(message interface{}, topic string) error {
	// Serialize the message struct to JSON
	serializedMessage, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("kafka - failed to serialize message %#v: %w", message, err)
	}

	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          serializedMessage,
	}

	// Produce the Kafka message
	deliveryChan := make(chan kafka.Event)
	err = kp.producer.Produce(kafkaMessage, deliveryChan)
	if err != nil {
		return fmt.Errorf("kafka - failed to produce message %#v: %w", kafkaMessage, err)
	}
	// Wait for delivery report or error
	e := <-deliveryChan
	m := e.(*kafka.Message)

	// Check for delivery errors
	if m.TopicPartition.Error != nil {
		return fmt.Errorf("delivery failed: %s", m.TopicPartition.Error)
	}

	// Close the delivery channel
	close(deliveryChan)

	return nil
}

func (kp *KafkaProducer) Close() error {
	err := kp.Close()
	if err != nil {
		return fmt.Errorf("failed to close Kafka producer: %w", err)
	}

	return nil
}
