package bootstrap

import (
	"github.com/gin-gonic/gin"

	httpapi "rating_service/internal/api/http"
	"rating_service/internal/services/ratingService"
)

func initHTTP(service *ratingService.Service) *gin.Engine {
	router := gin.Default()

	handler := httpapi.New(service)

	router.GET("/get_category_top", handler.GetCategoryTop)

	return router
}
