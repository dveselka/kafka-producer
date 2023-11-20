package main

import (
	"fmt"

	"github.com/dveselka/kafka-producer/golang/producer"
)

func main() {
	// send message to Kafka
	message := Producer.Produce()
	fmt.Println(message)
}
