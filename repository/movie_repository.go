package repository

import (
	"github.com/elangreza14/moviefestival/model"
)

type movieRepository struct {
	db QueryPgx
	*PostgresRepo[model.Movie]
}

func NewMovieRepository(
	dbPool QueryPgx,
) *movieRepository {
	return &movieRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Movie](dbPool),
	}
}
