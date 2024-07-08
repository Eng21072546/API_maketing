package mockRepository

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockProductRepo struct {
	mock.Mock
}

func (m *MockProductRepo) FindProductById(ctx context.Context, productId int) (*entity.Product, error) {
	args := m.Called(ctx, productId)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *MockProductRepo) InsertProduct(ctx context.Context, product *entity.Product) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, product)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockProductRepo) FindAllProducts(ctx context.Context) (*[]entity.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).(*[]entity.Product), args.Error(1)
}

func (m *MockProductRepo) UpdateProduct(ctx context.Context, id int, updateDocument bson.M) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, id, updateDocument)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockProductRepo) DeleteProductById(ctx context.Context, productId int) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, productId)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func (m *MockProductRepo) UpdateStock(ctx context.Context, productID int, quantity int) error {
	args := m.Called(ctx, productID, quantity)
	return args.Error(0)
}

func (m *MockProductRepo) CheckStock(ctx context.Context, productID int, quantity int) error {
	args := m.Called(ctx, productID, quantity)
	return args.Error(0)
}

func (m *MockProductRepo) DecreaseStock(ctx context.Context, productOrder []entity.ProductOrder) error {
	args := m.Called(ctx, productOrder)
	return args.Error(0)
}

func (m *MockProductRepo) GenProductID() int {
	args := m.Called()
	return args.Get(0).(int)
}
