package utils

import (
	"log"
	"bd2-backend/src/config"
	"github.com/streadway/amqp"
)

var amqpURI string
var channel *amqp.Channel
var queue amqp.Queue

func init() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		ErrorLogger.Fatal("cannot load config:", err)
	}
	amqpURI = cfg.AmqpURI

	conn, err := amqp.Dial(amqpURI)
    failOnError(err, "Failed to connect to RabbitMQ")
    //defer conn.Close()

	ch, err := conn.Channel()
	//defer ch.Close()
	if err != nil {
		ErrorLogger.Fatal("Failed to open a channel:", err)
	}
	channel = ch

    q, err := ch.QueueDeclare(
        "notif", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
	if err != nil {
		ErrorLogger.Fatal("Failed to declare a queue:", err)
	}
	queue = q

}

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}

func AddmessageToQueue(body []byte) error {


	err := channel.Publish(
        "",     // exchange
        queue.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
	
	if err != nil {
		ErrorLogger.Fatal("Failed to publish a message:", err)
		return err
	}
	
	return nil
}