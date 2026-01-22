package main_db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseUrl string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
