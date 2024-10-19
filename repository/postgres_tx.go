package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type (
	PgxTXer interface {
		Begin(context.Context) (pgx.Tx, error)
	}

	PostgresTransactionRepo struct {
		TXer PgxTXer
	}
)

func NewPostgresTransactionRepo(tx PgxTXer) *PostgresTransactionRepo {
	return &PostgresTransactionRepo{
		TXer: tx,
	}
}

func (pr PostgresTransactionRepo) WithTX(ctx context.Context, callback func(tx QueryPgx) error) error {
	tx, err := pr.TXer.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit(context.Background())
		default:
			_ = tx.Rollback(context.Background())
		}
	}()

	err = callback(tx)
	if err != nil {
		return err
	}

	return err
}
