package user

import (
	"auth-service/internal/delivery/grpc_v1/dto"
	"auth-service/internal/service/converter"
	"errors"
	"os"
)

func (s *service) CreateUser(createDTO *dto.UserCreateDTO) (int64, error) {
	if !dto.IsValidEmail(createDTO.Email) {
		return -1, errors.New("service: not valid email format")
	}
	if !dto.ArePasswordsMatch(createDTO.Password, createDTO.ConfirmPassword) {
		return -1, errors.New("service: password and confirm password does not match")
	}

	adminPass := os.Getenv("ADMIN_KEY")

	if dto.IsAdminRole(createDTO.Role) && !dto.IsAdminPasswordValid(adminPass, createDTO.AdminKey) {
		return -1, errors.New("service: admin password not valid")
	}

	d, err := converter.CreateDTOToUserDAO(createDTO)
	if err != nil {
		return -1, err
	}

	return s.userRepository.CreateUser(d)
}

func (s *service) LoginUser(dto *dto.UserLoginDTO) (accessToken, refreshToken string, err error) {
	d, err := converter.LoginDTOToUserDAO(dto)
	if err != nil {
		return "", "", err
	}

	d, err = s.userRepository.GetUser(d)
	if err != nil {
		return "", "", err
	}

	accessToken, err = s.jwtHelper.GenerateAccessToken(d)
	if err != nil {
		return "", "", err
	}

	return accessToken, d.UUID, nil
}
