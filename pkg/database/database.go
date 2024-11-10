package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
)

var (
	databaseOnce sync.Once
	database     *pgxpool.Pool
)

func GetDatabase() *pgxpool.Pool {
	databaseOnce.Do(func() {
		pool, err := pgxpool.Connect(context.Background(), "postgresql://postgres:1212@localhost:5432/postgres?pool_max_conns=10")
		if err != nil {
			return
		}
		database = pool
	})
	return database
}
