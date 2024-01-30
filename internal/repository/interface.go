package repository

import (
	"auth-service/internal/repository/user/postgres/dao"
)

type UserRepository interface {
	CreateUser(user *dao.SaveUserDAO) (int64, error)
}
