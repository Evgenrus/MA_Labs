package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
)

func rabbit() {
	conn, err := amqp.Dial("amqp://rabbitmq")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cha, err := conn.Channel()
	if err != nil {
		log.Fatal("Error in creating channel")
	}
	defer cha.Close()

	messages, err := cha.Consume(
		"lab3",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Println(string(message.Body))
		}
	}()

	<-forever
}

func main() {
	go rabbit()

	http.HandleFunc("/service2/getString", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Service2")
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8084", nil))
}
