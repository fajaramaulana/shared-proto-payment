.PHONY: all auth payment transaction notification

PROTOC_GEN_GO := protoc-gen-go
PROTOC_GEN_GO_GRPC := protoc-gen-go-grpc
PROTO_PATH := proto

all: auth payment transaction notification

auth:
	protoc --proto_path=$(PROTO_PATH) \
	       --go_out=. --go_opt=module=github.com/your-username/proto-repo \
	       --go-grpc_out=. --go-grpc_opt=module=github.com/your-username/proto-repo \
	       $(PROTO_PATH)/auth/*.proto

payment:
	protoc --proto_path=$(PROTO_PATH) \
	       --go_out=. --go_opt=module=github.com/your-username/proto-repo \
	       --go-grpc_out=. --go-grpc_opt=module=github.com/your-username/proto-repo \
	       $(PROTO_PATH)/payment/*.proto

transaction:
	protoc --proto_path=$(PROTO_PATH) \
	       --go_out=. --go_opt=module=github.com/your-username/proto-repo \
	       --go-grpc_out=. --go-grpc_opt=module=github.com/your-username/proto-repo \
	       $(PROTO_PATH)/transaction/*.proto

notification:
	protoc --proto_path=$(PROTO_PATH) \
	       --go_out=. --go_opt=module=github.com/your-username/proto-repo \
	       --go-grpc_out=. --go-grpc_opt=module=github.com/your-username/proto-repo \
	       $(PROTO_PATH)/notification/*.proto
