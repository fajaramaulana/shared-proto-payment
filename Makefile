.PHONY: all auth payment transaction notification user account clean check-tools generate-mocks

# Protoc and Go plugin versions
PROTOC_GEN_GO_VERSION := v1.28
PROTOC_GEN_GO_GRPC_VERSION := v1.2
PROTOC_GEN_GO := protoc-gen-go
PROTOC_GEN_GO_GRPC := protoc-gen-go-grpc
PROTO_PATH := proto
OUTPUT_DIR := .

# Define the command to generate Protobuf files
define generate_proto
	protoc --proto_path=$(PROTO_PATH) \
           --go_out=$(OUTPUT_DIR) --go_opt=module=github.com/fajaramaulana/shared-proto-payment \
           $(1)
    protoc --proto_path=$(PROTO_PATH) \
           --go_out=$(OUTPUT_DIR) --go_opt=module=github.com/fajaramaulana/shared-proto-payment \
           --go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=module=github.com/fajaramaulana/shared-proto-payment \
           $(2)
endef

# Check if necessary tools are installed
check-tools:
	@which protoc > /dev/null || (echo "Protoc is not installed" && exit 1)
	@which $(PROTOC_GEN_GO) > /dev/null || (echo "protoc-gen-go is not installed" && exit 1)
	@which $(PROTOC_GEN_GO_GRPC) > /dev/null || (echo "protoc-gen-go-grpc is not installed" && exit 1)
	@which mockery > /dev/null || (echo "mockery is not installed" && exit 1)

# Main target to run all tasks
all: clean check-tools auth payment transaction notification user account generate-mocks

# Generate Protobuf files for each service
auth:
	$(call generate_proto, $(PROTO_PATH)/auth/auth.proto, $(PROTO_PATH)/auth/auth_service.proto)

payment:
	$(call generate_proto, $(PROTO_PATH)/payment/payment.proto, $(PROTO_PATH)/payment/payment_service.proto)

transaction:
	$(call generate_proto, $(PROTO_PATH)/transaction/transaction.proto, $(PROTO_PATH)/transaction/transaction_service.proto)

notification:
	$(call generate_proto, $(PROTO_PATH)/notification/notification.proto, $(PROTO_PATH)/notification/notification_service.proto)

user:
	$(call generate_proto, $(PROTO_PATH)/user/user.proto, $(PROTO_PATH)/user/user_service.proto)

account:
	$(call generate_proto, $(PROTO_PATH)/account/account.proto, $(PROTO_PATH)/account/account_service.proto)

# Generate mocks for all interfaces
generate-mocks:
	@echo "Generating mocks for all interfaces..."
	mockery --all --output=$(OUTPUT_DIR)/mocks --outpkg=mocks --case=underscore

# Clean up generated files
clean:
	rm -f $(OUTPUT_DIR)/proto/auth/*.pb.go $(OUTPUT_DIR)/proto/payment/*.pb.go $(OUTPUT_DIR)/proto/transaction/*.pb.go $(OUTPUT_DIR)/proto/notification/*.pb.go $(OUTPUT_DIR)/proto/user/*.pb.go $(OUTPUT_DIR)/proto/account/*.pb.go
	rm -f $(OUTPUT_DIR)/mocks/*.go
