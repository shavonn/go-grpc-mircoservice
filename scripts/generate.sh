#!/bin/bash

# Generate protobuf code
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/proto/user/v1/user.proto

# Move generated files to pkg/pb
mkdir -p pkg/pb
mv api/proto/user/v1/*.pb.go pkg/pb/

echo "Protobuf code generated successfully"
