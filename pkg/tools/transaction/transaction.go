package transaction

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func Finish(ctx context.Context, tx pgx.Tx, err error) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return errors.New("rollback error")
		}
		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			return errors.New("commit error")
		}
		return nil
	}
}
