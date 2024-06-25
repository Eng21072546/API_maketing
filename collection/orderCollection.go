package collection

import (
	"github.com/Eng21072546/API_maketing/entity"
	"time"
)

type Order struct {
	ID           string        `json:"id"`
	CustomerName string        `json:"customer_name"`
	Status       entity.Status `json:"status"`
	Transaction  []string      `json:"transaction"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

func NewOrder(orderEntity *entity.Order) Order {
	order := Order{}
	var transactionId []string
	order.ID = orderEntity.ID
	order.CustomerName = orderEntity.CustomerName
	order.Status = orderEntity.Status
	for _, transaction := range orderEntity.Transaction {
		transactionId = append(transactionId, transaction.ID)
	}
	order.Transaction = transactionId
	return order
}
