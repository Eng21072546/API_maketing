package useCase

import (
	"context"
	"errors"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestInsertProduct(t *testing.T) {
	t.Run("should insert product", func(t *testing.T) {
		productRepo := &mockProductRepo{
			insertFunc: func(ctx context.Context, product *entity.Product) (*mongo.InsertOneResult, error) {
				return &mongo.InsertOneResult{}, nil
			},
		}
		productUseCase := NewProductUseCase(productRepo)
		_, err := productUseCase.CreateProduct(context.Background(), &entity.Product{
			ID:    12345,
			Name:  "Test",
			Price: 40,
			Stock: 50,
		})
		assert.NoError(t, err)
	})
	t.Run("should return error when insert product fails", func(t *testing.T) {
		productRepo := &mockProductRepo{
			insertFunc: func(ctx context.Context, product *entity.Product) (*mongo.InsertOneResult, error) {
				return &mongo.InsertOneResult{}, errors.New("error inserting product")
			},
		}
		productUseCase := NewProductUseCase(productRepo)
		_, err := productUseCase.CreateProduct(context.Background(), &entity.Product{
			ID:    12345,
			Name:  "Test",
			Price: 40,
			Stock: 50,
		})
		assert.Error(t, err)
	})
}
func TestGetProduct(t *testing.T) {
	t.Run("should get product", func(t *testing.T) {

		productRepo := &mockProductRepo{
			findFunc: func(ctx context.Context, productId int) (*entity.Product, error) {
				return &entity.Product{
					ID:    12345,
					Name:  "Test",
					Price: 40,
					Stock: 50,
				}, nil
			},
		}
		productUseCase := NewProductUseCase(productRepo)
		product, err := productUseCase.GetProduct(context.Background(), 12345)
		assert.NoError(t, err)
		assert.Equal(t, &entity.Product{
			ID:    12345,
			Name:  "Test",
			Price: 40,
			Stock: 50,
		}, product)
	})
	t.Run("should return error when get product fails", func(t *testing.T) {
		productRepo := &mockProductRepo{
			findFunc: func(ctx context.Context, productId int) (*entity.Product, error) {
				return nil, errors.New("error getting product")
			},
		}
		productUseCase := NewProductUseCase(productRepo)
		_, err := productUseCase.GetProduct(context.Background(), 12345)
		assert.Error(t, err)
	})
}

func TestUpdateProduct(t *testing.T) {
	name := "test"
	price := 40.00
	stock := 50
	emptyStock := -1
	//Test case
	tests := []struct {
		name        string
		product     *entity.ProductUpdate
		wantProduct *entity.Product
		wantErr     error
	}{
		{
			name: "should update product",
			product: &entity.ProductUpdate{
				ID:    12345,
				Name:  &name,
				Price: &price,
				Stock: &stock,
			},
			wantProduct: nil,
			wantErr:     nil,
		},
		{
			name: "should update product with only name",
			product: &entity.ProductUpdate{
				ID:    12345,
				Name:  &name,
				Price: nil,
				Stock: nil,
			},
			wantProduct: nil,
			wantErr:     nil,
		},
		{
			name: "should update product with only price",
			product: &entity.ProductUpdate{
				ID:    12345,
				Name:  nil,
				Price: &price,
				Stock: nil,
			},
			wantProduct: nil,
			wantErr:     nil,
		},
		{
			name: "should update product with only stock",
			product: &entity.ProductUpdate{
				ID:    12345,
				Name:  nil,
				Price: nil,
				Stock: &stock,
			},
			wantProduct: nil,
			wantErr:     nil,
		},
		{
			name: "should return error when update product with nil product ",
			product: &entity.ProductUpdate{
				ID:    12345,
				Name:  nil,
				Price: nil,
				Stock: nil,
			},
			wantProduct: nil,
			wantErr:     errors.New("product cannot be nil"),
		},
		{
			name: "should return error when stock negative ",
			product: &entity.ProductUpdate{
				ID:    12345,
				Name:  nil,
				Price: nil,
				Stock: &emptyStock,
			},
			wantProduct: nil,
			wantErr:     errors.New("product stock cannot be negative"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := &mockProductRepo{
				updateFunc: func(ctx context.Context, id int, updateDocument bson.M) (*mongo.UpdateResult, error) {
					return &mongo.UpdateResult{}, nil
				},
			}
			productUseCase := NewProductUseCase(productRepo)
			_, err := productUseCase.UpdateProduct(context.Background(), tt.product)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func TestDeleteProduct(t *testing.T) {
	t.Run("should delete product", func(t *testing.T) {
		productRepo := &mockProductRepo{
			deleteFunc: func(ctx context.Context, id int) (*mongo.DeleteResult, error) {
				return &mongo.DeleteResult{DeletedCount: 1}, nil
			},
		}
		productUseCase := NewProductUseCase(productRepo)
		_, err := productUseCase.DeleteProduct(context.Background(), 12345)
		assert.NoError(t, err)
	})

	t.Run("should return error when delete product fails", func(t *testing.T) {
		productRepo := &mockProductRepo{
			deleteFunc: func(ctx context.Context, id int) (*mongo.DeleteResult, error) {
				return nil, errors.New("error deleting product")
			},
		}
		productUseCase := NewProductUseCase(productRepo)
		_, err := productUseCase.DeleteProduct(context.Background(), 12345)
		assert.Error(t, err)
	})

}
