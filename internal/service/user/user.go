package user

import (
	"auth-service/internal/domain/user"
	"auth-service/internal/service/converter"
)

func (s *service) CreateUser(user *user.User) (int64, error) {
	return s.userRepository.CreateUser(converter.UserDomainToUserDAO(user))
}

func (s *service) CheckEmail(email string) bool {
	return user.IsValidEmail(email)
}
