package domain

import "context"

type UserRepository interface {
	FindById(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, name string) (*User, error)
	GetAll(ctx context.Context) ([]User, error)
	ExistById(ctx context.Context, id string) (bool, error)
	ExistByEmail(ctx context.Context, email string) (bool, error)
	Add(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
}
