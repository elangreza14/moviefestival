package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgxutil"
)

type (
	QueryPgx interface {
		Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
		Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	}

	Entity interface {
		TableName() string
		Columns() []string
		Data() map[string]any
	}

	PostgresRepo[T Entity] struct {
		db               QueryPgx
		QueryBasicSelect string
		tableName        string
	}
)

func NewPostgresRepo[T Entity](dbPool QueryPgx) *PostgresRepo[T] {
	var table T

	columns := []string{"created_at", "updated_at"}
	columns = append(columns, table.Columns()...)

	tableName := table.TableName()

	return &PostgresRepo[T]{
		db: dbPool,
		QueryBasicSelect: fmt.Sprintf(`SELECT %s FROM %s `,
			strings.Join(columns, ","),
			tableName,
		),
		tableName: tableName,
	}
}

func (pr *PostgresRepo[T]) Get(ctx context.Context, by string, val any, columns ...string) (*T, error) {
	q := fmt.Sprintf(pr.QueryBasicSelect+` WHERE %s = $1 LIMIT 1`, by)
	if len(columns) > 0 {
		q = fmt.Sprintf(`select %s from %s  WHERE %s = $1 LIMIT 1`, strings.Join(columns, ", "), pr.tableName, by)
	}

	v, err := pgxutil.SelectRow(ctx, pr.db, q, []any{val}, pgx.RowToAddrOfStructByNameLax[T])
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (pr *PostgresRepo[T]) GetAll(ctx context.Context) ([]T, error) {
	v, err := pgxutil.Select(ctx, pr.db, pr.QueryBasicSelect, nil, pgx.RowToStructByNameLax[T])
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (pr *PostgresRepo[T]) Create(ctx context.Context, payloads ...T) error {
	for _, payload := range payloads {
		err := pgxutil.InsertRow(ctx, pr.db, pr.tableName, payload.Data())
		if err != nil {
			return err
		}
	}

	return nil
}

func (pr *PostgresRepo[T]) Edit(ctx context.Context, payload T, whereValues map[string]any) error {
	_, err := pgxutil.Update(ctx, pr.db, pr.tableName, payload.Data(), whereValues)
	return err
}
