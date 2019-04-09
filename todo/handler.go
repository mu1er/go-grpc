package todo

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
)

type Server struct {
	Todos []*TodoObject
}

func (s *Server) AddTodo(ctx context.Context, newTodo *AddTodoParams) (*TodoObject, error) {
	log.Printf("Received new task %s", newTodo.Task)
	todoObject := &TodoObject{
		Id:   uuid.NewV1().String(),
		Task: newTodo.Task,
	}
	s.Todos = append(s.Todos, todoObject)
	return todoObject, nil
}

func (s *Server) GetTodo(ctx context.Context, _ *GetTodoParams) (*TodoResponse, error) {
	log.Printf("get tasks")
	return &TodoResponse{Todos: s.Todos}, nil
}

func (s *Server) DelTodo(ctx context.Context, delTodo *DelTodoParams) (*DelResponse, error) {
	log.Printf("delete todo id %s", delTodo.Id)
	var updateTodos []*TodoObject
	for index, todo := range s.Todos {
		if delTodo.Id == todo.Id {
			updateTodos = append(s.Todos[:index], s.Todos[index+1:]...)
			break
		}
	}
	s.Todos = updateTodos
	return &DelResponse{Message: "success"}, nil
}
