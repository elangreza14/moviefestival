package repository

import (
	"github.com/elangreza14/moviefestival/model"
)

type movieViewHistoryRepository struct {
	db QueryPgx
	*PostgresRepo[model.MovieViewHistory]
}

func NewMovieViewHistoryRepository(
	dbPool QueryPgx,
) *movieViewHistoryRepository {
	return &movieViewHistoryRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.MovieViewHistory](dbPool),
	}
}
