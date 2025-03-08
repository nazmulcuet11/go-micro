package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "8080"

type Config struct {
	rabbit *amqp.Connection
}

func main() {
	rabbitconn, err := connect()
	if err != nil {
		log.Panic(err)
	}
	defer rabbitconn.Close()

	app := Config{
		rabbit: rabbitconn,
	}

	log.Printf("Starting broker service on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
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
