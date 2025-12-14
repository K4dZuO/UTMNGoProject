package recommend

import (
    "context"
    "github.com/redis/go-redis/v9"
)

type RedisRepository interface {
    Get(ctx context.Context, key string) (string, error)
    Set(ctx context.Context, key, val string) error
}

type redisRepo struct {
    client *redis.Client
}

func NewRedisRepository(c *redis.Client) RedisRepository {
    return &redisRepo{client: c}
}

func (r *redisRepo) Get(ctx context.Context, key string) (string, error) {
    return r.client.Get(ctx, key).Result()
}

func (r *redisRepo) Set(ctx context.Context, key, val string) error {
    return r.client.Set(ctx, key, val, 0).Err()
}
