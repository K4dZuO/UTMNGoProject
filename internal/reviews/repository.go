package reviews

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
	"github.com/google/uuid"
)

type Repository interface {
    InsertReview(ctx context.Context, req CreateReviewRequest) (string, error)
}

type PgRepository struct {
    db *pgxpool.Pool
}

func NewPgRepository(db *pgxpool.Pool) *PgRepository {
    return &PgRepository{db: db}
}

func (r *PgRepository) InsertReview(ctx context.Context, req CreateReviewRequest) (string, error) {
    id := uuid.New().String()

    query := `
        INSERT INTO reviews (id, product_id, rate)
        VALUES ($1, $2, $3)
        RETURNING id;
    `

    // Для RETURNING можно либо использовать ту же переменную, либо отдельную.
    var returnedID string

    err := r.db.QueryRow(ctx, query, id, req.ProductID, req.Rate).Scan(&returnedID)
    return returnedID, err
}
