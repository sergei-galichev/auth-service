package service

import (
	"auth-service/internal/delivery/grpc_v1/dto"
)

type UserService interface {
	CreateUser(dto *dto.UserCreateDTO) (int64, error)
	LoginUser(dto *dto.UserLoginDTO) (accessToken, refreshToken string, err error)
}
