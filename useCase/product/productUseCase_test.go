package product

import (
	"context"
	"errors"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/repo/product"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

type mockRepo struct {
	methodName string
	mock       interface{}
}

type mockProductUseCase struct {
	mockProductRepo *product.MockProductRepo
	productUseCase  ProductUseCase
}

func setupProductUseCase() mockProductUseCase {
	repo := new(product.MockProductRepo)
	useCase := NewProductUseCase(repo)
	return mockProductUseCase{repo, useCase}
}

func TestProductUseCase(t *testing.T) {

	t.Run("findProduct by id success", func(t *testing.T) {
		expect := &entity.Product{
			ID:    12345,
			Name:  "test",
			Price: 432.10,
			Stock: 12,
		}
		mockProduct := setupProductUseCase()
		mockProduct.mockProductRepo.On("FindProductById", mock.Anything, 12345).Return(expect, nil)
		result, err := mockProduct.productUseCase.GetProduct(context.TODO(), 12345)
		assert.NoError(t, err)
		assert.Equal(t, expect, result)
	})
	t.Run("findProduct by id fail", func(t *testing.T) {
		expect := &entity.Product{}
		mockProduct := setupProductUseCase()
		mockProduct.mockProductRepo.On("FindProductById", mock.Anything, 12345).Return(&entity.Product{}, errors.New("product not found"))
		result, err := mockProduct.productUseCase.GetProduct(context.TODO(), 12345)
		assert.Error(t, err)
		assert.Equal(t, expect, result)
	})
	t.Run("CreateProduct success", func(t *testing.T) {
		product := &entity.Product{
			ID:    12345,
			Name:  "test",
			Price: 123.45,
			Stock: 67,
		}
		mockProduct := setupProductUseCase()
		mockProduct.mockProductRepo.On("InsertProduct", mock.Anything, mock.Anything).Return(product, nil)
		mockProduct.mockProductRepo.On("FindProductById", mock.Anything, mock.Anything).Return(product, nil)

		result, err := mockProduct.productUseCase.CreateProduct(context.TODO(), product)
		assert.NoError(t, err)
		assert.Equal(t, product, result)
	})
	t.Run("CreateProduct fail", func(t *testing.T) {
		product := &entity.Product{
			ID:    12345,
			Name:  "test",
			Price: 123.45,
			Stock: 67,
		}
		mockProduct := setupProductUseCase()
		mockProduct.mockProductRepo.On("InsertProduct", mock.Anything, mock.Anything).Return(&entity.Product{}, errors.New("product not added"))
		mockProduct.mockProductRepo.On("FindProductById", mock.Anything, mock.Anything).Return(product, nil)
		result, err := mockProduct.productUseCase.CreateProduct(context.TODO(), product)
		assert.Error(t, err)
		assert.Equal(t, (*entity.Product)(nil), result)
	})
	t.Run("UpdateProduct success", func(t *testing.T) {
		name := "Test"
		price := 123.45
		stock := 67
		products := []struct {
			product *entity.ProductUpdate
		}{
			{&entity.ProductUpdate{ID: 12345,
				Name:  &name,
				Price: &price,
				Stock: &stock}},
			{&entity.ProductUpdate{ID: 12345,
				Name:  nil,
				Price: &price,
				Stock: &stock}},
			{&entity.ProductUpdate{ID: 12345,
				Name:  &name,
				Price: nil,
				Stock: &stock}},
			{&entity.ProductUpdate{ID: 12345,
				Name:  &name,
				Price: &price,
				Stock: nil}},
		}
		mockProduct := setupProductUseCase()
		mockProduct.mockProductRepo.On("UpdateProduct", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{
			MatchedCount:  1,
			ModifiedCount: 1,
			UpsertedCount: 0,
			UpsertedID:    nil,
		}, nil)
		mockProduct.mockProductRepo.On("FindProductById", mock.Anything, mock.Anything).Return(&entity.Product{
			ID:    123456,
			Name:  "test",
			Price: 12.34,
			Stock: 56,
		}, nil)
		for _, product := range products {
			_, err := mockProduct.productUseCase.UpdateProduct(context.TODO(), product.product)
			assert.NoError(t, err)
			//assert.Equal(t, product.product, result)
		}

	})
	t.Run("UpdateProduct fail", func(t *testing.T) {

		product := &entity.ProductUpdate{
			ID:    0,
			Name:  nil,
			Price: nil,
			Stock: nil,
		}
		mockProduct := setupProductUseCase()
		mockProduct.mockProductRepo.On("UpdateProduct", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{
			MatchedCount:  0,
			ModifiedCount: 0,
			UpsertedCount: 0,
			UpsertedID:    nil,
		}, errors.New("product not updated"))
		result, err := mockProduct.productUseCase.UpdateProduct(context.TODO(), product)
		assert.Error(t, err)
		assert.Equal(t, &mongo.UpdateResult{
			MatchedCount:  0,
			ModifiedCount: 0,
			UpsertedCount: 0,
			UpsertedID:    nil,
		}, result)
	})
	t.Run("DeleteProduct success", func(t *testing.T) {
		mockProduct := setupProductUseCase()
		mockProduct.mockProductRepo.On("DeleteProductById", mock.Anything, mock.Anything).Return(&mongo.DeleteResult{DeletedCount: 1}, nil)
		result, err := mockProduct.productUseCase.DeleteProduct(context.TODO(), 12345)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), result.DeletedCount)
	})
	t.Run("DeleteProduct fail", func(t *testing.T) {
		mockProduct := setupProductUseCase()
		mockProduct.mockProductRepo.On("DeleteProductById", mock.Anything, mock.Anything).Return(&mongo.DeleteResult{DeletedCount: 0}, errors.New("product not deleted"))
		result, err := mockProduct.productUseCase.DeleteProduct(context.TODO(), 12345)
		assert.Error(t, err)
		assert.Equal(t, int64(0), result.DeletedCount)
	})

}
