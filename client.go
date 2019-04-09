package main

import (
	"context"
	"fmt"
	pb "go-grpc/todo"

	"google.golang.org/grpc"
)

const (
	address = "localhost:14586"
)

var (
	conn *grpc.ClientConn
	err  error
)

func open() {
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect :%s\n", err)
	}
}
func close() {
	err = conn.Close()
	if err != nil {
		fmt.Printf("close connect error :%s", err)
	}
}

func main() {
	open()
	defer close()
	// Get Grpc
	addTodo("3")
	getTodo()
}

func addTodo(task string) {
	c := pb.NewTodoServiceClient(conn)
	r1, err := c.AddTodo(context.Background(), &pb.AddTodoParams{Task: task})
	if err != nil {
		fmt.Printf("clond not addtodo ..\n", err)
		return
	}
	fmt.Printf("resp id is %s,\ntask is %s\n", r1.Id, r1.Task)
}

func getTodo() {
	c := pb.NewTodoServiceClient(conn)
	resp, err := c.GetTodo(context.Background(), &pb.GetTodoParams{})
	if err != nil {
		fmt.Printf("get todo error :%s\n", err)
	}
	fmt.Println("resp is ", resp)
}
