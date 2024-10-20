package model

import (
	"database/sql"
	"time"
)

type Movie struct {
	ID          int           `db:"id"`
	Title       string        `db:"title"`
	Description string        `db:"description"`
	WatchUrl    string        `db:"watch_url"`
	Duration    time.Duration `db:"duration"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (c Movie) TableName() string {
	return "movies"
}

func NewMovie(title, description, watchUrl string, duration time.Duration) *Movie {
	return &Movie{
		Title:       title,
		Description: description,
		WatchUrl:    watchUrl,
		Duration:    duration,
	}
}

// to create in DB
func (c Movie) Data() map[string]any {
	return map[string]any{
		"title":       c.Title,
		"description": c.Description,
		"watch_url":   c.WatchUrl,
		"duration":    c.Duration,
	}
}

func (c Movie) Columns() []string {
	return []string{
		"id",
		"title",
		"description",
		"watch_url",
		"duration",
	}
}
