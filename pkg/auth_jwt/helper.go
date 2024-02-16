package auth_jwt

import (
	"auth-service/internal/config"
	"auth-service/internal/repository/user/postgres/dao"
	"auth-service/pkg/logging"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

type UserClaims struct {
	jwt.RegisteredClaims
}

type jwtHelper struct {
	logger *logging.Logger
	cfg    config.AuthConfig
}

func NewHelper(cfg config.AuthConfig) JWTHelper {
	logger := logging.GetLogger()
	return &jwtHelper{
		logger: logger,
		cfg:    cfg,
	}
}

func (h *jwtHelper) GenerateAccessToken(dao *dao.UserDAO) (string, error) {
	accessTTL := h.cfg.AccessTTL()

	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:     dao.UUID,
			Issuer: "auth-service",
			Audience: []string{
				"users",
			},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return h.generateToken(&claims)
}

func (h *jwtHelper) GenerateRefreshToken(dao *dao.UserDAO) (string, error) {
	refreshTTL := h.cfg.RefreshTTL()
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:     dao.UUID,
			Issuer: "auth-service",
			Audience: []string{
				"users",
			},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTTL)),
		},
	}

	return h.generateToken(&claims)
}

func (h *jwtHelper) generateToken(claims *UserClaims) (string, error) {
	key := []byte(h.cfg.Secret())

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	if err != nil {
		return "", errors.New("jwt: [GenerateAccessToken.SignedString]: " + err.Error())
	}

	return token, nil
}

func (h *jwtHelper) validateToken(token string) (bool, error) {
	key := []byte(h.cfg.Secret())

	t, err := jwt.Parse(
		token, func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, errors.New("jwt: signing method invalid: ")
			}
			return key, nil
		},
	)
	if err != nil {
		return false, errors.New("jwt: error parsing token")
	}
	if t.Valid {
		return true, nil
	}

	return false, errors.New("jwt: invalid token")
}

func (h *jwtHelper) ExchangeRefreshToken(accessToken, refreshToken string) (at, rt string, err error) {
	//key := []byte(h.cfg.Secret())
	//t, err := jwt.ParseWithClaims(
	//	accessToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
	//		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
	//			return nil, errors.New("jwt: [GetUUIDFromAccessToken.Alg]: " + err.Error())
	//		}
	//		return key, nil
	//	},
	//)
	//if err != nil {
	//	return "", errors.New("jwt: [GetUUIDFromAccessToken.ParseWithClaims]: " + err.Error())
	//}
	//claims := t.Claims.(*UserClaims)
	//return claims.uuid, nil

	return "", "", nil
}

func (h *jwtHelper) AccessTokenTTL() time.Duration {
	return h.cfg.AccessTTL()
}

func (h *jwtHelper) RefreshTokenTTL() time.Duration {
	return h.cfg.RefreshTTL()
}
