package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Eng21072546/API_maketing/controller/mockUseCase"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/payload"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupMockOrder() (*mockUseCase.MockOrderUseCase, *HttpOrderHandler) {
	mockOrderUseCase := &mockUseCase.MockOrderUseCase{}
	orderHandler := NewHttpOrderHandler(mockOrderUseCase)
	return mockOrderUseCase, orderHandler
}

func TestNewHttpOrderHandler(t *testing.T) {
	t.Run("should success create new http order handler", func(t *testing.T) {
		useCase, handler := setupMockOrder()
		app := fiber.New()
		app.Post("/order", handler.CreateOrder)
		transactionID := uuid.New().String()
		useCase.On("NewOrder", mock.Anything, mock.Anything).Return(&entity.Order{
			ID:            uuid.New().String(),
			CustomerName:  "test",
			Status:        1,
			TransactionId: transactionID,
			Transaction: &entity.Transaction{
				ID:           transactionID,
				Address:      "international",
				Amount:       2,
				TotalPrice:   3,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
				ProductOrder: []entity.ProductOrder{{12345, 1}, {56789, 1}},
			},
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}, []error{})
		order := payload.Order{
			CustomerName:  "test",
			TransactionId: transactionID,
		}
		orderJson, err := json.Marshal(order)
		if err != nil {
			t.Error(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewReader(orderJson))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != fiber.StatusCreated {
			assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
			t.Error("response status code is wrong")
		}
		//assert.Equal(t, transactionID, resp.Header.Get("X-Transaction-Id"))
	})
	t.Run("should fail create new http order handler (transaction not found)", func(t *testing.T) {
		useCase, handler := setupMockOrder()
		app := fiber.New()
		app.Post("/order", handler.CreateOrder)
		useCase.On("NewOrder", mock.Anything, mock.Anything).Return(&entity.Order{}, []error{errors.New("transaction not found")})
		transactionID := uuid.New().String()
		order := payload.Order{
			CustomerName:  "test",
			TransactionId: transactionID,
		}
		orderJson, err := json.Marshal(order)
		if err != nil {
			t.Error(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewReader(orderJson))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("should fail create new http order handler (transaction not found)", func(t *testing.T) {
		useCase, handler := setupMockOrder()
		app := fiber.New()
		app.Post("/order", handler.CreateOrder)
		useCase.On("NewOrder", mock.Anything, mock.Anything).Return(&entity.Order{}, []error{errors.New("transaction not found")})
		transactionID := uuid.New().String()
		order := payload.Order{
			CustomerName:  "test",
			TransactionId: transactionID,
		}
		orderJson, err := json.Marshal(order)
		if err != nil {
			t.Error(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewReader(orderJson))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("should fail create new http order handler (request missing some field)", func(t *testing.T) {
		useCase, handler := setupMockOrder()
		app := fiber.New()
		app.Post("/order", handler.CreateOrder)
		useCase.On("NewOrder", mock.Anything, mock.Anything).Return(&entity.Order{}, []error{errors.New("transaction not found")})

		order := payload.Order{
			CustomerName:  "",
			TransactionId: "",
		}
		orderJson, err := json.Marshal(order)
		if err != nil {
			t.Error(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/order", bytes.NewReader(orderJson))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

}
