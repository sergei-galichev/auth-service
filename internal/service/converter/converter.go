package converter

import (
	"auth-service/internal/delivery/grpc_v1/dto"
	"auth-service/internal/repository/user/postgres/dao"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	errGeneratePassHash = errors.New("user service: could not hash password")
)

func ConvertCreateDTOToUserDAO(createDTO *dto.UserCreateDTO) (*dao.UserDAO, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(createDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errGeneratePassHash
	}

	return &dao.UserDAO{
		UUID:      uuid.New().String(),
		Email:     createDTO.Email,
		PassHash:  passHash,
		Role:      createDTO.Role,
		CreatedAt: time.Now(),
	}, nil
}

func LoginDTOToUserDAO(loginDTO *dto.UserLoginDTO) (*dao.UserDAO, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(loginDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errGeneratePassHash
	}

	return &dao.UserDAO{
		Email:    loginDTO.Email,
		PassHash: passHash,
	}, nil
}
