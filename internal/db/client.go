package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDb(ctx context.Context) (*Database, error) {
	pool, err := pgxpool.Connect(ctx, fmt.Sprintf("host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"))
	if err != nil {
		return nil, err
	}
	return NewDatabase(pool), nil
}
