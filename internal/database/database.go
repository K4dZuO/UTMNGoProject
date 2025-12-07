package database

import (
	"context"
	"fmt"
	"go_back/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool() (*pgxpool.Pool, error) {
    appCfg, err := config.Load("config.yaml")
    if err != nil {
        return nil, fmt.Errorf("cannot parse yaml: %w", err)
    }

    pgcfg, err := pgxpool.ParseConfig(appCfg.Postgres.DSN)
    if err != nil {
        return nil, fmt.Errorf("cannot parse config: %w", err)
    }

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
