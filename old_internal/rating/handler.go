package rating

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

type Handler struct {
    service *Service
}

func NewHandler(service *Service) *Handler {
    return &Handler{service: service}
}

func (h *Handler) GetCategoryTop(c *gin.Context) {

    categoryName := c.Query("categoryName")
    if categoryName == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "categoryName is required"})
        return
    }

    log.Printf("[HTTP] Rebuilding top for category: %s", categoryName)

    if err := h.service.BuildTopByCategoryName(c, categoryName); err != nil {
        log.Printf("[HTTP] Error building category top: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "failed to build category top",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": "ok",
    })
}
