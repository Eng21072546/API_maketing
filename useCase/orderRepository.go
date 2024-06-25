package useCase

import (
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	InsertOrder(order collection.Order) (*mongo.InsertOneResult, error)
	FindOrderById(orderId string) (*entity.Order, error)
	UpdateOrderStatus(orderID string, newStatus entity.Status) (err error)
}
