PROTO_DIR=internal/grpc/proto
OUT_DIR=internal/grpc/gen

generate:
	protoc --proto_path=$(PROTO_DIR) --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) $(PROTO_DIR)/v1/todos.proto
