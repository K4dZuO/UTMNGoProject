package reviews


type CreateReviewRequest struct {
    ProductID int       `json:"product_id" example:"123"`
    Rate      int       `json:"rate" example:"5"`
}

type CreateReviewResponse struct {
    ReviewID string `json:"id"`
}
