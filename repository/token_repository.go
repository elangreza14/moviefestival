package repository

import (
	"github.com/elangreza14/moviefestival/model"
)

type tokenRepository struct {
	db QueryPgx
	*PostgresRepo[model.Token]
}

func NewTokenRepository(
	dbPool QueryPgx,
) *tokenRepository {
	return &tokenRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Token](dbPool),
	}
}
