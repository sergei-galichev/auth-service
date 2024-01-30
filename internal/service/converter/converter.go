package converter

import (
	"auth-service/internal/delivery/grpc_v1/dto"
	"auth-service/internal/repository/user/postgres/dao"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//func UserDomainToUserDAO(user *user.User) *dao.SaveUserDAO {
//	return &dao.SaveUserDAO{
//		UUID:      user.UUID,
//		Email:     user.Email,
//		PassHash:  []byte(""),
//		Role:      user.Role,
//		CreatedAt: time.Now(),
//	}
//}

func CreateDTOToUserDAO(createDTO *dto.UserCreateDTO) (*dao.SaveUserDAO, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(createDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("user service: could not hash password")
	}

	return &dao.SaveUserDAO{
		UUID:      uuid.New().String(),
		Email:     createDTO.Email,
		PassHash:  passHash,
		Role:      createDTO.Role,
		CreatedAt: time.Now(),
	}, nil
}
