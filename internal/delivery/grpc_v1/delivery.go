package grpc_v1

import (
	services "auth-service/internal/service"
	auth_v1 "auth-service/pkg/grpc/v1/auth"
	"auth-service/pkg/logging"
)

type AuthImplementation struct {
	auth_v1.UnimplementedAuthServiceV1Server

	userService services.UserService

	logger *logging.Logger
}

func NewImplementation(
	userService services.UserService,
) *AuthImplementation {
	return &AuthImplementation{
		userService: userService,
		logger:      logging.GetLogger(),
	}
}
