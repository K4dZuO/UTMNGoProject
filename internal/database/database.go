package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool() (*pgxpool.Pool, error) {
    pgcfg, err := pgxpool.ParseConfig("postgres://root:root@postgres:5432/marketdb")    
    if err != nil {
        return nil, fmt.Errorf("cannot parse config: %w", err)
    }

    pgcfg.MaxConns = 50
    pool, err := pgxpool.NewWithConfig(context.Background(), pgcfg)
    if err != nil {
        return nil, fmt.Errorf("cannot create pool: %w", err)
    }

    // Проверяем подключение
    if err := pool.Ping(context.Background()); err != nil {
        return nil, fmt.Errorf("cannot ping database: %w", err)
    }

    return pool, nil
}
