.PHONY: build run test clean generate docker-build docker-run

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=grpc-microservice
BINARY_UNIX=$(BINARY_NAME)_unix

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/server

# Run the application
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/server
	./$(BINARY_NAME)

# Test the application
test:
	$(GOTEST) -v ./...

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Generate protobuf code
generate:
	@chmod +x scripts/generate.sh
	@./scripts/generate.sh

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./cmd/server

# Docker
docker-build:
	docker build -t $(BINARY_NAME):latest .

docker-run:
	docker run -p 8080:8080 $(BINARY_NAME):latest

# Development
dev:
	air -c .air.toml

deps:
	$(GOGET) -d ./...
	$(GOGET) -u ./...
