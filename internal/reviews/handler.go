package reviews

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Handler struct {
    svc *Service
}

func NewHandler(svc *Service) *Handler {
    return &Handler{svc: svc}
}

// @Summary Create review
// @Description Create new review and trigger rating recalculation
// @Tags reviews
// @Accept json
// @Produce json
// @Param review body CreateReviewRequest true "Review data"
// @Success 201 {object} CreateReviewResponse
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /reviews [post]
func (h *Handler) CreateReview(c *gin.Context) {
    var req CreateReviewRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
        return
    }

    if req.Rate < 1 || req.Rate > 5 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "rating must be 1-5"})
        return
    }

    id, err := h.svc.CreateReview(c, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, CreateReviewResponse{ReviewID: id})
}
