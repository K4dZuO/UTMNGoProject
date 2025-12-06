package database

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool() (*pgxpool.Pool, error) {
    dsn := "postgres://root:root@localhost:5432/marketdb"

    cfg, err := pgxpool.ParseConfig(dsn)
    if err != nil {
        return nil, fmt.Errorf("cannot parse config: %w", err)
    }

    pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
    if err != nil {
        return nil, fmt.Errorf("cannot create pool: %w", err)
    }

    // Проверяем подключение
    if err := pool.Ping(context.Background()); err != nil {
        return nil, fmt.Errorf("cannot ping database: %w", err)
    }

    return pool, nil
}
