package main

import (
	"log"
	"net/http"

	"github.com/raphacm/grpc-learning/service-b/internal/grpc"
	"github.com/raphacm/grpc-learning/service-b/internal/handlers"
	"github.com/raphacm/grpc-learning/service-b/internal/repositories"
	"github.com/raphacm/grpc-learning/service-b/internal/routers"
	"github.com/raphacm/grpc-learning/service-b/internal/services"
	"github.com/raphacm/grpc-learning/service-b/pkg/db"
)

func main() {
	database := db.New()
	todosRepositories := repositories.New(database)
	todosServices := services.New(todosRepositories)
	todoHandlers := handlers.New(todosServices)
	r := routers.New(todoHandlers)

	go grpc.NewgRPCServer(todosServices)

	e := r.Initialize()
	err := e.Start(":3003")
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
