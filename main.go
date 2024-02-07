package main

import (
	"fmt"
	"log"
	"net/http"
	"rmq/rabbitMQ"
	"rmq/sse"
)

var (
	routingKey = "go-test"
	queueName  = fmt.Sprintf("%s:queue", routingKey)
)

func main() {
	// message := []byte("Hello, World!")
	// Publish(message, routingKey)
	run_rmq()
	fmt.Println("back to main routine")
	r := http.NewServeMux()
	r.HandleFunc("/events", sse.EventsHandler)

	viewsDir := http.Dir("./views")
	r.Handle("/", http.StripPrefix("/", http.FileServer(viewsDir)))
	log.Fatal(http.ListenAndServe(":8080", r))

	select {}
}
func rmq_publish(rw http.ResponseWriter, r *http.Request) {
	str := r.URL.Query().Get("id")
	data := []byte(str)
	rabbitMQ.Publish([]byte(data), routingKey)
}
func run_rmq() {
	go rabbitMQ.Consume(queueName, routingKey)
	fmt.Println("Consumers started. RUNNING !@#$$%%^")
}
