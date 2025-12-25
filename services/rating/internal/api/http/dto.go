package http

type RebuildCategoryTopRequest struct {
	CategoryName string `form:"categoryName"`
}

type RebuildCategoryTopResponse struct {
	Status string `json:"status"`
}
