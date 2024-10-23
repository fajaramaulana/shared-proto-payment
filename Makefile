.PHONY: all auth payment transaction notification user account

PROTOC_GEN_GO := protoc-gen-go
PROTOC_GEN_GO_GRPC := protoc-gen-go-grpc
PROTO_PATH := proto

all: auth payment transaction notification

auth:
	protoc --proto_path=$(PROTO_PATH) \
	       --go_out=. --go_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       --go-grpc_out=. --go-grpc_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       $(PROTO_PATH)/auth/*.proto

payment:
	protoc --proto_path=$(PROTO_PATH) \
	       --go_out=. --go_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       --go-grpc_out=. --go-grpc_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       $(PROTO_PATH)/payment/*.proto

transaction:
	protoc --proto_path=$(PROTO_PATH) \
	       --go_out=. --go_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       --go-grpc_out=. --go-grpc_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       $(PROTO_PATH)/transaction/*.proto

notification:
	protoc --proto_path=$(PROTO_PATH) \
	       --go_out=. --go_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       --go-grpc_out=. --go-grpc_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       $(PROTO_PATH)/notification/*.proto

user:
	protoc --proto_path=$(PROTO_PATH) \
	       --go_out=. --go_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       --go-grpc_out=. --go-grpc_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       $(PROTO_PATH)/user/*.proto

account:
	protoc --proto_path=$(PROTO_PATH) \
	       --go_out=. --go_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       --go-grpc_out=. --go-grpc_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       $(PROTO_PATH)/account/*.proto
