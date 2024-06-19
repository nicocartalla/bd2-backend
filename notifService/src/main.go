package main

import (
    "log"
	"notifService/src/config"
	"notifService/src/services"
	"notifService/src/utils"
)

var amqpURI string


func init() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		utils.ErrorLogger.Fatal("cannot load config:", err)
	}
	amqpURI = cfg.AmqpURI
}

func main() {
    conn, ch, err := utils.ConnectToRabbitMQ(amqpURI)
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()
    defer ch.Close()

    queueName := "notif"
    msgs, err := utils.ConsumeMessages(ch, queueName)
    if err != nil {
        log.Fatalf("Failed to register a consumer: %v", err)
    }

    utils.InfoLogger.Println("Successfully connected to RabbitMQ")
    utils.InfoLogger.Println("Waiting for messages...")

    services.ProcessMessages(msgs)
}