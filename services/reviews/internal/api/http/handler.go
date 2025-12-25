package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Handler struct {
	svc ReviewService
}

type ReviewService interface {
	CreateReview(ctx context.Context, productID int, rate int) (string, error)
}


func New(svc ReviewService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateReview(c *gin.Context) {
    var req CreateReviewRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
        return
    }

    id, err := h.svc.CreateReview(c.Request.Context(), req.ProductID, req.Rate)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, CreateReviewResponse{ReviewID: id})
}
