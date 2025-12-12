package recommend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler { return &Handler{s: s} }

func (h *Handler) GetTop10(c *gin.Context) {
	category := c.Query("categoryName")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "categoryName required"})
		return
	}

	jsonStr, err := h.s.GetTop10(c, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// JSON уже подготовлен в redis как строка, возвращаем корректно
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(jsonStr))
}
