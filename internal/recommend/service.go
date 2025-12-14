package recommend

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	redis      *redis.Client
	ratingBase string
}

func NewService(rdb *redis.Client, ratingBaseURL string) *Service {
	return &Service{
		redis:      rdb,
		ratingBase: ratingBaseURL,
	}
}

// get_top_category: проверяет redis, если нет — вызывает rating endpoint и снова читает
func (s *Service) GetTop10(ctx context.Context, category string) (string, error) {
	// 1) попытка получить из Redis
	val, err := s.redis.Get(ctx, category).Result()
	if err == nil && val != "" {
		return val, nil
	}

	// если key не найден, Redis возвращает Nil error
	if err != nil && err != redis.Nil {
		// реальная ошибка соединения
		return "", fmt.Errorf("redis get error: %w", err)
	}

	// 2) вызов сервиса rating для обновления топа
	if err := s.callRatingUpdate(category); err != nil {
		return "", fmt.Errorf("call to rating failed: %w", err)
	}

	time.Sleep(time.Second)

	val, err = s.redis.Get(ctx, category).Result()
	if err == redis.Nil || val == "" {
		return "", errors.New("no data in redis after rating update")
	}
	if err != nil {
		return "", fmt.Errorf("redis get after update error: %w", err)
	}

	return val, nil
}
