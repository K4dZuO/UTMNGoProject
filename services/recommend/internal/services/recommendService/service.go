package recommendService

import (
	"context"
	"errors"
	"fmt"
)

type Cache interface {
	Get(ctx context.Context, category string) (string, error)
}

type RatingClient interface {
	RebuildCategoryTop(category string) error
}

type Service struct {
	cache  Cache
	rating RatingClient
}

func New(cache Cache, rating RatingClient) *Service {
	return &Service{
		cache:  cache,
		rating: rating,
	}
}

func (s *Service) GetTopByCategory(ctx context.Context, category string) (string, error) {
	if category == "" {
		return "", errors.New("category is required")
	}

	val, err := s.cache.Get(ctx, category)
	if err == nil {
		return val, nil
	}

	if err := s.rating.RebuildCategoryTop(category); err != nil {
		return "", fmt.Errorf("failed to rebuild category top: %w", err)
	}

	val, err = s.cache.Get(ctx, category)
	if err != nil {
		return "", errors.New("category top not found after rebuild")
	}

	return val, nil
}
