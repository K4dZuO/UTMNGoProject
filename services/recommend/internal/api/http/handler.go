package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecommendService interface {
	GetTopByCategory(ctx context.Context, category string) (string, error)
}

type Handler struct {
	service RecommendService
}

func New(service RecommendService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetTop(c *gin.Context) {
	category := c.Query("categoryName")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "categoryName is required",
		})
		return
	}

	result, err := h.service.GetTopByCategory(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Data(
		http.StatusOK,
		"application/json; charset=utf-8",
		[]byte(result),
	)
}
