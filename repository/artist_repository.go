package repository

import (
	"context"

	"github.com/elangreza14/moviefestival/model"
)

type artistRepository struct {
	db QueryPgx
	*PostgresRepo[model.Artist]
}

func NewArtistRepository(
	dbPool QueryPgx,
) *artistRepository {
	return &artistRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Artist](dbPool),
	}
}

func (ar *artistRepository) UpsertArtist(ctx context.Context, name string) error {
	sql := `INSERT INTO artists ("name") VALUES ($1) ON CONFLICT (name) DO NOTHING`
	_, err := ar.db.Exec(ctx, sql, name)
	if err != nil {
		return err
	}
	return err
}

func (ar *artistRepository) InsertArtistWithMovieMapping(ctx context.Context, name string, movieID int) error {
	sql := `INSERT INTO movie_artists ("movie_id", "artist_name") VALUES ($1, $2)
			ON CONFLICT (movie_id,artist_name) DO NOTHING`
	_, err := ar.db.Exec(ctx, sql, movieID, name)
	if err != nil {
		return err
	}
	return err
}
