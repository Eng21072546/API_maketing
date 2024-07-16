package product

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/payload"
	"github.com/Eng21072546/API_maketing/useCase/product"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http/httptest"
	"testing"
)

func TestNewHttpProductHandler(t *testing.T) {
	mockProductUseCase := new(product.MockProductUseCase)
	productHandler := NewHttpProductHandler(mockProductUseCase)

	app := fiber.New()
	app.Get("/product", productHandler.GetAllProducts)
	app.Get("product/:id", productHandler.GetProductById)
	app.Post("/product", productHandler.CreateProduct)
	app.Put("/product/:id", productHandler.UpdateProduct)
	app.Delete("/product/:id", productHandler.DeleteProduct)

	t.Run("Get Product by id success", func(t *testing.T) {

		mockProductUseCase.On("GetProduct", mock.Anything, mock.Anything).Return(&entity.Product{
			ID:    1,
			Name:  "test",
			Price: 12.34,
			Stock: 56,
		}, nil)
		resp, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/product/1", nil))
		assert.NoError(t, err)
		utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode)
		mockProductUseCase.AssertExpectations(t)
	})
	t.Run("Get Product by id fail", func(t *testing.T) {
		mockProductUseCase.ExpectedCalls = nil
		mockProductUseCase.On("GetProduct", mock.Anything, mock.Anything).Return(&entity.Product{}, errors.New("product not found"))
		resp, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/product/1", nil))
		assert.NoError(t, err)
		utils.AssertEqual(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockProductUseCase.AssertExpectations(t)
	})
	t.Run("Get All Products success", func(t *testing.T) {
		products := &[]entity.Product{
			{
				ID:    12345,
				Name:  "test01",
				Price: 12.34,
				Stock: 56,
			},
			{
				ID:    54321,
				Name:  "test02",
				Price: 56.78,
				Stock: 43,
			},
		}
		mockProductUseCase.On("GetAllProduct", mock.Anything, mock.Anything).Return(products, nil)
		resp, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/product", nil))
		assert.NoError(t, err)
		utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode)
		utils.AssertEqual(t, []string{"application/json"}, resp.Header.Values("content-type"))
		mockProductUseCase.AssertExpectations(t)
	})
	t.Run("Get All Products fail", func(t *testing.T) {
		//products := &[]entity.Product{}
		mockProductUseCase.On("GetAllProduct", mock.Anything, mock.Anything).Return(&[]entity.Product{}, errors.New("product not found"))
		resp, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/product", nil))
		assert.NoError(t, err)
		utils.AssertEqual(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockProductUseCase.AssertExpectations(t)

	})
	t.Run("Create Product success", func(t *testing.T) {
		product := payload.ProductCreate{
			Name:  "test",
			Price: 12.34,
			Stock: 56,
		}
		productJson, err := json.Marshal(product)
		if err != nil {
			t.Fatal(err)
		}
		mockProductUseCase.On("CreateProduct", mock.Anything, mock.Anything).Return(&entity.Product{
			ID:    12345,
			Name:  "test",
			Price: 12.34,
			Stock: 56,
		}, nil)
		req := httptest.NewRequest(fiber.MethodPost, "/product", bytes.NewBuffer(productJson))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode)
		utils.AssertEqual(t, []string{"application/json"}, resp.Header.Values("content-type"))
		mockProductUseCase.AssertExpectations(t)
	})
	t.Run("Create Product fail(missing some field)", func(t *testing.T) {
		product := payload.ProductCreate{
			Name:  "test",
			Price: 12.34,
			//Stock: 56, //missing some stock field
		}
		productJson, err := json.Marshal(product)
		if err != nil {
			t.Fatal(err)
		}
		mockProductUseCase.On("CreateProduct", mock.Anything, mock.Anything).Return(&entity.Product{}, errors.New("product not found"))
		req := httptest.NewRequest(fiber.MethodPost, "/product", bytes.NewBuffer(productJson))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		utils.AssertEqual(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockProductUseCase.AssertExpectations(t)

	})
	t.Run("Update Product success", func(t *testing.T) {
		name := "test"
		price := 12.34
		stock := 56
		product := payload.ProductUpdate{
			Name:  &name,
			Price: &price,
			Stock: &stock,
		}
		productJson, err := json.Marshal(product)
		if err != nil {
			t.Fatal(err)
		}
		mockProductUseCase.On("UpdateProduct", mock.Anything, mock.Anything).Return(&entity.Product{
			ID:    12345,
			Name:  name,
			Price: price,
			Stock: stock,
		}, nil)
		req := httptest.NewRequest(fiber.MethodPut, "/product/12345", bytes.NewBuffer(productJson))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode)
		utils.AssertEqual(t, []string{"application/json"}, resp.Header.Values("content-type"))
		mockProductUseCase.AssertExpectations(t)
	})
	t.Run("Update Product fail (missing all field)", func(t *testing.T) {
		product := payload.ProductUpdate{}
		productJson, err := json.Marshal(product)
		if err != nil {
			t.Fatal(err)
		}
		mockProductUseCase.On("UpdateProduct", mock.Anything, mock.Anything).Return(&entity.Product{}, errors.New("product not found"))
		req := httptest.NewRequest(fiber.MethodPut, "/product/12345", bytes.NewBuffer(productJson))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		utils.AssertEqual(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockProductUseCase.AssertExpectations(t)
	})
	t.Run("Delete Product success", func(t *testing.T) {
		mockProductUseCase.On("DeleteProduct", mock.Anything, mock.Anything).Return(&mongo.DeleteResult{DeletedCount: 4}, nil)
		req := httptest.NewRequest(fiber.MethodDelete, "/product/0", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode)
		utils.AssertEqual(t, []string{"application/json"}, resp.Header.Values("content-type"))
		mockProductUseCase.AssertExpectations(t)
	})
	t.Run("Delete Product fail (mongo delete result = 0)", func(t *testing.T) {
		mockProductUseCase.On("DeleteProduct", mock.Anything, mock.Anything).Return(&mongo.DeleteResult{DeletedCount: 0}, errors.New("product not found"))
		req := httptest.NewRequest(fiber.MethodDelete, "/product/0", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		utils.AssertEqual(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockProductUseCase.AssertExpectations(t)

	})

}
