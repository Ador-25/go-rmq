package main

import (
	"fmt"
	"log"
	"net/http"
	"rmq/middlewares"
	"rmq/rabbitMQ"
	"rmq/sse"
)

var (
	routingKey = "go-test"
	queueName  = fmt.Sprintf("%s:queue", routingKey)
)

func main() {
	run_rmq()
	go func() {
		r := http.NewServeMux()
		r.HandleFunc("/events", sse.EventsHandler)
		r.HandleFunc("/send", middlewares.Post(Send))
		viewsDir := http.Dir("./views")
		r.Handle("/", http.StripPrefix("/", http.FileServer(viewsDir)))
		log.Fatal(http.ListenAndServe(":8080", r))
	}()
	select {}
}

func run_rmq() {
	go rabbitMQ.Consume(queueName, routingKey)
	fmt.Println("Consumers started. RUNNING !@#$$%%^")
}
func Send(rw http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	fmt.Println(msg)
	rabbitMQ.Publish([]byte(msg), routingKey)
	rw.Write([]byte(fmt.Sprintf("message published to rmq")))
}
