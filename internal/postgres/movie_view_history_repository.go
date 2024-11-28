package postgres

import (
	"github.com/elangreza14/moviefestival/internal/domain"
)

type movieViewHistoryRepository struct {
	db QueryPgx
	*PostgresRepo[domain.MovieViewHistory]
}

func NewMovieViewHistoryRepository(
	dbPool QueryPgx,
) *movieViewHistoryRepository {
	return &movieViewHistoryRepository{
		db:           dbPool,
		PostgresRepo: NewRepo[domain.MovieViewHistory](dbPool),
	}
}
