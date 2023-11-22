package main

import (
	"fmt"

	"log"

	"github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func init() {
	initAmqp()
}

var conn *amqp091.Connection
var ch *amqp091.Channel
var replies <-chan amqp091.Delivery

func initAmqp() {
	var err error
	var q amqp091.Queue

	conn, err = amqp091.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	log.Printf("got Connection, getting Channel...")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")

	log.Printf("got Channel, declaring Exchange (%s)", "go-test-exchange")

	err = ch.ExchangeDeclare(
		"booking-result", // name of the exchange
		"direct",         // type
		true,             // durable
		false,            // delete when complete
		false,            // internal
		false,            // noWait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare the Exchange")

	log.Printf("declared Exchange, declaring Queue (%s)", "go-test-queue")

	q, err = ch.QueueDeclare(
		"go-test-queue", // name, leave empty to generate a unique name
		true,            // durable
		false,           // delete when usused
		false,           // exclusive
		false,           // noWait
		nil,             // arguments
	)
	failOnError(err, "Error declaring the Queue")

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		q.Name, q.Messages, q.Consumers, "go-test-key")

	err = ch.QueueBind(
		q.Name,           // name of the queue
		"go-booking-key", // bindingKey
		"booking-result", // sourceExchange
		false,            // noWait
		nil,              // arguments
	)
	failOnError(err, "Error binding to the Queue")

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", "go-amqp-example")

	replies, err = ch.Consume(
		q.Name,            // queue
		"go-amqp-example", // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	failOnError(err, "Error consuming the Queue")
}

func main() {
	defer ch.Close()
	defer conn.Close()

	log.Println("Start consuming the Queue...")
	var count int = 1

	for r := range replies {
		log.Printf("Consuming reply number %d", count)
		fmt.Printf("Message: %s\n", r.Body)
		count++
	}
}
