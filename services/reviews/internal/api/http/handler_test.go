package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)


type mockReviewService struct {
	createFn func(ctx context.Context, productID int, rate int) (string, error)
}

func (m *mockReviewService) CreateReview(
	ctx context.Context,
	productID int,
	rate int,
) (string, error) {
	return m.createFn(ctx, productID, rate)
}


func TestCreateReview_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &mockReviewService{
		createFn: func(ctx context.Context, productID int, rate int) (string, error) {
			require.Equal(t, 1, productID)
			require.Equal(t, 5, rate)
			return "test-review-id", nil
		},
	}

	handler := New(mockSvc)

	r := gin.New()
	r.POST("/reviews", handler.CreateReview)

	body := `{"product_id":1,"rate":5}`
	req := httptest.NewRequest(
		http.MethodPost,
		"/reviews",
		bytes.NewBufferString(body),
	)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	require.Contains(t, w.Body.String(), "test-review-id")
}
