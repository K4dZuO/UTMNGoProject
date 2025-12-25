package bootstrap

import (
	"recommend_service/internal/services/recommendService"
)

func initService(
	cache recommendService.Cache,
	rating recommendService.RatingClient,
) *recommendService.Service {
	return recommendService.New(cache, rating)
}
