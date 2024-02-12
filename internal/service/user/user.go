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

func (s *service) LoginUser(loginDTO *dto.UserLoginDTO) (accessToken, refreshToken string, err error) {
	d, err := s.userRepository.GetUserByEmail(loginDTO.Email)
	if err != nil {
		return "", "", err
	}

	if !dto.ComparePassAndHash(loginDTO.Password, d.PassHash) {
		return "", "", errors.New("service: invalid credentials")
	}

	accessToken, err = s.jwtHelper.GenerateAccessToken(d)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = s.jwtHelper.GenerateRefreshToken(d)
	if err != nil {
		return "", "", err
	}

	err = s.userCache.Set([]byte(refreshToken), []byte(accessToken), s.jwtHelper.RefreshTokenTTL())
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *service) LogoutUser(logoutDTO *dto.UserLogoutDTO) error {
	return nil
}
