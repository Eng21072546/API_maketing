package controller

import (
	"bytes"
	"encoding/json"
	"github.com/Eng21072546/API_maketing/controller/mockUseCase"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/payload"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

func SetupMockTransaction() (*mockUseCase.MockTransactionUseCase, *HttpTransactionHandler) {
	mockTransactionUseCase := new(mockUseCase.MockTransactionUseCase)
	transactionHandler := NewHttpTransactionHandler(mockTransactionUseCase)
	return mockTransactionUseCase, transactionHandler
}

func TestNewHttpTransactionHandler(t *testing.T) {

	t.Run("should create new transaction handler", func(t *testing.T) {
		mockTransactionUseCase, transactionHandler := SetupMockTransaction()
		app := fiber.New()
		app.Post("/order/calculation", transactionHandler.PostTransaction)
		mockTransactionUseCase.On("NewTransaction", mock.Anything, mock.Anything).Return(&entity.Transaction{}, []error{})

		transaction := &payload.Transaction{
			AccountName: "test",
			Address:     "international",
			ProductOrder: []entity.ProductOrder{
				{12345, 1}, {56789, 2},
			},
		}
		transactionJson, err := json.Marshal(transaction)
		if err != nil {
			t.Error(err)
		}
		req, err := http.NewRequest(http.MethodPost, "/order/calculation", bytes.NewBuffer(transactionJson))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Error(err)
		}
		resp, err := app.Test(req)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
		//assert.Equal(t, body, transactionJson)
		if resp.StatusCode != http.StatusCreated {
			t.Error(string(body))
		}
		mockTransactionUseCase.AssertExpectations(t)

	})
	t.Run("should return error when create transaction fails (Invalid request body)", func(t *testing.T) {
		mockTransactionUseCase, transactionHandler := SetupMockTransaction()
		app := fiber.New()
		app.Post("/order/calculation", transactionHandler.PostTransaction)
		mockTransactionUseCase.On("NewTransaction", mock.Anything, mock.Anything).Return(&entity.Transaction{}, []error{})
		transaction := &payload.Transaction{
			AccountName:  "test",
			Address:      "international",
			ProductOrder: nil,
		}
		transactionJson, err := json.Marshal(transaction)
		if err != nil {
			t.Error(err)
		}
		req, err := http.NewRequest(http.MethodPost, "/order/calculation", bytes.NewBuffer(transactionJson))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Error(err)
		}
		resp, err := app.Test(req)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
		if resp.StatusCode != http.StatusBadRequest {
			t.Error(string(body))
		}
	})
}
