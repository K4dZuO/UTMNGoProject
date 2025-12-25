package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"reviews_service/tools/seed/internal/seeder"
)

func main() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN is required")
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	log.Println("Seeding categories...")
	if err := seeder.SeedCategories(ctx, pool); err != nil {
		log.Fatal(err)
	}

	log.Println("Seeding products...")
	if err := seeder.SeedProducts(ctx, pool); err != nil {
		log.Fatal(err)
	}

	log.Println("Seeding completed successfully")
}
