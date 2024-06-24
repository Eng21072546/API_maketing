package entity

import "time"

type Transaction struct {
	ID           string         `json:"id"`
	Address      string         `json:"address"`
	AccountName  string         `json:"account_name"`
	Amount       int64          `json:"amount"`
	TotalPrice   float32        `json:"total_price"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	ProductOrder []ProductOrder `json:"product_order"`
}

type ProductOrder struct {
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}
