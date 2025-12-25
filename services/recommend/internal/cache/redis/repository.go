package redis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

var ErrNotFound = errors.New("category top not found")

type Repository struct {
	client *redis.Client
}

func New(client *redis.Client) *Repository {
	return &Repository{client: client}
}

// Get возвращает JSON-строку топа по названию категории
func (r *Repository) Get(ctx context.Context, category string) (string, error) {
	val, err := r.client.Get(ctx, category).Result()
	if err == redis.Nil {
		return "", ErrNotFound
	}
	if err != nil {
		return "", err
	}
	return val, nil
}
