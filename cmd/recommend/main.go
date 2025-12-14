package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"go_back/config"
	"go_back/internal/recommend"
)

func main() {
	cfg, err := config.Load("recommend_config.yaml")
	if err != nil {
		log.Printf("warning: config load failed: %v — продолжим с env", err)
	}

	redisAddr := cfg.Redis.Addr

	ratingURL := "http://rating:8082"

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	svc := recommend.NewService(rdb, ratingURL)
	h := recommend.NewHandler(svc)

	router := gin.Default()
	router.GET("/get_top_category", h.GetTop10)

	port := getEnv("RECOMMEND_PORT", "8083")
	log.Printf("recommend service started on :%s (redis=%s, rating=%s)", port, redisAddr, ratingURL)
	router.Run(":" + port)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
