package auth_jwt

import (
	"auth-service/internal/repository/user/postgres/dao"
)

type JWTHelper interface {
	GenerateAccessToken(dao *dao.UserDAO) (string, error)
}
