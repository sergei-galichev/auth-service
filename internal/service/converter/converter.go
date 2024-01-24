package converter

import (
	"auth-service/internal/domain/user"
	"auth-service/internal/repository/user/postgres/dao"
	"time"
)

func UserDomainToUserDAO(user *user.User) *dao.UserDAO {
	return &dao.UserDAO{
		UUID:      user.UUID,
		Email:     user.Email,
		PassHash:  []byte(""),
		Role:      user.Role,
		CreatedAt: time.Now(),
	}
}
