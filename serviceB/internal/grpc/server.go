package grpc

import (
	"log"
	"net"

	// todo_proto "github.com/grpc-learning/service-b/internal/grpc/gen/github.com/grpc-learning/service-b/gen/proto/v1"
	todo_proto "github.com/grpc-learning/grpc-protos/serviceB/gen/proto/v1"
	grpc_handler "github.com/grpc-learning/service-b/internal/grpc/handlers"
	"github.com/grpc-learning/service-b/internal/services"
	"google.golang.org/grpc"
)

func NewgRPCServer(service services.ToDoServicer) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register your service
	todo_proto.RegisterTodosServiceServer(grpcServer, &grpc_handler.TodosServer{
		Service: service,
	})

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
