package repository

import (
	"context"

	"github.com/elangreza14/moviefestival/model"
)

type movieViewRepository struct {
	db QueryPgx
	*PostgresRepo[model.Movie]
	txRepo *PostgresTransactionRepo
}

func NewMovieViewRepository(
	dbPool QueryPgx,
	tx PgxTXer,
) *movieViewRepository {
	return &movieViewRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Movie](dbPool),
		txRepo:       NewPostgresTransactionRepo(tx),
	}
}

func (ur *movieViewRepository) AddMovieViewTX(ctx context.Context, movieID int) error {
	return ur.txRepo.WithTX(ctx, func(tx QueryPgx) error {
		movieRepo := NewMovieViewRepository(tx, nil)
		return movieRepo.UpsertMovieView(ctx, movieID)
	})
}

func (ar *movieViewRepository) UpsertMovieView(ctx context.Context, movieID int) error {
	sql := `INSERT INTO movie_views
		(movie_id, "views")
		VALUES ($1, 1) 
		ON CONFLICT (movie_id)
		DO UPDATE 
		SET views=movie_views.views+1`
	_, err := ar.db.Exec(ctx, sql, movieID)
	if err != nil {
		return err
	}
	return err
}
