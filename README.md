# Shared Proto

This README provides step-by-step instructions to set up the shared proto for the microservice-payment-go-grpc project.

## Prerequisites

Ensure you have the following installed on your system:
- Go (https://golang.org/doc/install)
- Make (https://www.gnu.org/software/make/)

## Installation Steps

### 1. Install Protocol Buffers

Download and install Protocol Buffers from the official site:
- [Protocol Buffers](https://developers.google.com/protocol-buffers)

### 2. Install Go Protobuf Plugins

Install the necessary Go plugins for Protocol Buffers:
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 3. Install Make

Make sure `make` is installed on your system. You can install it using the package manager for your OS.

For Debian/Ubuntu:
```sh
sudo apt-get install build-essential
```

For macOS:
```sh
brew install make
```

### 4. Run Make

Navigate to the project directory and run:
```sh
make all
```

This will compile the protobuf files and generate the necessary Go code.

## Conclusion

You have successfully set up the shared proto for the microservice-payment-go-grpc project. You can now proceed with the development.
