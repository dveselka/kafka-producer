package main

import (
	"log"

	"github.com/Shopify/sarama"
)

const (
	// KafkaConnectionString hostname:port
	KafkaConnectionString = "localhost:9092"

	// KafkaTopic topic name
	KafkaTopic = "test-topic"
)

func main() {
	// create prodcuer
	producer, err := sarama.NewSyncProducer([]string{KafkaConnectionString}, nil)

	// check if connected
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to %s", KafkaConnectionString)

	// close connection
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// send message
	msg := &sarama.ProducerMessage{Topic: KafkaTopic, Value: sarama.StringEncoder("testing 123")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}

	log.Print("Done")
}
