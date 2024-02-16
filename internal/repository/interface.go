package repository

import (
	"auth-service/internal/repository/user/postgres/dao"
)

type UserRepository interface {
	CreateUser(user *dao.UserDAO) (int64, error)
	IsUserExists(email string) (bool, error)
	GetUserByEmail(email string) (*dao.UserDAO, error)
}
