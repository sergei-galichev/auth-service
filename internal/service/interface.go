package service

import (
	"auth-service/internal/domain/user"
)

type UserService interface {
	CreateUser(user *user.User) (int64, error)
	CheckEmail(email string) bool
}
