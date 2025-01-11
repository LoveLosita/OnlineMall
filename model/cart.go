package model

type AddToCart struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ProductInCart struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
