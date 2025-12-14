package rating

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "go_back/internal/models"

    "github.com/redis/go-redis/v9"
)

type Service struct {
    repo  Repository
    redis *redis.Client
}

func NewService(repo Repository, redisClient *redis.Client) *Service {
    return &Service{
        repo:  repo,
        redis: redisClient,
    }
}

// Структура для записи в Redis
type TopProductItem struct {
    ID    int     `json:"id"`
    Rate  float64 `json:"rate"`
    Name  string  `json:"name"`
}

// ПЕРЕСЧЁТ РЕЙТИНГА
func (s *Service) RecalculateAndUpdate(ctx context.Context, productID int) error {

    // Пересчёт рейтинга
    _, err := s.repo.RecalculateProductRating(ctx, productID)
    if err != nil {
        return fmt.Errorf("failed to recalc rating: %w", err)
    }

    // Получить категорию товара
    categoryID, err := s.repo.GetProductCategory(ctx, productID)
    if err != nil {
        return fmt.Errorf("failed to get product category: %w", err)
    }

    // Получить имя категории
    categoryName, err := s.repo.GetCategoryName(ctx, categoryID)
    if err != nil {
        return fmt.Errorf("failed to get category name: %w", err)
    }

    // Получить топ-10 товаров категории
    topProducts, err := s.repo.GetTopProductsByCategory(ctx, categoryID)
    if err != nil {
        return fmt.Errorf("failed to get top products: %w", err)
    }

    //  входит ли товар в топ-10
    inTop := false
    for _, p := range topProducts {
        if p.ID == productID {
            inTop = true
            break
        }
    }

    if !inTop {
        return nil // ничего обновлять в Redis не надо
    }

    return s.writeTopToRedis(ctx, categoryName, topProducts)
}

// ОБРАБОТЧИК ДЛЯ HTTP
func (s *Service) BuildTopByCategoryName(ctx context.Context, categoryName string) error {

    // получить id категории	
    categoryID, err := s.findCategoryID(ctx, categoryName)
    if err != nil {
        return err
    }

    topProducts, err := s.repo.GetTopProductsByCategory(ctx, categoryID)
    if err != nil {
        return err
    }

    return s.writeTopToRedis(ctx, categoryName, topProducts)
}

// получить id категории по имени
func (s *Service) findCategoryID(ctx context.Context, name string) (int, error) {

    query := `
        SELECT id
        FROM categories
        WHERE name = $1
        LIMIT 1;
    `

    var id int
    err := s.repo.(*PgRepository).db.QueryRow(ctx, query, name).Scan(&id)
    if err != nil {
        return 0, errors.New("category not found")
    }
    return id, nil
}


func (s *Service) writeTopToRedis(ctx context.Context, categoryName string, products []models.Product) error {

    topMap := make(map[string]TopProductItem)

    for i, p := range products {
        key := fmt.Sprintf("%d", i+1)

        topMap[key] = TopProductItem{
            ID:   p.ID,
            Rate: p.Rate,
            Name: p.Name,
        }
    }

    jsonData, err := json.Marshal(topMap)
    if err != nil {
        return err
    }

    return s.redis.Set(ctx, categoryName, string(jsonData), 0).Err()
}
