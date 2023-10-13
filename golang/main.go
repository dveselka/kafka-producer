package main

import (
	"fmt"

	"github.com/dveselka/kafka-producer/golang/producer"
)

func main() {
	// send message to Kafka
	message := producer.Produce()
	fmt.Println(message)
}
