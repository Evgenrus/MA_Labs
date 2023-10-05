package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"time"
)

func main() {
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

	_, err = cha.QueueDeclare(
		"lab3", // queue name
		true,   // durable
		false,  // auto delete
		false,  // exclusive
		false,  // no wait
		nil,    // arguments
	)

	http.HandleFunc("/service1/getString", func(w http.ResponseWriter, r *http.Request) {
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(time.Now().GoString()),
		}
		if err := cha.Publish("", "lab3", false, false, message); err != nil {
			fmt.Fprintf(w, err.Error())
		}
		fmt.Fprintf(w, "Service1")
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}
