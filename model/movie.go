package model

import (
	"database/sql"
	"time"
)

type Movie struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	Manufacturer string `db:"manufacturer"`
	Price        int    `db:"price"`
	Image        string `db:"image"`
	Stock        int    `db:"stock"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (c Movie) TableName() string {
	return "movies"
}

// to create in DB
func (c Movie) Data() map[string]any {
	return map[string]any{
		"id":           c.ID,
		"name":         c.Name,
		"manufacturer": c.Manufacturer,
		"price":        c.Price,
		"image":        c.Image,
		"stock":        c.Stock,
	}
}

func (c Movie) Columns() []string {
	return []string{
		"id",
		"name",
		"manufacturer",
		"price",
		"image",
		"stock",
	}
}
