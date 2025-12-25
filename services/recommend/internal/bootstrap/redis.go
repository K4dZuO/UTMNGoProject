package bootstrap

import (
	"context"

	redis2 "github.com/redis/go-redis/v9"

	"recommend_service/internal/cache/redis"
)

func initRedis(addr string) (*redis.Repository, error) {
	client := redis2.NewClient(&redis2.Options{
		Addr: addr,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return redis.New(client), nil
}
