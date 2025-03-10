package main

import (
	"context"
	"fmt"
	"log"
	"logger/data"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRPCPort = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: *data.New(client),
	}

	err = rpc.Register(new(RPCServer))
	if err != nil {
		panic(err)
	}
	go app.listenRPC()

	go app.gRPCListen()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Println("Starting server")
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func connectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting to mongo")
		return nil, err
	}

	log.Println("Connected to mongo db")
	return conn, nil
}

func (app *Config) listenRPC() error {
	log.Println("starting rpc server on port: " + rpcPort)

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		return err
	}

	for {
		rpcconn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcconn)
	}
}
