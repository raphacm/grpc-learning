package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/raphacm/grpc-learning/service-b/internal/handlers"
)

type Routers struct {
	toDoHandlers handlers.ToDoHandlER
}

func New(toDoHandler handlers.ToDoHandlER) *Routers {
	return &Routers{
		toDoHandlers: toDoHandler,
	}
}

func (router *Routers) Initialize() *echo.Echo {
	e := echo.New()

	e.POST("/v1/todos", router.toDoHandlers.CreateTodo)
	e.GET("/v1/todos/:todo_id", router.toDoHandlers.GetTodo)
	e.GET("/v1/todos", router.toDoHandlers.GetTodos)

	return e
}
