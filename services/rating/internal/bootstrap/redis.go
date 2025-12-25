package bootstrap

import (
	"context"

	redisclient "github.com/redis/go-redis/v9"

	"rating_service/internal/cache/redis"
)

func initRedis(addr string) (*redis.Repository, error) {
	client := redisclient.NewClient(&redisclient.Options{
		Addr: addr,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return redis.New(client), nil
}
