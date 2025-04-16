package repositories

import (
	"context"
	"time"

	"github.com/grpc-learning/service-b/internal/domains"
	"github.com/grpc-learning/service-b/pkg/db"
)

type ToDoRepositorer interface {
	CreateTodo(ctx context.Context, todo domains.ToDoDomain) error
	GetTodos(ctx context.Context) []*domains.ToDoDomain
	GetTodo(ctx context.Context, todoID string) (*domains.ToDoDomain, error)
}

func New(db db.Databaser) ToDoRepositorer {
	return &toDoRepository{
		db: db,
	}
}

type toDoRepository struct {
	db db.Databaser
}

func (repo *toDoRepository) CreateTodo(ctx context.Context, todo domains.ToDoDomain) error {
	time.Sleep(time.Millisecond * 15)
	err := repo.db.Insert(ctx, todo)
	if err != nil {
		return err
	}
	return nil
}

func (repo *toDoRepository) GetTodos(ctx context.Context) []*domains.ToDoDomain {
	time.Sleep(time.Millisecond * 15)
	return repo.db.List(ctx)

}

func (repo *toDoRepository) GetTodo(ctx context.Context, todoID string) (*domains.ToDoDomain, error) {
	time.Sleep(time.Millisecond * 15)
	todo, err := repo.db.Get(ctx, todoID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
