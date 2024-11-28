package postgres

import (
	"context"

	"github.com/elangreza14/moviefestival/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type movieViewRepository struct {
	db QueryPgx
	*PostgresRepo[domain.MovieView]
	txRepo *PostgresTransactionRepo
}

func NewMovieViewRepository(
	dbPool *pgxpool.Pool,
	querier QueryPgx,
) *movieViewRepository {
	if dbPool != nil {
		return &movieViewRepository{
			db:           dbPool,
			PostgresRepo: NewRepo[domain.MovieView](dbPool),
			txRepo:       NewTXRepo(dbPool),
		}
	}

	if querier == nil {
		return nil
	}

	return &movieViewRepository{
		db:           querier,
		PostgresRepo: NewRepo[domain.MovieView](querier),
	}
}

func (ur *movieViewRepository) AddMovieViewTX(ctx context.Context, movieID int) error {
	return ur.txRepo.WithTX(ctx, func(tx QueryPgx) error {
		movieViewHistoryRepo := NewMovieViewHistoryRepository(tx)
		err := movieViewHistoryRepo.Create(ctx, *domain.NewMovieViewHistory(movieID))
		if err != nil {
			return err
		}

		movieRepo := NewMovieViewRepository(nil, tx)
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
