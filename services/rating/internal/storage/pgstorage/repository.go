package pgstorage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"

	"rating_service/internal/models"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{db: db}
}

func (s *Storage) RecalculateProductRating(
	ctx context.Context,
	productID int,
) (float64, error) {

	var sum float64
	var count int

	err := s.db.QueryRow(
		ctx,
		`
		SELECT COALESCE(SUM(rate), 0), COUNT(*)
		FROM reviews
		WHERE product_id = $1;
		`,
		productID,
	).Scan(&sum, &count)
	if err != nil {
		return 0, err
	}

	if count == 0 {
		// нет отзывов — рейтинг 0
		_, err = s.db.Exec(
			ctx,
			`UPDATE products SET rate = 0 WHERE id = $1`,
			productID,
		)
		return 0, err
	}

	newRate := sum / float64(count)

	_, err = s.db.Exec(
		ctx,
		`
		UPDATE products
		SET rate = $1
		WHERE id = $2;
		`,
		newRate,
		productID,
	)
	if err != nil {
		return 0, err
	}

	return newRate, nil
}

// GetProductCategory возвращает category_id товара
func (s *Storage) GetProductCategory(
	ctx context.Context,
	productID int,
) (int, error) {

	var categoryID int

	err := s.db.QueryRow(
		ctx,
		`
		SELECT category_id
		FROM products
		WHERE id = $1;
		`,
		productID,
	).Scan(&categoryID)

	return categoryID, err
}

// GetCategoryName возвращает имя категории по id
func (s *Storage) GetCategoryName(
	ctx context.Context,
	categoryID int,
) (string, error) {

	var name string

	err := s.db.QueryRow(
		ctx,
		`
		SELECT name
		FROM categories
		WHERE id = $1;
		`,
		categoryID,
	).Scan(&name)

	return name, err
}

// GetCategoryIDByName возвращает id категории по имени
func (s *Storage) GetCategoryIDByName(
	ctx context.Context,
	name string,
) (int, error) {

	var id int

	err := s.db.QueryRow(
		ctx,
		`
		SELECT id
		FROM categories
		WHERE name = $1
		LIMIT 1;
		`,
		name,
	).Scan(&id)

	if err != nil {
		return 0, errors.New("category not found")
	}

	return id, nil
}

// GetTopProductsByCategory возвращает топ-10 товаров категории
func (s *Storage) GetTopProductsByCategory(
	ctx context.Context,
	categoryID int,
) ([]models.Product, error) {

	rows, err := s.db.Query(
		ctx,
		`
		SELECT id, name, rate, category_id
		FROM products
		WHERE category_id = $1
		ORDER BY rate DESC
		LIMIT 10;
		`,
		categoryID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Rate,
			&p.CategoryID,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
