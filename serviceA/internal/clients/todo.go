package clients

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/grpc-learning/service-a/internal/domains"
	"github.com/grpc-learning/service-a/pkg/http"
)

type ToDoClienter interface {
	CreateTodo(ctx context.Context, todo *domains.ToDoRequest) error
	GetTodos(ctx context.Context) ([]*domains.ToDoDomain, error)
	GetTodo(ctx context.Context, todoID string) (*domains.ToDoDomain, error)
}

func New(http http.HttpClienter, baseURL string) ToDoClienter {
	return &toDoClient{
		http:    http,
		baseURL: baseURL,
	}
}

type toDoClient struct {
	http    http.HttpClienter
	baseURL string
}

func (client *toDoClient) CreateTodo(ctx context.Context, todo *domains.ToDoRequest) error {
	url := fmt.Sprintf("%s/v1/todos", client.baseURL)
	fmt.Println(url)
	body, _ := json.Marshal(todo)

	response, err := client.http.Post(ctx, &http.HttpClientConfig{
		URL:     url,
		Headers: nil,
	}, body)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(response.StatusCode)
	if response.StatusCode != 201 {
		return fmt.Errorf("error when call todo API to post a task")
	}

	return nil
}

func (client *toDoClient) GetTodos(ctx context.Context) ([]*domains.ToDoDomain, error) {
	url := fmt.Sprintf("%s/v1/todos", client.baseURL)

	response, err := client.http.Get(ctx, &http.HttpClientConfig{
		URL:     url,
		Headers: nil,
	})

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error when call todo API to get a task by id")
	}

	todoDomain := []*domains.ToDoDomain{}

	err = json.Unmarshal(response.Body, &todoDomain)
	if err != nil {
		return nil, fmt.Errorf("error when unmarshall response body to get task")
	}

	return todoDomain, nil
}

func (client *toDoClient) GetTodo(ctx context.Context, todoID string) (*domains.ToDoDomain, error) {
	url := fmt.Sprintf("%s/v1/todos/%s", client.baseURL, todoID)

	response, err := client.http.Get(ctx, &http.HttpClientConfig{
		URL:     url,
		Headers: nil,
	})

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error when call todo API to get a task by id")
	}

	todoDomain := &domains.ToDoDomain{}

	err = json.Unmarshal(response.Body, todoDomain)
	if err != nil {
		return nil, fmt.Errorf("error when unmarshall response body to get task")
	}

	return todoDomain, nil
}
