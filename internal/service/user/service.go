package user

import (
	"auth-service/internal/repository"
	services "auth-service/internal/service"
	"auth-service/pkg/auth_jwt"
	"auth-service/pkg/cache"
)

var _ services.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
	userCache      cache.Repository
	jwtHelper      auth_jwt.JWTHelper
}

func NewService(
	userRepository repository.UserRepository,
	userCache cache.Repository,
) *service {
	return &service{
		userRepository: userRepository,
		userCache:      userCache,
		jwtHelper:      auth_jwt.NewHelper(),
	}
}
