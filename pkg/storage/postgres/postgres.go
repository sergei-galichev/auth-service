package postgres

import (
	"auth-service/pkg/logging"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

var (
	errorDatabaseConnect = errors.New("postgres: database connection error")
	errorDatabaseSession = errors.New("postgres: get session error")
	errorDatabasePing    = errors.New("postgres: database ping error")
	errorDatabaseClose   = errors.New("postgres: database close connection error")
)

type Storage struct {
	Session db.Session
	logger  *logging.Logger
}

func New(dsn string) (*Storage, error) {
	logger := logging.GetLogger()

	connDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, errorDatabaseConnect
	}

	session, err := postgresql.New(connDB)
	if err != nil {
		return nil, errorDatabaseSession
	}

	err = session.Ping()
	if err != nil {
		return nil, errorDatabasePing
	}

	logger.Info("Postgres is connected")

	return &Storage{
		Session: session,
		logger:  logger,
	}, nil
}

func (s *Storage) Close() {
	err := s.Session.Close()
	if err != nil {
		s.logger.Fatal(errorDatabaseClose)
	}
}
