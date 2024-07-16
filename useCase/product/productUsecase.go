package product

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/repo/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
)

var (
	name   = "productUseCase"
	tracer = otel.GetTracerProvider().Tracer(name)
)

type ProductUseCaseImpl struct {
	repo product.ProductRepository
}

func NewProductUseCase(repo product.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{repo: repo}
}

func (p *ProductUseCaseImpl) CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	ctx, span := tracer.Start(ctx, "CreateProduct")
	defer span.End()
	result, err := p.repo.InsertProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProductUseCaseImpl) GetProduct(ctx context.Context, id int) (*entity.Product, error) {
	ctx, span := tracer.Start(ctx, "GetProduct")
	defer span.End()
	return p.repo.FindProductById(ctx, id)
}

func (p *ProductUseCaseImpl) GetAllProduct(ctx context.Context) (*[]entity.Product, error) {
	ctx, span := tracer.Start(ctx, "GetAllProduct")
	defer span.End()
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
