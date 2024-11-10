package domain

import "time"

type Entity interface {
	GetId() string
	GetDeletedAt() time.Time
	GetUpdateAt() time.Time
	GetCreatedAt() time.Time
}
