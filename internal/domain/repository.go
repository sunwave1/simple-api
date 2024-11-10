package domain

import "context"

type Repository interface {
	FindById(ctx context.Context, id string) (interface{}, error)
	FindAll(ctx context.Context) ([]interface{}, error)
	Update(ctx context.Context, id string, object interface{}) error
	Delete(ctx context.Context, id string) error
	Add(ctx context.Context, object interface{}) error
}
