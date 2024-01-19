package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// Here we set the way error messages are displayed in the terminal.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What message do you want to send?")
	mPayload, _ := reader.ReadString('\n')

	// Here we connect to RabbitMQ or send a message if there are any errors connecting.
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// We create a Queue to send the message to.
	q, err := ch.QueueDeclare(
		"golang-queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// We set the payload for the message.
	// body := "Golang is awesome - Keep Moving Forward!"
	body := mPayload
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	// If there is an error publishing the message, a log will be displayed in the terminal.
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Congrats, sending message: %s", body)
}
