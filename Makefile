.PHONY: all auth payment transaction notification user account clean check-tools

PROTOC_GEN_GO_VERSION := v1.28
PROTOC_GEN_GO_GRPC_VERSION := v1.2
PROTOC_GEN_GO := protoc-gen-go
PROTOC_GEN_GO_GRPC := protoc-gen-go-grpc
PROTO_PATH := proto
OUTPUT_DIR := .

define generate_proto
	protoc --proto_path=$(PROTO_PATH) \
	       --go_out=$(OUTPUT_DIR) --go_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       --go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=module=github.com/fajaramaulana/shared-proto-payment \
	       $(1)
endef

check-tools:
	@which protoc > /dev/null || (echo "Protoc is not installed" && exit 1)
	@which $(PROTOC_GEN_GO) > /dev/null || (echo "protoc-gen-go is not installed" && exit 1)
	@which $(PROTOC_GEN_GO_GRPC) > /dev/null || (echo "protoc-gen-go-grpc is not installed" && exit 1)

all: check-tools auth payment transaction notification

auth:
	$(call generate_proto, $(PROTO_PATH)/auth/*.proto)

payment:
	$(call generate_proto, $(PROTO_PATH)/payment/*.proto)

transaction:
	$(call generate_proto, $(PROTO_PATH)/transaction/*.proto)

notification:
	$(call generate_proto, $(PROTO_PATH)/notification/*.proto)

user:
	$(call generate_proto, $(PROTO_PATH)/user/*.proto)

account:
	$(call generate_proto, $(PROTO_PATH)/account/*.proto)

clean:
	rm -f auth/*.pb.go payment/*.pb.go transaction/*.pb.go notification/*.pb.go user/*.pb.go account/*.pb.go
