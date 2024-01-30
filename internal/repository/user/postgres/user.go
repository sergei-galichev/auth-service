package postgres

import (
	"auth-service/internal/repository/user/postgres/dao"
	"auth-service/pkg/tools/transaction"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (r *repository) CreateUser(saveUserDAO *dao.SaveUserDAO) (int64, error) {
	ctx := context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return -1, errors.New("user repo: begin transaction error")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, tx, err)
	}(ctx, tx)

	builder := r.genSQL.Insert("users").Values(saveUserDAO).Suffix("RETURNING id")
	query, args, err := builder.ToSql()
	if err != nil {
		return -1, errors.New("user repo: build sql error")
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return -1, errors.New("user repo: begin transaction error")
	}

	defer rows.Close()
	_ = rows

	return 0, nil
}
