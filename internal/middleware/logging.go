package middleware

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func LoggingUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		resp, err := handler(ctx, req)

		duration := time.Since(start)

		logger := log.With().
			Str("method", info.FullMethod).
			Dur("duration", duration).
			Logger()

		if err != nil {
			logger.Error().Err(err).Msg("gRPC request failed")
		} else {
			logger.Info().Msg("gRPC request completed")
		}

		return resp, err
	}
}

func LoggingStreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()

		err := handler(srv, stream)

		duration := time.Since(start)

		logger := log.With().
			Str("method", info.FullMethod).
			Dur("duration", duration).
			Logger()

		if err != nil {
			logger.Error().Err(err).Msg("gRPC stream failed")
		} else {
			logger.Info().Msg("gRPC stream completed")
		}

		return err
	}
}
