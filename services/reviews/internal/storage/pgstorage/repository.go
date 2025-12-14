package pgstorage

import (
	"context"
	"reviews_service/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
    db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Save(ctx context.Context, review models.ReviewInfo) error {
    id := uuid.New()

    query := `
		INSERT INTO reviews (id, product_id, rate)
		VALUES ($1, $2, $3)
	`
    _, err := s.db.Exec(
		ctx,
		query,
		id,
		review.ProductID,
		review.Rate,
	)
    return err
}
