package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	Configuration Config
	Conn          *amqp.Connection
	Ch            *amqp.Channel
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RMQ_SECRET struct {
	HostName string `json:"hostname"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Exchange string `json:"exchange"`
}
type Config struct {
	Creds Credentials `json:"creds"`
	RMQ   RMQ_SECRET  `json:"RabbitMQConnection"`
}

func GetConnectionString() string {
	rabbitMQEndpoint := Configuration.RMQ.HostName
	rabbitMQUsername := Configuration.RMQ.UserName
	rabbitMQPassword := Configuration.RMQ.Password
	str := fmt.Sprintf("amqp://%s:%s@%s", rabbitMQUsername, rabbitMQPassword, rabbitMQEndpoint)
	return str
}

func init() {
	data, err := ioutil.ReadFile("applicationsettings.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err = json.Unmarshal(data, &Configuration)
	if err != nil {
		log.Fatalf("Error parsing config JSON: %v", err)
	}
	fmt.Println("Username:", Configuration.Creds)
	fmt.Println("RMQ:", Configuration.RMQ)
	amqpURI := GetConnectionString()
	conn, err := amqp.Dial(amqpURI)
	FailOnError(err, "Failed to connect to RabbitMQ")
	Ch, err = conn.Channel()
	FailOnError(err, "Failed to open a channel")
}
