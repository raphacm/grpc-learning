package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raphacm/grpc-learning/service-b/internal/domains"
	"github.com/raphacm/grpc-learning/service-b/internal/services"
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
		fmt.Println("service-b: create todo bind error: ", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	err = handler.service.CreateTodo(ctx, todoRequest.ToToDoDomain())
	if err != nil {
		fmt.Println("service-b: create todo service error: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	c.Response().Status = 201
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
		fmt.Println("service-b: get todo id error")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "todo_id field is required",
		})
	}

	todo, err := handler.service.GetTodo(ctx, todoID)
	if err != nil {
		fmt.Println("service-b: get todo service error: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, todo)
}

func (handler *todoHandler) GetTodos(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	todos := handler.service.GetTodos(ctx)

	return c.JSON(http.StatusOK, todos)
}
