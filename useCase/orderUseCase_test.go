package useCase

import (
	"context"
	"errors"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/useCase/interface"
	"github.com/Eng21072546/API_maketing/useCase/mockRepository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

type mockOrder struct {
	useCase         _interface.OrderUseCase
	repo            *mockRepository.MockOrderRepo
	productRepo     *mockRepository.MockProductRepo
	transactionRepo *mockRepository.MockTransactionRepo
	logisticRepo    *mockRepository.MockLogisticRepo
}

func setupOrderTest() mockOrder {
	pd, tx, or, lo := mockRepository.NewMockRepository()
	useCase := NewOrderUseCase(or, pd, tx)
	return mockOrder{useCase, or, pd, tx, lo}
}

func TestOrderUseCaseImpl(t *testing.T) {
	t.Run("New Order success", func(t *testing.T) {
		timeStamp := time.Now()
		transaction := &entity.Transaction{
			ID:         uuid.New().String(),
			Address:    "international address",
			Amount:     0,
			TotalPrice: 0,
			CreatedAt:  timeStamp,
			UpdatedAt:  timeStamp,
			ProductOrder: []entity.ProductOrder{
				{12345, 1}, {56789, 1},
			},
		}
		order := &entity.Order{
			ID:            "",
			CustomerName:  "test",
			Status:        0,
			TransactionId: transaction.ID,
			Transaction:   nil,
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
		}
		want := &entity.Order{
			ID:            uuid.New().String(),
			CustomerName:  "test",
			Status:        1,
			TransactionId: transaction.ID,
			Transaction:   transaction,
			CreatedAt:     timeStamp,
			UpdatedAt:     timeStamp,
		}

		orderMock := setupOrderTest()
		orderMock.repo.On("InsertOrder", mock.Anything, mock.Anything).Return(&entity.Order{
			ID:            want.ID,
			CustomerName:  "test",
			Status:        1,
			TransactionId: transaction.ID,
			Transaction:   transaction,
			CreatedAt:     timeStamp,
			UpdatedAt:     timeStamp,
		}, nil)
		orderMock.transactionRepo.On("FindTransaction", mock.Anything, mock.Anything).Return(transaction, nil)
		orderMock.productRepo.On("CheckStock", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		orderMock.productRepo.On("DecreaseStock", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		result, err := orderMock.useCase.NewOrder(context.TODO(), order)

		assert.Equal(t, []error(nil), err)
		assert.Equal(t, want, result)
	})
	t.Run("New Order failure", func(t *testing.T) {
		timeStamp := time.Now()
		transaction := &entity.Transaction{
			ID:         uuid.New().String(),
			Address:    "international address",
			Amount:     0,
			TotalPrice: 0,
			CreatedAt:  timeStamp,
			UpdatedAt:  timeStamp,
			ProductOrder: []entity.ProductOrder{
				{12345, 1}, {56789, 1},
			},
		}
		order := &entity.Order{
			ID:            "",
			CustomerName:  "test",
			Status:        0,
			TransactionId: transaction.ID,
			Transaction:   nil,
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
		}
		want := &entity.Transaction{}

		orderMock := setupOrderTest()
		orderMock.repo.On("InsertOrder", mock.Anything, mock.Anything).Return(&entity.Order{
			ID:            want.ID,
			CustomerName:  "test",
			Status:        1,
			TransactionId: transaction.ID,
			Transaction:   transaction,
			CreatedAt:     timeStamp,
			UpdatedAt:     timeStamp,
		}, nil)
		orderMock.transactionRepo.On("FindTransaction", mock.Anything, mock.Anything).Return(&entity.Transaction{}, errors.New("transaction not found"))
		orderMock.productRepo.On("CheckStock", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		orderMock.productRepo.On("DecreaseStock", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		result, err := orderMock.useCase.NewOrder(context.TODO(), order)

		assert.Error(t, err[0])
		assert.Equal(t, (*entity.Order)(nil), result)
	})
	t.Run("Patch Status success", func(t *testing.T) {
		timeStamp := time.Now()
		transaction := &entity.Transaction{
			ID:         uuid.New().String(),
			Address:    "international address",
			Amount:     0,
			TotalPrice: 0,
			CreatedAt:  timeStamp,
			UpdatedAt:  timeStamp,
			ProductOrder: []entity.ProductOrder{
				{12345, 1}, {56789, 1},
			},
		}
		order := &entity.Order{
			ID:            uuid.New().String(),
			CustomerName:  "test",
			Status:        1,
			TransactionId: transaction.ID,
			Transaction:   transaction,
			CreatedAt:     timeStamp,
			UpdatedAt:     timeStamp,
		}
		want := &entity.Order{
			ID:            order.ID,
			CustomerName:  order.CustomerName,
			Status:        2,
			TransactionId: transaction.ID,
			Transaction:   transaction,
			CreatedAt:     timeStamp,
			UpdatedAt:     timeStamp,
		}
		orderMock := setupOrderTest()
		orderMock.repo.On("FindOrderById", mock.Anything, mock.Anything).Return(order, nil)
		orderMock.transactionRepo.On("FindTransaction", mock.Anything, mock.Anything).Return(transaction, nil)
		orderMock.repo.On("UpdateOrderStatus", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{
			MatchedCount:  1,
			ModifiedCount: 1,
			UpsertedCount: 0,
			UpsertedID:    nil,
		}, nil)

		result, err := orderMock.useCase.PatchOrderStatus(context.TODO(), order.ID)
		assert.Equal(t, want, result)
		assert.Nil(t, err)
	})
	t.Run("Patch Status failure (order not found)", func(t *testing.T) {})
	orderMock := setupOrderTest()
	orderMock.repo.On("FindOrderById", mock.Anything, mock.Anything).Return(&entity.Order{}, errors.New("order not found"))
	_, err := orderMock.useCase.PatchOrderStatus(context.TODO(), mock.Anything)
	assert.Error(t, err)

}
