package postgres

import (
	"auth-service/internal/repository/user/postgres/dao"
	"github.com/pkg/errors"
)

var (
	errUserNotFound = errors.New("user not found")
)

func (r *repository) CreateUser(user *dao.UserDAO) (int64, error) {
	res, err := r.session.Collection("users").Insert(user)
	if err != nil {
		return -1, errors.New("repo: [CreateUser.InsertInto] : " + err.Error())
	}

	return res.ID().(int64), nil
}

func (r *repository) IsUserExists(email string) (bool, error) {
	return r.session.Collection("users").Find("email = ?", email).Exists()
}

func (r *repository) GetUserByEmail(email string) (*dao.UserDAO, error) {
	var userDAO dao.UserDAO
	err := r.session.Collection("users").
		Find("email = ?", email).
		One(&userDAO)
	if err != nil {
		return nil, errUserNotFound
	}

	return &userDAO, nil
}
