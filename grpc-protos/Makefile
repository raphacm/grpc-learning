PROJECT_FOLDER=$(PROJECT)
PROTO_DIR=$(PROJECT_FOLDER)/proto
OUT_DIR=$(PROJECT_FOLDER)/gen/proto

init: 


generate: 
	$(if $(strip $(PROJECT)),,$(error PROJECT var is required. Usage: make generate PROJECT=<project_folder_name>))
	@echo "Generating proto files for project $(PROJECT) in folder: $(OUT_DIR)"
	mkdir -p $(OUT_DIR)
	protoc --proto_path=$(PROTO_DIR) \
		   --go_out=paths=source_relative:$(OUT_DIR) \
		   --go-grpc_out=paths=source_relative:$(OUT_DIR) \
		   $(PROTO_DIR)/v1/todos.proto
	cd $(PROJECT_FOLDER) && go mod tidy

