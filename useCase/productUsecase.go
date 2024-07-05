package useCase

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"time"
)

type ProductUseCaseImpl struct {
	repo repo.ProductRepository
}

func NewProductUseCase(repo repo.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{repo: repo}
}

func (p *ProductUseCaseImpl) CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	rand.Seed(time.Now().UnixNano()) // random id product
	randomNumber := 10000 + rand.Intn(90001)
	product.ID = randomNumber
	_, err := p.repo.InsertProduct(ctx, product)
	return product, err
}

func (p *ProductUseCaseImpl) GetProduct(ctx context.Context, id int) (*entity.Product, error) {
	return p.repo.FindProductById(ctx, id)
}

func (p *ProductUseCaseImpl) GetAllProduct(ctx context.Context) (*[]entity.Product, error) {
	return p.repo.FindAllProducts(ctx)
}

func (p *ProductUseCaseImpl) UpdateProduct(ctx context.Context, productUpdate *entity.ProductUpdate) (*mongo.UpdateResult, error) {
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
	return p.repo.UpdateProduct(ctx, productUpdate.ID, updateQuery)
}

func (p *ProductUseCaseImpl) DeleteProduct(ctx context.Context, id int) (*mongo.DeleteResult, error) {
	return p.repo.DeleteProductById(ctx, id)
}
