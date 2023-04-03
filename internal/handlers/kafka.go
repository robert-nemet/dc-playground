package handlers

import (
	"dc-playground/internal/services"
	"io/ioutil"
	"log"
	"net/http"
)

type KafkaHandlers interface {
	KafkaHandler(w http.ResponseWriter, r *http.Request)
}

type kafkaHandler struct {
	kafka      services.KafkaService
	kafkaTopic string
}

func NewKafkaHandler(svc services.KafkaService, kafkaTopic string) KafkaHandlers {
	return &kafkaHandler{
		kafka:      svc,
		kafkaTopic: kafkaTopic,
	}
}

func (k *kafkaHandler) KafkaHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[Kafka producer error, reading message] ", err)
	}
	err = k.kafka.WriteMessage(string(msg), k.kafkaTopic)
	if err != nil {
		log.Println("[Kafka producer error, writing message] ", err)
	}
}
