package postgres

import (
	repositories "auth-service/internal/repository"
	"auth-service/pkg/logging"
	"github.com/upper/db/v4"
)

var (
	_      repositories.UserRepository = (*repository)(nil)
	logger                             = logging.GetLogger()
)

type repository struct {
	session db.Session
}

func New(session db.Session) *repository {
	return &repository{
		session: session,
	}
}
