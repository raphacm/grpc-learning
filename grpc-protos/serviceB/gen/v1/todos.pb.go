// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: v1/todos.proto

package todo_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TodosRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          string                 `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	IsDone        bool                   `protobuf:"varint,2,opt,name=is_done,json=isDone,proto3" json:"is_done,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodosRequest) Reset() {
	*x = TodosRequest{}
	mi := &file_v1_todos_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodosRequest) ProtoMessage() {}

func (x *TodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_todos_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodosRequest.ProtoReflect.Descriptor instead.
func (*TodosRequest) Descriptor() ([]byte, []int) {
	return file_v1_todos_proto_rawDescGZIP(), []int{0}
}

func (x *TodosRequest) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *TodosRequest) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

type TodosResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Task          string                 `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	IsDone        bool                   `protobuf:"varint,3,opt,name=is_done,json=isDone,proto3" json:"is_done,omitempty"`
	CreatedAt     int64                  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodosResponse) Reset() {
	*x = TodosResponse{}
	mi := &file_v1_todos_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodosResponse) ProtoMessage() {}

func (x *TodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_todos_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodosResponse.ProtoReflect.Descriptor instead.
func (*TodosResponse) Descriptor() ([]byte, []int) {
	return file_v1_todos_proto_rawDescGZIP(), []int{1}
}

func (x *TodosResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TodosResponse) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *TodosResponse) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *TodosResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type GetTodoIdParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTodoIdParams) Reset() {
	*x = GetTodoIdParams{}
	mi := &file_v1_todos_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTodoIdParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoIdParams) ProtoMessage() {}

func (x *GetTodoIdParams) ProtoReflect() protoreflect.Message {
	mi := &file_v1_todos_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoIdParams.ProtoReflect.Descriptor instead.
func (*GetTodoIdParams) Descriptor() ([]byte, []int) {
	return file_v1_todos_proto_rawDescGZIP(), []int{2}
}

func (x *GetTodoIdParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListTodosResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Todos         []*TodosResponse       `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTodosResponse) Reset() {
	*x = ListTodosResponse{}
	mi := &file_v1_todos_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodosResponse) ProtoMessage() {}

func (x *ListTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_todos_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodosResponse.ProtoReflect.Descriptor instead.
func (*ListTodosResponse) Descriptor() ([]byte, []int) {
	return file_v1_todos_proto_rawDescGZIP(), []int{3}
}

func (x *ListTodosResponse) GetTodos() []*TodosResponse {
	if x != nil {
		return x.Todos
	}
	return nil
}

var File_v1_todos_proto protoreflect.FileDescriptor

const file_v1_todos_proto_rawDesc = "" +
	"\n" +
	"\x0ev1/todos.proto\x12\rtodo_proto.v1\x1a\x1bgoogle/protobuf/empty.proto\";\n" +
	"\fTodosRequest\x12\x12\n" +
	"\x04task\x18\x01 \x01(\tR\x04task\x12\x17\n" +
	"\ais_done\x18\x02 \x01(\bR\x06isDone\"k\n" +
	"\rTodosResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04task\x18\x02 \x01(\tR\x04task\x12\x17\n" +
	"\ais_done\x18\x03 \x01(\bR\x06isDone\x12\x1d\n" +
	"\n" +
	"created_at\x18\x04 \x01(\x03R\tcreatedAt\"!\n" +
	"\x0fGetTodoIdParams\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"G\n" +
	"\x11ListTodosResponse\x122\n" +
	"\x05todos\x18\x01 \x03(\v2\x1c.todo_proto.v1.TodosResponseR\x05todos2\xe0\x01\n" +
	"\fTodosService\x12D\n" +
	"\bGetTodos\x12\x16.google.protobuf.Empty\x1a .todo_proto.v1.ListTodosResponse\x12G\n" +
	"\aGetTodo\x12\x1e.todo_proto.v1.GetTodoIdParams\x1a\x1c.todo_proto.v1.TodosResponse\x12A\n" +
	"\n" +
	"CreateTodo\x12\x1b.todo_proto.v1.TodosRequest\x1a\x16.google.protobuf.EmptyBOZMgithub.com/raphacm/grpc-learning/grpc-protos/serviceB/gen/proto/v1;todo_protob\x06proto3"

var (
	file_v1_todos_proto_rawDescOnce sync.Once
	file_v1_todos_proto_rawDescData []byte
)

func file_v1_todos_proto_rawDescGZIP() []byte {
	file_v1_todos_proto_rawDescOnce.Do(func() {
		file_v1_todos_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_todos_proto_rawDesc), len(file_v1_todos_proto_rawDesc)))
	})
	return file_v1_todos_proto_rawDescData
}

var file_v1_todos_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_todos_proto_goTypes = []any{
	(*TodosRequest)(nil),      // 0: todo_proto.v1.TodosRequest
	(*TodosResponse)(nil),     // 1: todo_proto.v1.TodosResponse
	(*GetTodoIdParams)(nil),   // 2: todo_proto.v1.GetTodoIdParams
	(*ListTodosResponse)(nil), // 3: todo_proto.v1.ListTodosResponse
	(*emptypb.Empty)(nil),     // 4: google.protobuf.Empty
}
var file_v1_todos_proto_depIdxs = []int32{
	1, // 0: todo_proto.v1.ListTodosResponse.todos:type_name -> todo_proto.v1.TodosResponse
	4, // 1: todo_proto.v1.TodosService.GetTodos:input_type -> google.protobuf.Empty
	2, // 2: todo_proto.v1.TodosService.GetTodo:input_type -> todo_proto.v1.GetTodoIdParams
	0, // 3: todo_proto.v1.TodosService.CreateTodo:input_type -> todo_proto.v1.TodosRequest
	3, // 4: todo_proto.v1.TodosService.GetTodos:output_type -> todo_proto.v1.ListTodosResponse
	1, // 5: todo_proto.v1.TodosService.GetTodo:output_type -> todo_proto.v1.TodosResponse
	4, // 6: todo_proto.v1.TodosService.CreateTodo:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_todos_proto_init() }
func file_v1_todos_proto_init() {
	if File_v1_todos_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_todos_proto_rawDesc), len(file_v1_todos_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_todos_proto_goTypes,
		DependencyIndexes: file_v1_todos_proto_depIdxs,
		MessageInfos:      file_v1_todos_proto_msgTypes,
	}.Build()
	File_v1_todos_proto = out.File
	file_v1_todos_proto_goTypes = nil
	file_v1_todos_proto_depIdxs = nil
}
