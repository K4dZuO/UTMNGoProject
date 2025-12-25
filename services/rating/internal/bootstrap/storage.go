package bootstrap

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"rating_service/internal/storage/pgstorage"
)

func initStorage(ctx context.Context, dsn string) (*pgstorage.Storage, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return pgstorage.New(pool), nil
}
