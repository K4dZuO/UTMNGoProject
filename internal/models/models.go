package models

import "time"

// Product represents a product row from `products` table.
type Product struct {
    ID         string  `json:"id" db:"id"`             // base58 id, stored as TEXT in Postgres
    Name       string  `json:"name" db:"name"`
    Rate       float64 `json:"rate" db:"rate"`         // aggregated rating (double precision)
    CategoryID int     `json:"category_id" db:"category_id"`
}

type Category struct {
	ID int
	name string
}

type Review struct {
    ID        string    `json:"id" db:"id"`               
    Rate      int       `json:"rate" db:"rate"`           
    ProductID string    `json:"product_id" db:"product_id"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}
