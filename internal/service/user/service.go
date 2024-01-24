package user

import (
	"auth-service/internal/repository"
	services "auth-service/internal/service"
)

var _ services.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}
