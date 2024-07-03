package useCase

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"time"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error)
	GetProduct(ctx context.Context, id int) (*entity.Product, error)
	GetAllProduct(ctx context.Context) (*[]entity.Product, error)
	UpdateProduct(ctx context.Context, id int, updateDocument bson.M) (*mongo.UpdateResult, error)
	DeleteProduct(ctx context.Context, id int) (*mongo.DeleteResult, error)
}

type ProductUseCaseImpl struct {
	repo ProductRepository
}

func NewProductUseCase(repo ProductRepository) ProductUseCase {
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

func (p *ProductUseCaseImpl) UpdateProduct(ctx context.Context, id int, updateDocument bson.M) (*mongo.UpdateResult, error) {
	return p.repo.UpdateProduct(ctx, id, updateDocument)
}

func (p *ProductUseCaseImpl) DeleteProduct(ctx context.Context, id int) (*mongo.DeleteResult, error) {
	return p.repo.DeleteProductById(ctx, id)
}
