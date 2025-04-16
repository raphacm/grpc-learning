package services

import (
	"context"

	"github.com/grpc-learning/service-a/internal/clients"
	"github.com/grpc-learning/service-a/internal/domains"
)

type ToDoServicer interface {
	CreateTodo(ctx context.Context, todo *domains.ToDoRequest) error
	GetTodo(ctx context.Context, todoID string) (*domains.ToDoDomain, error)
	GetTodos(ctx context.Context) ([]*domains.ToDoDomain, error)
}

type toDoService struct {
	client clients.ToDoClienter
}

func New(client clients.ToDoClienter) ToDoServicer {
	return &toDoService{
		client: client,
	}
}

func (service *toDoService) CreateTodo(ctx context.Context, todo *domains.ToDoRequest) error {
	return service.client.CreateTodo(ctx, todo)
}

func (service *toDoService) GetTodo(ctx context.Context, todoID string) (*domains.ToDoDomain, error) {
	return service.client.GetTodo(ctx, todoID)
}

func (service *toDoService) GetTodos(ctx context.Context) ([]*domains.ToDoDomain, error) {
	return service.client.GetTodos(ctx)
}
