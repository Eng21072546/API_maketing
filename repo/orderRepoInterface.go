package repo

import (
	"context"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type OrderRepository interface {
	InsertOrder(ctx context.Context, order collection.Order) (*entity.Order, error)
	FindOrderById(ctx context.Context, orderId string) (*entity.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID string, newStatus entity.Status) (*mongo.UpdateResult, error)
	SetTime() time.Time
	SetId() string
}
