package _interface

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error)
	GetProduct(ctx context.Context, id int) (*entity.Product, error)
	GetAllProduct(ctx context.Context) (*[]entity.Product, error)
	UpdateProduct(ctx context.Context, productUpdate *entity.ProductUpdate) (*entity.Product, error)
	DeleteProduct(ctx context.Context, id int) (*mongo.DeleteResult, error)
}
