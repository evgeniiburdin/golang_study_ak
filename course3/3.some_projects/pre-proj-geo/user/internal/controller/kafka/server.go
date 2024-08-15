package kafka

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
	"user-service/internal/usecase"
	kafka_pkg "user-service/pkg/kafka"
	"user-service/pkg/logger"
)

var topics []string = []string{
	"CreateUser",
}

type KafkaController struct {
	uc usecase.Userer
	lg logger.Interface
	c  kafka_pkg.KafkaConsumer
}

func New(uc usecase.Userer, lg logger.Interface, c kafka_pkg.KafkaConsumer) *KafkaController {
	return &KafkaController{
		uc: uc,
		lg: lg,
		c:  c,
	}
}

func (kc *KafkaController) Start(attempts int, interval time.Duration) error {
	var err error

	for attempts > 0 {
		// Subscribe to the Kafka topic
		if attempts == 0 {
			return fmt.Errorf("kafka controller - failed to subscribe to topics: %+v, error: %w", topics, err)
		}

		err = kc.c.Consumer.SubscribeTopics(topics, nil)
		if err != nil {
			kc.lg.Error(fmt.Errorf("kafka controller - failed to subscribe to topics: %w, attempts left: %d", err, attempts))
		}
		attempts--

		time.Sleep(interval)
	}

	// Start consuming messages
	run := true
	for run == true {
		select {
		case sig := <-kc.c.NotifyChan:
			kc.lg.Info(fmt.Sprintf("kafka controller - received signal %v: terminating\n", sig))
			fmt.Printf("kafka controller - received signal %v: terminating\n", sig)
			run = false
		default:
			// Poll for Kafka messages
			ev := kc.c.Consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				errr := kc.uc.Write(context.Background(), e.Value)
				if errr != nil {
					kc.lg.Error(fmt.Errorf("kafka controller - failed to write user: %w", errr))
				}
			case kafka.Error:
				// Handle Kafka errors
				kc.lg.Info(fmt.Sprintf("kafka controller - error: %+v", e))
			}
		}
	}

	return nil
}
