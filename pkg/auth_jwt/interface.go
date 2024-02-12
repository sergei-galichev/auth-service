package auth_jwt

import (
	"auth-service/internal/repository/user/postgres/dao"
	"time"
)

type JWTHelper interface {
	GenerateAccessToken(dao *dao.UserDAO) (string, error)
	GenerateRefreshToken(dao *dao.UserDAO) (string, error)
	ExchangeRefreshToken(accessToken, refreshToken string) (at, rt string, err error)
	AccessTokenTTL() time.Duration
	RefreshTokenTTL() time.Duration
}
