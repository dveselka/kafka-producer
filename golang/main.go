package main

import (
	"fmt"

	"producer"
)

func main() {
	// send message to Kafka
	message := producer.Produce()
	fmt.Println(message)
}
