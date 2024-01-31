package postgres

import (
	"auth-service/internal/repository/user/postgres/dao"
	"github.com/pkg/errors"
)

func (r *repository) CreateUser(user *dao.UserDAO) (int64, error) {
	res, err := r.session.Collection("users").Insert(user)
	//_, err := r.session.SQL().InsertInto("users").Values(user).Exec()
	if err != nil {
		return -1, errors.New("repo: [CreateUser.InsertInto] : " + err.Error())
	}

	return res.ID().(int64), nil
}

func (r *repository) GetUser(user *dao.UserDAO) (*dao.UserDAO, error) {
	var userDAO dao.UserDAO
	ok, _ := r.session.Collection("users").
		Find("email = ?", user.Email).
		And("password = ?", user.PassHash).
		Exists()
	if !ok {
		return nil, errors.New("repo: [GetUser.Find] : invalid credentials")
	}
	//err = r.session.SQL().Select("*").From("users").Where("email = ?", email).One(&userDAO)
	//if err != nil {
	//	return nil, errors.New("repo: [GetUser.Select] : " + err.Error())
	//}
	return &userDAO, nil
}
