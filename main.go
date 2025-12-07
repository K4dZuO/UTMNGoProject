package main

import (
    "fmt"
    "context"
    "log"
    "go_back/config"
    "go_back/internal/database"
    "go_back/internal/seeder"
)

func main() {
    appCfg, err := config.Load("config.yaml")
    ctx := context.Background()

    if err != nil {
        log.Fatalf("config load error: %v", err)
    }

    pool, err := database.NewPostgresPool()
    if err != nil {
        log.Fatalf("DB error: %v", err)
    }
    defer pool.Close()

    fmt.Println("PostgreSQL connected!")

    // запускаем миграции
    if err := database.RunMigrations(appCfg.Postgres.DSN, appCfg.Migrations.Path); err != nil {
    log.Fatal(err)
    }

    /// Генерируем данные в таблицы
    conn, err := pool.Acquire(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Release()

    if err := seeder.SeedCategories(ctx, conn.Conn()); err != nil {
        log.Fatal(err)
    }
}
