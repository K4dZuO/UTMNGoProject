package models

type TopProduct struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}

// CategoryTop — топ товаров по категории
type CategoryTop map[string]TopProduct
