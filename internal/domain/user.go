package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		ID          uuid.UUID        `db:"id"`
		Email       string           `db:"email"`
		Name        string           `db:"name"`
		Password    []byte           `db:"password"`
		Permission  int              `db:"permission"`
		Permissions []UserPermission `db:"-"`

		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
	}

	UserPermission struct {
		Val  int    `json:"val"`
		Name string `json:"name"`
	}
)

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

func (u *User) ValidPermission(reqPermission int) bool {
	return (u.Permission & reqPermission) > 0
}

var UserValPermission = UserPermission{
	Val:  1,
	Name: "user",
}

var AdminValPermission = UserPermission{
	Val:  2,
	Name: "admin",
}

var DefaultUserPermissions = []UserPermission{UserValPermission, AdminValPermission}

func (u *User) LoadPermissions() {
	if len(u.Permissions) > 0 {
		return
	}

	for _, permission := range DefaultUserPermissions {
		if u.ValidPermission(permission.Val) && permission.Val <= u.Permission {
			u.Permissions = append(u.Permissions, permission)
		}
	}
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
