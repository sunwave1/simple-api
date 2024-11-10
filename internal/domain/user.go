package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id              string     `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Password        string     `json:"password"`
	EmailVerified   bool       `json:"email_verified"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

func NewUser(name string, email string, password string) *User {
	return &User{
		Id:              uuid.NewString(),
		Name:            name,
		Email:           email,
		Password:        password,
		EmailVerified:   false,
		EmailVerifiedAt: nil,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		DeletedAt:       nil,
	}
}

func BuildUser(
	id string,
	name string,
	email string,
	password string,
	emailVerified bool,
	emailVerifiedAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *User {
	return &User{
		Id:              id,
		Name:            name,
		Email:           email,
		Password:        password,
		EmailVerified:   emailVerified,
		EmailVerifiedAt: emailVerifiedAt,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		DeletedAt:       deletedAt,
	}
}
