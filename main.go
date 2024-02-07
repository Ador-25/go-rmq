package main

import (
	"fmt"
	rabbitMQ "rmq/rmq"
)

func main() {
	routingKey := "go-test"
	queueName := fmt.Sprintf("%s:queue", routingKey)
	rabbitMQ.Consume(queueName, routingKey)
	// message := []byte("Hello, World!")
	// Publish(message, routingKey)
}
