package repository

import (
	"context"

	"github.com/elangreza14/moviefestival/model"
)

type genreRepository struct {
	db QueryPgx
	*PostgresRepo[model.Genre]
}

func NewGenreRepository(
	dbPool QueryPgx,
) *genreRepository {
	return &genreRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Genre](dbPool),
	}
}

func (ar *genreRepository) UpsertGenre(ctx context.Context, name string) error {
	sql := `INSERT INTO genres ("name") VALUES ($1) ON CONFLICT (name) DO NOTHING`
	_, err := ar.db.Exec(ctx, sql, name)
	if err != nil {
		return err
	}
	return err
}

func (ar *genreRepository) InsertGenreWithMovieMapping(ctx context.Context, name string, movieID int) error {
	sql := `INSERT INTO movie_genres ("movie_id", "genre_name") VALUES ($1, $2)
			ON CONFLICT (movie_id,genre_name) DO NOTHING`
	_, err := ar.db.Exec(ctx, sql, movieID, name)
	if err != nil {
		return err
	}
	return err
}
