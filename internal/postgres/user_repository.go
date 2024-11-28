package postgres

import (
	"github.com/elangreza14/moviefestival/internal/domain"
)

type userRepository struct {
	db QueryPgx
	*PostgresRepo[domain.User]
}

func NewUserRepository(
	dbPool QueryPgx,
) *userRepository {
	return &userRepository{
		db:           dbPool,
		PostgresRepo: NewRepo[domain.User](dbPool),
	}
}
