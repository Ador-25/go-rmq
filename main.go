package main

import "fmt"

func main() {
	routingKey := "go-test"
	queueName := fmt.Sprintf("%s:queue", routingKey)
	Consume(queueName, routingKey)
	// message := []byte("Hello, World!")
	// Publish(message, routingKey)
}
