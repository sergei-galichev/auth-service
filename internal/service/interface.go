package service

import (
	"auth-service/internal/delivery/grpc_v1/dto"
)

type UserService interface {
	CreateUser(user *dto.UserCreateDTO) (int64, error)
}
