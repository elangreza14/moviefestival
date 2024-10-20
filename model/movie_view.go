package model

import (
	"database/sql"
	"time"
)

type MovieView struct {
	MovieID int `db:"id"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (c MovieView) TableName() string {
	return "movie_views"
}

func NewMovieView(movieID int) *MovieView {
	return &MovieView{MovieID: movieID}
}

// to create in DB
func (c MovieView) Data() map[string]any {
	return map[string]any{
		"movie_id": c.MovieID,
	}
}

func (c MovieView) Columns() []string {
	return []string{"movie_id"}
}
