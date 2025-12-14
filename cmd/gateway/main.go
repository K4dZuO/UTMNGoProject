package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go_back/docs" // swagger docs
)

// @title Market API Gateway
// @version 1.0
// @description Entry point for all Market microservices
// @BasePath /

// ----------------------- MAIN -----------------------

func main() {
	r := gin.Default()

	// Reviews service
	r.POST("/reviews", CreateReviewForward)

	// Top-category service
	r.GET("/get_top_category", GetTopCategoryForward)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// ----------------------- HANDLERS -----------------------

// CreateReviewForward godoc
// @Summary Create a new review
// @Description Forwards review creation to Reviews microservice
// @Accept json
// @Produce json
// @Param review body ReviewRequest true "Review data"
// @Success 201 {object} ReviewResponse
// @Failure 500 {object} ErrorResponse
// @Router /reviews [post]
func CreateReviewForward(c *gin.Context) {
	resp, err := http.Post("http://reviews:8081/reviews", "application/json", c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	c.Data(resp.StatusCode, "application/json", readAll(resp.Body))
}

// GetTopCategoryForward godoc
// @Summary Get top-10 products by category
// @Description Forwards request to Recommend microservice
// @Produce json
// @Param categoryName query string true "Category name"
// @Success 200 {object} TopResponse
// @Failure 500 {object} ErrorResponse
// @Router /get_top_category [get]
func GetTopCategoryForward(c *gin.Context) {
	url := "http://recommend:8083/get_top_category?" + c.Request.URL.RawQuery

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	c.Data(resp.StatusCode, "application/json", readAll(resp.Body))
}

// ----------------------- HELPERS -----------------------

func readAll(r io.Reader) []byte {
	b, _ := io.ReadAll(r)
	return b
}

// ----------------------- DTOs -----------------------

type ReviewRequest struct {
	ProductID int `json:"product_id"`
	Rate      int `json:"rate"`
}

type ReviewResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type TopResponse map[string]struct {
	ID   int     `json:"id"`
	Rate float64 `json:"rate"`
	Name string  `json:"name"`
}
