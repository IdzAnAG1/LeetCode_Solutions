package db

import (
	"LeetCode_Solutions/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect(dbc config.DatabaseConfiguration) error {
	cfg, err := pgxpool.ParseConfig(dbc.URL())
	if err != nil {
		return fmt.Errorf("Config parse error: %v\n", err)
	}
	pool, err := pgxpool.New(context.Background(), cfg.ConnString())
	if err != nil {
		return fmt.Errorf("Unable to connect: %v\n", err)
	}
	Pool = pool
	return nil
}

func Close() {
	if Pool != nil {
		Pool.Close()
	}
}
