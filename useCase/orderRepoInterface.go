package useCase

import (
	"context"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	InsertOrder(ctx context.Context, order collection.Order) (*mongo.InsertOneResult, error)
	FindOrderById(ctx context.Context, orderId string) (*collection.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID string, newStatus entity.Status) (err error)
}
