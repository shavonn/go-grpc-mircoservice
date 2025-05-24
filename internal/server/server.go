package server

import (
	"github.com/shavonn/go-grpc-microservice/internal/config"
	"github.com/shavonn/go-grpc-microservice/internal/middleware"
	"github.com/shavonn/go-grpc-microservice/internal/service"
	"github.com/shavonn/go-grpc-microservice/pkg/pb"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(cfg *config.Config) *grpc.Server {
	// Recovery options
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(middleware.RecoveryHandler),
	}

	// Create gRPC server with middleware
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.LoggingUnaryInterceptor(),
			recovery.UnaryServerInterceptor(recoveryOpts...),
		),
		grpc.ChainStreamInterceptor(
			middleware.LoggingStreamInterceptor(),
			recovery.StreamServerInterceptor(recoveryOpts...),
		),
	)

	// Register services
	userService := service.NewUserService()
	pb.RegisterUserServiceServer(server, userService)

	// Enable reflection for development
	reflection.Register(server)

	return server
}
