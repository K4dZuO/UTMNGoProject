package rating

import (
	"context"
	"go_back/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
    GetProductCategory(ctx context.Context, productID int) (int, error)
    GetCategoryName(ctx context.Context, categoryID int) (string, error)
    RecalculateProductRating(ctx context.Context, productID int) (float64, error)
    GetTopProductsByCategory(ctx context.Context, categoryID int) ([]models.Product, error)
}

type PgRepository struct {
    db *pgxpool.Pool
}

func NewPgRepository(db *pgxpool.Pool) *PgRepository {
    return &PgRepository{db: db}
}

// Получить category_id по product_id
func (r *PgRepository) GetProductCategory(ctx context.Context, productID int) (int, error) {
    var categoryID int

    query := `
        SELECT category_id
        FROM products
        WHERE id = $1;
    `

    err := r.db.QueryRow(ctx, query, productID).Scan(&categoryID)
    return categoryID, err
}

// Получить имя категории
func (r *PgRepository) GetCategoryName(ctx context.Context, categoryID int) (string, error) {
    var name string

    query := `
        SELECT name
        FROM categories
        WHERE id = $1;
    `

    err := r.db.QueryRow(ctx, query, categoryID).Scan(&name)
    return name, err
}

// Пересчитать рейтинг: SUM(rate) / COUNT(rate)
func (r *PgRepository) RecalculateProductRating(ctx context.Context, productID int) (float64, error) {
    var sum float64
    var count int

    query := `
        SELECT COALESCE(SUM(rate), 0), COUNT(*)
        FROM reviews
        WHERE product_id = $1;
    `

    err := r.db.QueryRow(ctx, query, productID).Scan(&sum, &count)
    if err != nil {
        return 0, err
    }
    if count == 0 {
        return 0, nil // товара ещё нет в рейтингах
    }

    newRate := sum / float64(count)

    // обновляем products.rate
    _, err = r.db.Exec(ctx, `
        UPDATE products
        SET rate = $1
        WHERE id = $2;
    `, newRate, productID )

    return newRate, err
}

// Получить топ товаров категории
func (r *PgRepository) GetTopProductsByCategory(ctx context.Context, categoryID int) ([]models.Product, error) {
    query := `
        SELECT id, name, rate, category_id
        FROM products
        WHERE category_id = $1
        ORDER BY rate DESC
        LIMIT 10;
    `

    rows, err := r.db.Query(ctx, query, categoryID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []models.Product

    for rows.Next() {
        var p models.Product
        if err := rows.Scan(&p.ID, &p.Name, &p.Rate, &p.CategoryID); err != nil {
            return nil, err
        }
        products = append(products, p)
    }

    return products, nil
}
