package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaService interface {
	WriteMessage(msg, topic string) error
}

type kafkaService struct {
	broker string
}

func NewKafkaService(broker string) KafkaService {
	fmt.Println("Kafka broker: ", broker)
	return kafkaService{
		broker: broker,
	}
}

func (ks kafkaService) WriteMessage(msg, topic string) error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", ks.broker, topic, 0)
	if err != nil {
		log.Println("failed to dial leader:", err)
		return err
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(msg)},
	)
	if err != nil {
		log.Println("failed to write messages:", err)
		return err
	}

	if err := conn.Close(); err != nil {
		log.Println("failed to close writer:", err)
		return err
	}
	return nil
}

func StartConsumer(broker string, topic string) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
