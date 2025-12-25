package bootstrap

import (
	"recommend_service/internal/client/rating"
)

func initRatingClient(baseURL string) *rating.Client {
	return rating.New(baseURL)
}
