
package utils

import (
    "github.com/streadway/amqp"
)

func ConnectToRabbitMQ(url string) (*amqp.Connection, *amqp.Channel, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, nil, err
    }

    return conn, ch, nil
}

func ConsumeMessages(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
    q, err := ch.QueueDeclare(
        queueName, // name
        false,     // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        return nil, err
    }

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        return nil, err
    }

    return msgs, nil
}
