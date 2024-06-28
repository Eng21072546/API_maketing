package collection

import (
	"github.com/Eng21072546/API_maketing/entity"
	"time"
)

type Order struct {
	ID            string        `json:"id"`
	CustomerName  string        `json:"customer_name"`
	Status        entity.Status `json:"status"`
	TransactionId string        `json:"transactionId"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

func NewOrder(orderEntity *entity.Order) Order {
	order := Order{}

	order.ID = orderEntity.ID
	order.CustomerName = orderEntity.CustomerName
	order.Status = orderEntity.Status
	order.TransactionId = orderEntity.TransactionId
	order.CreatedAt = orderEntity.CreatedAt
	order.UpdatedAt = orderEntity.UpdatedAt

	return order
}
