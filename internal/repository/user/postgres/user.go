package postgres

import (
	"auth-service/internal/repository/user/postgres/dao"
)

func (r *repository) CreateUser(user *dao.UserDAO) (int64, error) {
	// TODO
	r.logger.Debug("CreateUser: implement me")
	return 0, nil
}
