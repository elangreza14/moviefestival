package model

import (
	"database/sql"
	"time"
)

type (
	Genre struct {
		Name string `db:"name"`

		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
	}

	ViewedGenre struct {
		Name  string
		Views int
	}
)

func (c Genre) TableName() string {
	return "genres"
}

func NewGenre(name string) *Genre {
	return &Genre{
		Name: name,
	}
}

// to create in DB
func (c Genre) Data() map[string]any {
	return map[string]any{
		"name": c.Name,
	}
}

func (c Genre) Columns() []string {
	return []string{
		"name",
	}
}
