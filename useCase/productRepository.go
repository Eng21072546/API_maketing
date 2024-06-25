package useCase

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	InsertProduct(ctx context.Context, product *entity.Product) (*mongo.InsertOneResult, error)
	FindProductById(ctx context.Context, productId int) (*entity.Product, error)
	FindAllProducts(ctx context.Context) (*[]entity.Product, error)
	UpdateProduct(ctx context.Context, id int, productUpdate *entity.ProductUpdate) (*mongo.UpdateResult, error)
	DeleteProductById(ctx context.Context, productId int) (*mongo.DeleteResult, error)
	UpdateStock(ctx context.Context, productID int, quantity int) error
	CheckStock(ctx context.Context, productID int, quantity int) error
	DecreaseStock(ctx context.Context, productOrder []entity.ProductOrder) error
}
