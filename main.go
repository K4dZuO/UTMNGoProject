package main

import (
    "fmt"
    "context"
    "log"
    "go_back/config"
    "go_back/internal/database"
    "go_back/internal/seeder"
    "go_back/internal/kafka"
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

    if err := seeder.SeedProducts(ctx, pool); err != nil {
        log.Fatal(err)
    }


    producer, err := kafka.NewSyncProducer([]string{"localhost:9092"})
    if err != nil {
        log.Fatal("producer init error:", err)
    }
    defer producer.Close()

    err = kafka.SendMessage(producer, "recalculation-requests", "hello from Go!")
    if err != nil {
        log.Fatal("send error:", err)
    }
    
}
