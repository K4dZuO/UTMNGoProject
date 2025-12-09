package main

import (
    "fmt"
    "context"
    "log"
    // "go_back/config"
    "go_back/internal/database"
    "go_back/internal/seeder"
    "go_back/internal/kafka"
    "go_back/internal/reviews"
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

func main() {
    // path := "config.yaml"

    // appCfg, err := config.Load(path)
    ctx := context.Background()

    // if err != nil {
        // log.Fatalf("config load error: %v", err)
    // }

    pool, err := database.NewPostgresPool()
    if err != nil {
        log.Fatalf("DB error: %v", err)
    }
    defer pool.Close()

    fmt.Println("PostgreSQL connected!")

    // запускаем миграции
    // if err := database.RunMigrations(appCfg.Postgres.DSN, appCfg.Migrations.Path); err != nil {
    if err := database.RunMigrations("postgres://root:root@postgres:5432/marketdb", "migrations"); err != nil {
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


    producer, err := kafka.NewSyncProducer([]string{"kafka:9092"})
    if err != nil {
        log.Fatal("producer init error:", err)
    }
    defer producer.Close()

    repo := reviews.NewPgRepository(pool)
    svc := reviews.NewService(repo, producer)
    handler := reviews.NewHandler(svc)

    r := gin.Default()

    r.POST("/reviews", handler.CreateReview)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.Run(":8081")
}
