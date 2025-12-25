package seeder

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SeedProducts(ctx context.Context, pool *pgxpool.Pool) error {
	var count int
	if err := pool.QueryRow(ctx, "SELECT COUNT(*) FROM products").Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		fmt.Println("Products already seeded, skipping")
		return nil
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var maxCategoryID int
	if err := tx.QueryRow(ctx, "SELECT COUNT(*) FROM categories").Scan(&maxCategoryID); err != nil {
		return err
	}

	batch := &pgx.Batch{}

	for i := 1; i <= 100_000; i++ {
		categoryID := 1 + rand.Intn(maxCategoryID)
		name := fmt.Sprintf("Item #%d", i)

		batch.Queue(
			"INSERT INTO products(name, category_id) VALUES ($1, $2)",
			name, categoryID,
		)

		if i%1000 == 0 {
			br := tx.SendBatch(ctx, batch)
			if err := br.Close(); err != nil {
				return err
			}
			batch = &pgx.Batch{}
		}
	}

	if batch.Len() > 0 {
		br := tx.SendBatch(ctx, batch)
		if err := br.Close(); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
