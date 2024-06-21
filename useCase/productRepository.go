package useCase

import (
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	InsertProduct(product *entity.Product) (*mongo.InsertOneResult, error)
	FindProductById(productId string) (*entity.Product, error)
	FindAllProducts() (*[]entity.Product, error)
	UpdateProduct(id string, productUpdate *entity.ProductUpdate) (*mongo.UpdateResult, error)
	DeleteProductById(productId string) (*mongo.DeleteResult, error)
}
