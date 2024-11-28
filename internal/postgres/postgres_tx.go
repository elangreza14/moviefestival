package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type (
	TxPgx interface {
		Begin(context.Context) (pgx.Tx, error)
	}

	PostgresTransactionRepo struct {
		TXer TxPgx
	}
)

func NewTXRepo(tx TxPgx) *PostgresTransactionRepo {
	return &PostgresTransactionRepo{
		TXer: tx,
	}
}

func (pr PostgresTransactionRepo) WithTX(ctx context.Context, callback func(tx QueryPgx) error) error {
	tx, err := pr.TXer.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit(ctx)
		default:
			err = tx.Rollback(ctx)
		}
	}()

	err = callback(tx)
	if err != nil {
		return err
	}

	return err
}
