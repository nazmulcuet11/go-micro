package main

import (
	"fmt"
	"listener/event"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitconn, err := connect()
	if err != nil {
		log.Panic(err)
	}
	defer rabbitconn.Close()

	consumer, err := event.NewConsumer(rabbitconn)
	if err != nil {
		panic(err)
	}

	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	count := 0
	var connection *amqp.Connection
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			count++
			if count > 10 {
				fmt.Println("Could not connect to rabbitmq")
				return nil, err
			} else {
				fmt.Println("RabbitMQ not yet ready, backingOff for 1 sec")
				time.Sleep(1 * time.Second)
			}
		} else {
			log.Println("connected to rabbitmq")
			connection = c
			break
		}
	}
	return connection, nil
}
