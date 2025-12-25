package ratingService

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"rating_service/internal/models"
)

type Storage interface {
	RecalculateProductRating(ctx context.Context, productID int) (float64, error)
	GetProductCategory(ctx context.Context, productID int) (int, error)
	GetCategoryName(ctx context.Context, categoryID int) (string, error)
	GetCategoryIDByName(ctx context.Context, name string) (int, error)
	GetTopProductsByCategory(ctx context.Context, categoryID int) ([]models.Product, error)
}

type Cache interface {
	Set(ctx context.Context, key string, value string) error
}

// Serivce
type Service struct {
	storage Storage
	cache   Cache
}

func New(storage Storage, cache Cache) *Service {
	return &Service{
		storage: storage,
		cache:   cache,
	}
}

/*
   === Kafka entrypoint ===
*/

// HandleReviewCreated — основной event-driven вход
func (s *Service) HandleReviewCreated(ctx context.Context, productID int) error {

	// 1. Пересчитать рейтинг товара
	_, err := s.storage.RecalculateProductRating(ctx, productID)
	if err != nil {
		return fmt.Errorf("recalculate rating failed: %w", err)
	}

	// 2. Найти категорию товара
	categoryID, err := s.storage.GetProductCategory(ctx, productID)
	if err != nil {
		return fmt.Errorf("get product category failed: %w", err)
	}

	// 3. Имя категории
	categoryName, err := s.storage.GetCategoryName(ctx, categoryID)
	if err != nil {
		return fmt.Errorf("get category name failed: %w", err)
	}

	// 4. Топ-10 категории
	topProducts, err := s.storage.GetTopProductsByCategory(ctx, categoryID)
	if err != nil {
		return fmt.Errorf("get top products failed: %w", err)
	}

	inTop := false
	for _, p := range topProducts {
		if p.ID == productID {
			inTop = true
			break
		}
	}

	if !inTop {
		return nil
	}

	// 6. Обновляем Redis
	return s.writeTopToCache(ctx, categoryName, topProducts)
}


// RebuildCategoryTop — служебный HTTP-метод
func (s *Service) RebuildCategoryTop(ctx context.Context, categoryName string) error {
	if categoryName == "" {
		return errors.New("category name is empty")
	}

	categoryID, err := s.storage.GetCategoryIDByName(ctx, categoryName)
	if err != nil {
		return err
	}

	topProducts, err := s.storage.GetTopProductsByCategory(ctx, categoryID)
	if err != nil {
		return err
	}

	return s.writeTopToCache(ctx, categoryName, topProducts)
}


type topProductItem struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}

func (s *Service) writeTopToCache(
	ctx context.Context,
	categoryName string,
	products []models.Product,
) error {

	result := make(map[string]topProductItem)

	for i, p := range products {
		key := fmt.Sprintf("%d", i+1)
		result[key] = topProductItem{
			ID:   p.ID,
			Name: p.Name,
			Rate: p.Rate,
		}
	}

	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return s.cache.Set(ctx, categoryName, string(data))
}
