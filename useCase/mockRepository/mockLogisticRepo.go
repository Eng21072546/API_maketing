package mockRepository

import (
	"context"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockLogisticRepo struct {
	mock.Mock
}

func (l *MockLogisticRepo) FindLogisticCost(ctx context.Context, address string) (*entity.LogisticCost, error) {
	args := l.Called(ctx, address)
	return args.Get(0).(*entity.LogisticCost), args.Error(1)
}

func (l *MockLogisticRepo) InsertLogisticCost(ctx context.Context, cost collection.LogisticCost) (*mongo.InsertOneResult, error) {
	args := l.Called(ctx, cost)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}
