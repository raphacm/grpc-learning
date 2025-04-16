package services

import (
	"context"

	"github.com/raphacm/grpc-learning/service-b/internal/domains"
	"github.com/raphacm/grpc-learning/service-b/internal/repositories"
)

type ToDoServicer interface {
	CreateTodo(ctx context.Context, todo domains.ToDoDomain) error
	GetTodo(ctx context.Context, todoID string) (*domains.ToDoDomain, error)
	GetTodos(ctx context.Context) []*domains.ToDoDomain
}

type toDoService struct {
	repo repositories.ToDoRepositorer
}

func New(repo repositories.ToDoRepositorer) ToDoServicer {
	return &toDoService{
		repo: repo,
	}
}

func (service *toDoService) CreateTodo(ctx context.Context, todo domains.ToDoDomain) error {
	return service.repo.CreateTodo(ctx, todo)
}

func (service *toDoService) GetTodo(ctx context.Context, todoID string) (*domains.ToDoDomain, error) {
	return service.repo.GetTodo(ctx, todoID)
}

func (service *toDoService) GetTodos(ctx context.Context) []*domains.ToDoDomain {
	return service.repo.GetTodos(ctx)
}
