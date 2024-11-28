package domain

import (
	"database/sql"
	"time"
)

type Artist struct {
	Name string `db:"name"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (c Artist) TableName() string {
	return "artists"
}

func NewArtist(name string) *Artist {
	return &Artist{
		Name: name,
	}
}

// to create in DB
func (c Artist) Data() map[string]any {
	return map[string]any{
		"name": c.Name,
	}
}

func (c Artist) Columns() []string {
	return []string{
		"name",
	}
}
