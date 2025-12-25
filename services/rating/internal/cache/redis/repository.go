package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	client *redis.Client
}

func New(client *redis.Client) *Repository {
	return &Repository{client: client}
}

// Set сохраняет строковое значение по ключу без TTL
func (r *Repository) Set(ctx context.Context, key string, value string) error {
	return r.client.Set(ctx, key, value, 0).Err()
}
