package rabbitMQ

import (
	"log"
	"rmq/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(message []byte, routingKey string) {
	exchange := Configuration.RMQ.Exchange
	conn, err := amqp.Dial(GetConnectionString())
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", message)
}
