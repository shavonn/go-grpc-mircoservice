package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/shavonn/go-grpc-microservice/internal/config"
	"github.com/shavonn/go-grpc-microservice/internal/server"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Setup logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	// Create gRPC server
	grpcServer := server.NewGRPCServer(cfg)

	// Start server
	lis, err := net.Listen("tcp", cfg.Server.Address)
	if err != nil {
		log.Fatal().Err(err).Str("address", cfg.Server.Address).Msg("Failed to listen")
	}

	log.Info().Str("address", cfg.Server.Address).Msg("Starting gRPC server")

	// Graceful shutdown
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("Failed to serve gRPC server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server...")
	grpcServer.GracefulStop()
	log.Info().Msg("Server stopped")
}
