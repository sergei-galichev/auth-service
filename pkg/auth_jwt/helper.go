package auth_jwt

import (
	"auth-service/internal/config"
	"auth-service/internal/config/env"
	"auth-service/internal/repository/user/postgres/dao"
	"auth-service/pkg/logging"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Email string
}

type jwtHelper struct {
	logger *logging.Logger
	Cfg    config.AuthConfig
}

func NewHelper() JWTHelper {
	logger := logging.GetLogger()
	cfg, err := env.NewAuthConfig()
	if err != nil {
		logger.Fatal("JWT config error: ", err)
	}
	return &jwtHelper{
		logger: logger,
		Cfg:    cfg,
	}
}

func (h *jwtHelper) GenerateAccessToken(dao *dao.UserDAO) (string, error) {
	h.logger.Debug("Create access token")

	key := []byte(h.Cfg.Secret())
	accessTTL := h.Cfg.AccessTTL()

	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID: strconv.FormatInt(dao.ID, 10),
			Audience: []string{
				"users",
			},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTTL)),
		},
		Email: dao.Email,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(key)
	if err != nil {
		return "", errors.New("jwt: [GenerateAccessToken.SignedString]: " + err.Error())
	}

	return token, nil
}
