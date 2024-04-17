package db

import (
	"context"
	"film-library/internal/config"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDb(c config.Config, ctx context.Context) (*Database, error) {
	pool, err := pgxpool.Connect(ctx, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Database.Host, c.Database.Port, c.Database.User, "password", c.Database.Dbname))
	//pool, err := pgxpool.Connect(ctx, fmt.Sprintf("host=localhost port=5433 user=postgres password=password dbname=postgres sslmode=disable"))
	if err != nil {
		return nil, err
	}
	return NewDatabase(pool), nil
}
