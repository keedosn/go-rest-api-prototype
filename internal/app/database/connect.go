package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect(connStr string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
