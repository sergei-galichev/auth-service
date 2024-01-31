package converter

import (
	"auth-service/internal/delivery/grpc_v1/dto"
	"auth-service/internal/repository/user/postgres/dao"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//func UserDomainToUserDAO(user *user.User) *dao.UserDAO {
//	return &dao.UserDAO{
//		UUID:      user.UUID,
//		Email:     user.Email,
//		PassHash:  []byte(""),
//		Role:      user.Role,
//		CreatedAt: time.Now(),
//	}
//}

func CreateDTOToUserDAO(createDTO *dto.UserCreateDTO) (*dao.UserDAO, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(createDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("user service: could not hash password")
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
		return nil, errors.New("user service: could not hash password")
	}

	return &dao.UserDAO{
		Email:    loginDTO.Email,
		PassHash: passHash,
	}, nil
}
