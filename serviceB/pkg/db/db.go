package db

import (
	"context"
	"fmt"

	"github.com/grpc-learning/service-b/internal/domains"
)

type Databaser interface {
	Insert(ctx context.Context, data domains.ToDoDomain) error
	Get(ctx context.Context, id string) (*domains.ToDoDomain, error)
	List(ctx context.Context) []*domains.ToDoDomain
}

type database struct {
	table map[string]domains.ToDoDomain
}

func New() Databaser {
	return &database{
		table: make(map[string]domains.ToDoDomain, 0),
	}
}

func (db *database) Insert(ctx context.Context, data domains.ToDoDomain) error {
	db.table[data.ID] = data
	return nil
}

func (db *database) Get(ctx context.Context, id string) (*domains.ToDoDomain, error) {
	value, ok := db.table[id]
	if !ok {
		return nil, fmt.Errorf("todo not found")
	}

	return &value, nil
}

func (db *database) List(ctx context.Context) []*domains.ToDoDomain {
	list := make([]*domains.ToDoDomain, 0)
	for _, value := range db.table {
		list = append(list, &value)
	}

	return list
}
