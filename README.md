# gRPC Microservice

## Quick Start

1. **Generate protobuf code:**
   ```bash
   make generate
   ```

2. **Run locally:**
   ```bash
   make run
   ```

3. **Build Docker image:**
   ```bash
   make docker-build
   ```

4. **Run with Docker Compose:**
   ```bash
   docker-compose up
   ```

## Development

### Prerequisites

- Go 1.24+
- Protocol Buffers compiler (protoc)
- protoc-gen-go and protoc-gen-go-grpc plugins

### Install protoc plugins:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Testing

```bash
make test
```

### API Documentation

The service exposes the following gRPC methods:

- `GetUser(GetUserRequest) returns (GetUserResponse)`
- `CreateUser(CreateUserRequest) returns (CreateUserResponse)`
- `UpdateUser(UpdateUserRequest) returns (UpdateUserResponse)`
- `DeleteUser(DeleteUserRequest) returns (DeleteUserResponse)`

## Configuration

Config is environment variables and .env. `internal/config/config.go`

## Deployment

Kubernetes manifests: `deployments/k8s/`.