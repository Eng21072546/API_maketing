package product

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockProductUseCase struct {
	mock.Mock
}

func (p *MockProductUseCase) CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	args := p.Called(ctx, product)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (p *MockProductUseCase) GetProduct(ctx context.Context, id int) (*entity.Product, error) {
	args := p.Called(ctx, id)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (p *MockProductUseCase) GetAllProduct(ctx context.Context) (*[]entity.Product, error) {
	args := p.Called(ctx)
	if args.Get(1) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entity.Product), args.Error(1)
}

func (p *MockProductUseCase) UpdateProduct(ctx context.Context, productUpdate *entity.ProductUpdate) (*entity.Product, error) {
	args := p.Called(ctx, productUpdate)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (p *MockProductUseCase) DeleteProduct(ctx context.Context, id int) (*mongo.DeleteResult, error) {
	args := p.Called(ctx, id)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}
