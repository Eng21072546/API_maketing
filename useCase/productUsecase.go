package useCase

import (
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"time"
)

type ProductUseCase interface {
	CreateProduct(product *entity.Product) (*entity.Product, error)
	GetProduct(id int) (*entity.Product, error)
	GetAllProduct() (*[]entity.Product, error)
	UpdateProduct(id int, productUpdate *entity.ProductUpdate) (*mongo.UpdateResult, error)
	DeleteProduct(id int) (*mongo.DeleteResult, error)
}

type ProductUseCaseImpl struct {
	repo ProductRepository
}

func NewProductUseCase(repo ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{repo: repo}
}

func (p *ProductUseCaseImpl) CreateProduct(product *entity.Product) (*entity.Product, error) {
	rand.Seed(time.Now().UnixNano()) // random id product
	randomNumber := 10000 + rand.Intn(90001)
	product.ID = randomNumber
	_, err := p.repo.InsertProduct(product)
	return product, err
}

func (p *ProductUseCaseImpl) GetProduct(id int) (*entity.Product, error) {
	return p.repo.FindProductById(id)
}

func (p *ProductUseCaseImpl) GetAllProduct() (*[]entity.Product, error) {
	return p.repo.FindAllProducts()
}

func (p *ProductUseCaseImpl) UpdateProduct(id int, productUpdate *entity.ProductUpdate) (*mongo.UpdateResult, error) {
	return p.repo.UpdateProduct(id, productUpdate)
}

func (p *ProductUseCaseImpl) DeleteProduct(id int) (*mongo.DeleteResult, error) {
	return p.repo.DeleteProductById(id)
}
