package entity

import "time"

type Transaction struct {
	ID           string
	Address      string
	Amount       int
	TotalPrice   float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ProductOrder []ProductOrder
}

type ProductOrder struct {
	ProductID int
	Quantity  int
}

func NewTransaction(address string, productOrder []ProductOrder) *Transaction {
	return &Transaction{Address: address, ProductOrder: productOrder}
}
