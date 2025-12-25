package bootstrap

import (
	"log"

	"recommend_service/internal/config"
)

type App struct {
	HTTP *HTTPServer
}

func Build(cfg *config.Config) (*App, error) {
	redisRepo, err := initRedis(cfg.Redis.Addr)
	if err != nil {
		return nil, err
	}

	ratingClient := initRatingClient(cfg.Rating.BaseURL)

	service := initService(redisRepo, ratingClient)

	httpServer := initHTTP(cfg.HTTP.Addr, service)

	log.Println("recommend service bootstrap completed")

	return &App{
		HTTP: httpServer,
	}, nil
}
