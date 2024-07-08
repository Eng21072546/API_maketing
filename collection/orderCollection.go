package collection

import (
	"github.com/Eng21072546/API_maketing/entity"
	"time"
)

type Order struct {
	ID            string
	CustomerName  string
	Status        entity.Status
	TransactionId string
	CreatedAt     time.Time
	UpdatedAt     time.Time
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
