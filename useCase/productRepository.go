package useCase

import (
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	InsertProduct(product *entity.Product) (*mongo.InsertOneResult, error)
	FindProductById(productId int) (*entity.Product, error)
	FindAllProducts() (*[]entity.Product, error)
	UpdateProduct(id int, productUpdate *entity.ProductUpdate) (*mongo.UpdateResult, error)
	DeleteProductById(productId int) (*mongo.DeleteResult, error)
	UpdateStock(productID int, quantity int) error
	CheckStock(productID int, quantity int) error
	DecreaseStock(productOrder []entity.ProductOrder) error
}
