#!/bin/bash

# Build script for the gRPC microservice

set -e

echo "Building gRPC microservice..."

# Generate protobuf code
./scripts/generate.sh

# Build the binary
go build -o bin/grpc-microservice cmd/server/main.go

echo "Build completed successfully"
