package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RatingService interface {
	RebuildCategoryTop(ctx context.Context, categoryName string) error
}

type Handler struct {
	service RatingService
}

func New(service RatingService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetCategoryTop(c *gin.Context) {
	var req RebuildCategoryTopRequest

	if err := c.ShouldBindQuery(&req); err != nil || req.CategoryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "categoryName is required",
		})
		return
	}

	if err := h.service.RebuildCategoryTop(c.Request.Context(), req.CategoryName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, RebuildCategoryTopResponse{
		Status: "ok",
	})
}
