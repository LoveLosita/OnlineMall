package model

import "time"

type Product struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description,omitempty"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
	CategoryID   int       `json:"category_id,omitempty"`
	ProductImage string    `json:"product_image,omitempty"`
	Popularity   int       `json:"popularity,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
