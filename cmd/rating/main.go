package main

import (
    "context"
    "fmt"
    "log"

    "go_back/config"
    "go_back/internal/database"
    "go_back/internal/rating"

    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"

    "github.com/redis/go-redis/v9"
)

func main() {
    cfg, err := config.Load("rate_config.yaml")
    if err != nil {
        log.Fatalf("config load error: %v", err)
    }

    // === PostgreSQL ===
    pool, err := database.NewPostgresPool()
    if err != nil {
        log.Fatalf("DB error: %v", err)
    }
    defer pool.Close()

    fmt.Println("PostgreSQL connected!")

    // === Redis ===
    rdb := redis.NewClient(&redis.Options{
        Addr: cfg.Redis.Addr,
    })
    if err := rdb.Ping(context.Background()).Err(); err != nil {
        log.Fatalf("Redis error: %v", err)
    }
    fmt.Println("Redis connected!")

    repo := rating.NewPgRepository(pool)

    service := rating.NewService(repo, rdb)

    go rating.RunKafkaConsumer(
        cfg.Kafka.Brokers,
        cfg.Kafka.GroupID,
        cfg.Kafka.Topic,
        service,
    )

    // === HTTP SERVER WITH GIN ===
    handler := rating.NewHandler(service)

    r := gin.Default()

    r.GET("/get_category_top", handler.GetCategoryTop)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    log.Printf("Rating service started on %s", cfg.HTTP.Addr)
    r.Run(cfg.HTTP.Addr)
}
