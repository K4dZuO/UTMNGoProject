package models

import "time"

type Category struct {
    ID      int     `json:"id" db:"id"`
    Name    string  `json:"name" db:"name"`
}

type Product struct {
    ID       int     `json:"id" db:"id"`
    Name     string  `json:"name" db:"name"`
    Rate     float64 `json:"rate" db:"rate"`
    CategoryID int  `json:"category_id" db:"category_id"`
}

type Review struct {
    ID        string    `json:"id" db:"id"` // uuid храним
    Rate      int       `json:"rate" db:"rate"`
    ProductID int       `json:"product_id" db:"product_id"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}
