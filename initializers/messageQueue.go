package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

var Conn *amqp091.Connection
var Ch *amqp091.Channel

func ConnectToQueue() {
	var err error

	Conn, err = amqp091.Dial(os.Getenv("MQ_URI"))
	FailOnError(err, "Failed to connect to RabbitMQ")

	Ch, err = Conn.Channel()
	FailOnError(err, "Failed to open a channel")

	err = Ch.ExchangeDeclare(
		"booking-result", // name
		"direct",         // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // noWait
		nil,              // arguments
	)
	FailOnError(err, "Failed to declare the Exchange")
}
