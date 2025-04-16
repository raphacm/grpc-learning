// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: v1/todos.proto

package todo_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TodosService_GetTodos_FullMethodName   = "/todo_proto.v1.TodosService/GetTodos"
	TodosService_GetTodo_FullMethodName    = "/todo_proto.v1.TodosService/GetTodo"
	TodosService_CreateTodo_FullMethodName = "/todo_proto.v1.TodosService/CreateTodo"
)

// TodosServiceClient is the client API for TodosService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodosServiceClient interface {
	GetTodos(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListTodosResponse, error)
	GetTodo(ctx context.Context, in *GetTodoIdParams, opts ...grpc.CallOption) (*TodosResponse, error)
	CreateTodo(ctx context.Context, in *TodosRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type todosServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodosServiceClient(cc grpc.ClientConnInterface) TodosServiceClient {
	return &todosServiceClient{cc}
}

func (c *todosServiceClient) GetTodos(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListTodosResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTodosResponse)
	err := c.cc.Invoke(ctx, TodosService_GetTodos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosServiceClient) GetTodo(ctx context.Context, in *GetTodoIdParams, opts ...grpc.CallOption) (*TodosResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodosResponse)
	err := c.cc.Invoke(ctx, TodosService_GetTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosServiceClient) CreateTodo(ctx context.Context, in *TodosRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TodosService_CreateTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodosServiceServer is the server API for TodosService service.
// All implementations must embed UnimplementedTodosServiceServer
// for forward compatibility.
type TodosServiceServer interface {
	GetTodos(context.Context, *emptypb.Empty) (*ListTodosResponse, error)
	GetTodo(context.Context, *GetTodoIdParams) (*TodosResponse, error)
	CreateTodo(context.Context, *TodosRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTodosServiceServer()
}

// UnimplementedTodosServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTodosServiceServer struct{}

func (UnimplementedTodosServiceServer) GetTodos(context.Context, *emptypb.Empty) (*ListTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodos not implemented")
}
func (UnimplementedTodosServiceServer) GetTodo(context.Context, *GetTodoIdParams) (*TodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}
func (UnimplementedTodosServiceServer) CreateTodo(context.Context, *TodosRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodosServiceServer) mustEmbedUnimplementedTodosServiceServer() {}
func (UnimplementedTodosServiceServer) testEmbeddedByValue()                      {}

// UnsafeTodosServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodosServiceServer will
// result in compilation errors.
type UnsafeTodosServiceServer interface {
	mustEmbedUnimplementedTodosServiceServer()
}

func RegisterTodosServiceServer(s grpc.ServiceRegistrar, srv TodosServiceServer) {
	// If the following call pancis, it indicates UnimplementedTodosServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TodosService_ServiceDesc, srv)
}

func _TodosService_GetTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServiceServer).GetTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodosService_GetTodos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServiceServer).GetTodos(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodosService_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoIdParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServiceServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodosService_GetTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServiceServer).GetTodo(ctx, req.(*GetTodoIdParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodosService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodosService_CreateTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServiceServer).CreateTodo(ctx, req.(*TodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodosService_ServiceDesc is the grpc.ServiceDesc for TodosService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodosService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo_proto.v1.TodosService",
	HandlerType: (*TodosServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTodos",
			Handler:    _TodosService_GetTodos_Handler,
		},
		{
			MethodName: "GetTodo",
			Handler:    _TodosService_GetTodo_Handler,
		},
		{
			MethodName: "CreateTodo",
			Handler:    _TodosService_CreateTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/todos.proto",
}
