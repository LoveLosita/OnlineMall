package model

import "time"

type ShowProduct struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description,omitempty"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
	CategoryID   int       `json:"category_id,omitempty"`
	ProductImage string    `json:"product_image,omitempty"`
	Popularity   int       `json:"popularity,omitempty"`
	AveRating    float64   `json:"ave_rating,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type AddProduct struct {
	Name         string  `json:"name"`
	Description  string  `json:"description,omitempty"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategoryID   int     `json:"category_id,omitempty"`
	ProductImage string  `json:"product_image,omitempty"`
}
