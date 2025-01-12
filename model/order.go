package model

type Order struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Status    string `json:"status"`
	OrderDate string `json:"order_date"`
}

type PlaceOrder struct {
	UserID     int               `json:"user_id"`     //后端传入
	Items      []AProductInOrder `json:"items"`       //json传入
	TotalPrice float64           `json:"total_price"` //后端计算，items中每个商品的单价*数量的总和
	Status     string            `json:"status"`      //后端传入
	Address    string            `json:"address"`     //json传入
}

type AProductInOrder struct {
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"` //物品单价*数量
}

type ReturnOrder struct {
	ID        int    `json:"id"`
	Status    string `json:"status"`
	OrderDate string `json:"order_date"`
}
