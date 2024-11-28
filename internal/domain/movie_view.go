package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type MovieViewHistory struct {
	MovieID int       `db:"id"`
	UserID  uuid.UUID `db:"id"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (c MovieViewHistory) TableName() string {
	return "movie_views_histories"
}

func NewMovieViewHistory(movieID int) *MovieViewHistory {
	return &MovieViewHistory{MovieID: movieID}
}

// to create in DB
func (c MovieViewHistory) Data() map[string]any {
	maps := map[string]any{
		"movie_id": c.MovieID,
	}

	if uuid.Nil != c.UserID {
		maps["user_id"] = c.UserID
	}

	return maps
}

func (c MovieViewHistory) Columns() []string {
	return []string{"movie_id"}
}
