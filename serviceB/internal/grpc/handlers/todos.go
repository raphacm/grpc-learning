package grpc_handler

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/grpc-learning/service-b/internal/domains"

	// pb "github.com/grpc-learning/service-b/internal/grpc/gen/github.com/grpc-learning/service-b/gen/proto/v1"
	pb "github.com/grpc-learning/grpc-protos/serviceB/gen/proto/v1"
	"github.com/grpc-learning/service-b/internal/services"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodosServer struct {
	pb.UnimplementedTodosServiceServer
	Service services.ToDoServicer
}

func (server *TodosServer) CreateTodo(ctx context.Context, req *pb.TodosRequest) (*emptypb.Empty, error) {

	err := server.Service.CreateTodo(ctx, mapperTodosRequestToServiceDomain(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (server *TodosServer) GetTodo(ctx context.Context, id *pb.GetTodoIdParams) (*pb.TodosResponse, error) {
	todo, err := server.Service.GetTodo(ctx, id.Id)
	if err != nil {
		return nil, err
	}
	return mapTodoDomainTogRPCTodosResponse(todo), nil
}

func (server *TodosServer) GetTodos(ctx context.Context, req *emptypb.Empty) (*pb.ListTodosResponse, error) {
	todos := server.Service.GetTodos(ctx)

	listOfTodos := &pb.ListTodosResponse{}

	for _, todo := range todos {
		listOfTodos.Todos = append(listOfTodos.Todos, mapTodoDomainTogRPCTodosResponse(todo))
	}

	return listOfTodos, nil
}

func mapTodoDomainTogRPCTodosResponse(domain *domains.ToDoDomain) *pb.TodosResponse {
	return &pb.TodosResponse{
		Id:        domain.ID,
		Task:      domain.Task,
		IsDone:    domain.IsDone,
		CreatedAt: domain.Created_at,
	}
}

func mapperTodosRequestToServiceDomain(req *pb.TodosRequest) domains.ToDoDomain {
	return domains.ToDoDomain{
		ID:         uuid.NewString(),
		Task:       req.Task,
		IsDone:     req.IsDone,
		Created_at: time.Now().UnixMilli(),
	}
}
