package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/grpc-learning/service-a/internal/domains"
	"github.com/grpc-learning/service-a/internal/services"
	"github.com/labstack/echo/v4"
)

type ToDoHandlER interface {
	CreateTodo(c echo.Context) error
	GetTodo(c echo.Context) error
	GetTodos(c echo.Context) error
}

type todoHandler struct {
	service services.ToDoServicer
}

func New(service services.ToDoServicer) ToDoHandlER {
	return &todoHandler{
		service: service,
	}
}

func (handler *todoHandler) CreateTodo(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()
	todoRequest := &domains.ToDoRequest{}
	err := c.Bind(todoRequest)
	if err != nil {
		fmt.Println("service-a: create todo bind error: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	err = handler.service.CreateTodo(ctx, todoRequest)
	if err != nil {
		fmt.Println("service-a: create todo service error: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Created",
	})
}

func (handler *todoHandler) GetTodo(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	todoID := c.Param("todo_id")
	if todoID == "" {
		fmt.Println("service-a: get todo id error")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "todo id is required",
		})
	}

	todo, err := handler.service.GetTodo(ctx, todoID)
	if err != nil {
		fmt.Println("service-a: get todo service error: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, todo)
}

func (handler *todoHandler) GetTodos(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	todos, err := handler.service.GetTodos(ctx)
	if err != nil {
		fmt.Println("service-a: get todos service error: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, todos)
}
