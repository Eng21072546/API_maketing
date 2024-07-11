package mockRepository

import (
	"context"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockOrderRepo struct {
	mock.Mock
}

func (m *MockOrderRepo) InsertOrder(ctx context.Context, order collection.Order) (*entity.Order, error) {
	args := m.Called(ctx, order)
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (m *MockOrderRepo) FindOrderById(ctx context.Context, orderId string) (*entity.Order, error) {
	args := m.Called(ctx, orderId)
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (m *MockOrderRepo) UpdateOrderStatus(ctx context.Context, orderID string, newStatus entity.Status) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, orderID, newStatus)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}
