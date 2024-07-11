package useCase

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/repo/interface"
	_interface2 "github.com/Eng21072546/API_maketing/useCase/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductUseCaseImpl struct {
	repo _interface.ProductRepository
}

func NewProductUseCase(repo _interface.ProductRepository) _interface2.ProductUseCase {
	return &ProductUseCaseImpl{repo: repo}
}

func (p *ProductUseCaseImpl) CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	result, err := p.repo.InsertProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProductUseCaseImpl) GetProduct(ctx context.Context, id int) (*entity.Product, error) {
	return p.repo.FindProductById(ctx, id)
}

func (p *ProductUseCaseImpl) GetAllProduct(ctx context.Context) (*[]entity.Product, error) {
	return p.repo.FindAllProducts(ctx)
}

func (p *ProductUseCaseImpl) UpdateProduct(ctx context.Context, productUpdate *entity.ProductUpdate) (*entity.Product, error) {
	updateQuery := make(bson.M)
	if productUpdate.Name != nil {
		updateQuery["name"] = productUpdate.Name
	}
	if productUpdate.Price != nil {
		updateQuery["price"] = productUpdate.Price
	}
	if productUpdate.Stock != nil {
		updateQuery["stock"] = productUpdate.Stock
	}
	_, err := p.repo.UpdateProduct(ctx, productUpdate.ID, updateQuery)
	if err != nil {
		return nil, err
	}
	result, err := p.repo.FindProductById(ctx, productUpdate.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProductUseCaseImpl) DeleteProduct(ctx context.Context, id int) (*mongo.DeleteResult, error) {
	return p.repo.DeleteProductById(ctx, id)
}
