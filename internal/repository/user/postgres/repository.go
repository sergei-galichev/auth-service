package postgres

import (
	repositories "auth-service/internal/repository"
	"auth-service/pkg/logging"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ repositories.UserRepository = (*repository)(nil)

type repository struct {
	db     *pgxpool.Pool
	genSQL squirrel.StatementBuilderType
	logger *logging.Logger
}

func New(db *pgxpool.Pool) *repository {
	return &repository{
		db: db,
		genSQL: squirrel.StatementBuilder.PlaceholderFormat(
			squirrel.Dollar,
		),
		logger: logging.GetLogger(),
	}
}
