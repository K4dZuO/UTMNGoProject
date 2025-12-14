package seeder

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SeedProducts(ctx context.Context, pool *pgxpool.Pool) error {
	// Создаем транзакцию для ускорения
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Узнаем количество категорий
	var maxCategoryID int
	err = tx.QueryRow(ctx, "SELECT COUNT(*) FROM categories").Scan(&maxCategoryID)
	if err != nil {
		return fmt.Errorf("failed to get categories count: %w", err)
	}

	batch := &pgx.Batch{}
	
	for i := 1; i <= 100_000; i++ {
		categoryId := 1 + rand.Intn(maxCategoryID)
		name := fmt.Sprintf("Item #%d", i)

		batch.Queue(
			"INSERT INTO products(name, category_id) VALUES ($1, $2)",
			name,
			categoryId,
		)

		// Отправляем batch каждые 1000 записей
		if i%1000 == 0 {
			br := tx.SendBatch(ctx, batch)
			if err := br.Close(); err != nil {
				return fmt.Errorf("batch insert failed: %w", err)
			}
			batch = &pgx.Batch{} // Создаем новый batch
			
			// Логирование прогресса
			if i%10000 == 0 {
				fmt.Printf("Inserted %d products...\n", i)
			}
		}
	}

	// Отправляем оставшиеся записи
	if batch.Len() > 0 {
		br := tx.SendBatch(ctx, batch)
		if err := br.Close(); err != nil {
			return fmt.Errorf("final batch insert failed: %w", err)
		}
	}

	// Коммитим транзакцию
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Printf("Successfully inserted 1,000,000 products\n")
	return nil
}
