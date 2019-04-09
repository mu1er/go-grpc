package main

import (
	"fmt"
	pb "go-grpc/todo"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", 14586))
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}
	s := pb.Server{}
	grpcServer := grpc.NewServer()

	pb.RegisterTodoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server :%s", err)
	} else {
		log.Printf("Server started successfully")
	}
}
