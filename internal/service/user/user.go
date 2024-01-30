package user

import (
	"auth-service/internal/delivery/grpc_v1/dto"
	"auth-service/internal/service/converter"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) CreateUser(createDTO *dto.UserCreateDTO) (int64, error) {
	if !dto.IsValidEmail(createDTO.Email) {
		return -1, status.Error(codes.InvalidArgument, "invalid email")
	}
	if !dto.CheckPasswordsMatch(createDTO.Password, createDTO.ConfirmPassword) {
		return -1, status.Error(codes.InvalidArgument, "passwords do not match")
	}

	d, err := converter.CreateDTOToUserDAO(createDTO)
	if err != nil {
		return -1, err
	}

	return s.userRepository.CreateUser(d)
}
