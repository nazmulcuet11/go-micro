package main

import (
	"context"
	"fmt"
	"log"
	"logger/data"
	"logger/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		return &logs.LogResponse{Result: "Failed"}, err
	}
	return &logs.LogResponse{Result: "Success"}, nil
}

func (app *Config) gRPCListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	logs.RegisterLogServiceServer(
		s,
		&LogServer{Models: app.Models},
	)
	log.Println("grpc server started")

	err = s.Serve(listen)
}
