package models

import "time"

type Product struct {
    ID       int     `json:"id" db:"id"`
    Name     string  `json:"name" db:"name"`
    Rate     float64 `json:"rate" db:"rate"`
    Category string  `json:"category" db:"category"`
}

type Review struct {
    ID        string    `json:"id" db:"id"` // uuid храним
    Rate      int       `json:"rate" db:"rate"`
    ProductID int       `json:"product_id" db:"product_id"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}
