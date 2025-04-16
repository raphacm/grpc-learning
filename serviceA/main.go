package main

import (
	"log"
	"net/http"

	"github.com/grpc-learning/service-a/internal/clients"
	"github.com/grpc-learning/service-a/internal/handlers"
	"github.com/grpc-learning/service-a/internal/routers"
	"github.com/grpc-learning/service-a/internal/services"
	httpClient "github.com/grpc-learning/service-a/pkg/http"
)

func main() {
	todosClients := clients.New(httpClient.NewClient(http.Client{}), "http://localhost:3003")
	todosServices := services.New(todosClients)
	todosHandlers := handlers.New(todosServices)
	r := routers.New(todosHandlers)

	e := r.Initialize()
	err := e.Start(":3002")
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
