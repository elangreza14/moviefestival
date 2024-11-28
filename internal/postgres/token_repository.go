package postgres

import (
	"github.com/elangreza14/moviefestival/internal/domain"
)

type tokenRepository struct {
	db QueryPgx
	*PostgresRepo[domain.Token]
}

func NewTokenRepository(
	dbPool QueryPgx,
) *tokenRepository {
	return &tokenRepository{
		db:           dbPool,
		PostgresRepo: NewRepo[domain.Token](dbPool),
	}
}
