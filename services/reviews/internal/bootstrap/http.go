package bootstrap

import (
	"github.com/gin-gonic/gin"

	httpapi "reviews_service/internal/api/http"
)

func initHTTP(handler *httpapi.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/reviews", handler.CreateReview)

	return router
}

