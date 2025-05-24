package middleware

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryHandler(p interface{}) error {
	log.Error().Interface("panic", p).Msg("gRPC panic recovered")
	return status.Errorf(codes.Internal, "internal server error")
}
