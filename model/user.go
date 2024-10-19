package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uuid.UUID `db:"id"`
	Email      string    `db:"email"`
	Name       string    `db:"name"`
	Password   []byte    `db:"password"`
	Permission int       `db:"permission"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func NewUser(email, password, name string) (*User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id,
		Email:    email,
		Name:     name,
		Password: pass,
	}, nil
}

func (u *User) IsPasswordValid(reqPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(reqPassword))
	return err == nil
}

func (u User) TableName() string {
	return "users"
}

// to create in DB
func (u User) Data() map[string]any {
	return map[string]any{
		"id":       u.ID,
		"email":    u.Email,
		"name":     u.Name,
		"password": u.Password,
	}
}

func (u User) Columns() []string {
	return []string{
		"id",
		"email",
		"name",
		"password",
	}
}
