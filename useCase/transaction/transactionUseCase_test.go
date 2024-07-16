package transaction

import (
	"context"
	"errors"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/repo"
	"github.com/Eng21072546/API_maketing/repo/logistic"
	"github.com/Eng21072546/API_maketing/repo/order"
	"github.com/Eng21072546/API_maketing/repo/product"
	"github.com/Eng21072546/API_maketing/repo/transaction"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
	"time"
)

type mockTransaction struct {
	useCase      TransactionUseCase
	repo         *transaction.MockTransactionRepo
	productRepo  *product.MockProductRepo
	orderRepo    *order.MockOrderRepo
	logisticRepo *logistic.MockLogisticRepo
}

func setupTransactionTest() mockTransaction {
	pd, tx, or, lo := repo.NewMockRepository()
	useCase := NewTransactionUseCase(tx, pd, or, lo)
	return mockTransaction{
		useCase:      useCase,
		repo:         tx,
		productRepo:  pd,
		orderRepo:    or,
		logisticRepo: lo,
	}
}

func TestTransactionUseCase(t *testing.T) {
	t.Run("FindTransactionById success", func(t *testing.T) {
		transaction := &entity.Transaction{
			ID:         "adc1234",
			Address:    "International",
			Amount:     1234,
			TotalPrice: 5678.90,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			ProductOrder: []entity.ProductOrder{
				entity.ProductOrder{
					ProductID: 123456,
					Quantity:  12,
				},
				entity.ProductOrder{
					ProductID: 654321,
					Quantity:  34,
				},
			},
		}
		mockTran := setupTransactionTest()
		mockTran.repo.On("FindTransaction", mock.Anything, "abc1234").Return(transaction, nil)

		result, err := mockTran.useCase.FindTransactionById(context.TODO(), "abc1234")
		assert.NoError(t, err)
		assert.Equal(t, transaction, result)

	})
	t.Run("FindTransactionById fail", func(t *testing.T) {
		mockTran := setupTransactionTest()
		mockTran.repo.On("FindTransaction", mock.Anything, "abc1234").Return(&entity.Transaction{}, errors.New("transaction not found"))
		result, err := mockTran.useCase.FindTransactionById(context.TODO(), "abc1234")
		assert.Error(t, err)
		assert.Equal(t, &entity.Transaction{}, result)
	})
	t.Run("InsertTransaction success", func(t *testing.T) {
		timeStamp := time.Now()
		Id := uuid.New()
		transaction := &entity.Transaction{
			ID:         "",
			Address:    "international",
			Amount:     0,
			TotalPrice: 0,
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
			ProductOrder: []entity.ProductOrder{
				entity.ProductOrder{
					ProductID: 123456,
					Quantity:  1,
				},
				entity.ProductOrder{
					ProductID: 654321,
					Quantity:  1,
				},
			},
		}
		want := &entity.Transaction{
			ID:           Id.String(),
			Address:      transaction.Address,
			Amount:       len(transaction.ProductOrder),
			TotalPrice:   40.00,
			CreatedAt:    timeStamp,
			UpdatedAt:    timeStamp,
			ProductOrder: transaction.ProductOrder,
		}
		mockTran := setupTransactionTest()
		mockTran.repo.On("InsertTransaction", mock.Anything, mock.Anything).Return(&entity.Transaction{
			ID:           Id.String(),
			Address:      transaction.Address,
			Amount:       len(transaction.ProductOrder),
			TotalPrice:   40.00,
			CreatedAt:    timeStamp,
			UpdatedAt:    timeStamp,
			ProductOrder: transaction.ProductOrder,
		}, nil)
		mockTran.repo.On("FindTransaction", mock.Anything, "abc1234").Return(nil, nil)
		mockTran.repo.On("SetTime").Return(timeStamp)
		mockTran.repo.On("SetId").Return(Id.String())
		mockTran.productRepo.On("FindProductById", mock.Anything, mock.Anything).Return(&entity.Product{
			ID:    1234,
			Name:  "test",
			Price: 10.00,
			Stock: 1,
		}, nil)
		mockTran.logisticRepo.On("FindLogisticCost", mock.Anything, mock.Anything).Return(&entity.LogisticCost{
			Address: "International",
			Cost:    20,
		}, nil)

		result, err := mockTran.useCase.NewTransaction(context.TODO(), transaction)

		assert.Equal(t, []error(nil), err)
		assert.Equal(t, want, result)

	})
	t.Run("InsertTransaction fail", func(t *testing.T) {
		transaction := &entity.Transaction{
			ID:         "",
			Address:    "BKK",
			Amount:     0,
			TotalPrice: 0,
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
			ProductOrder: []entity.ProductOrder{
				entity.ProductOrder{
					ProductID: 123456,
					Quantity:  1,
				},
				entity.ProductOrder{
					ProductID: 654321,
					Quantity:  1,
				},
			},
		}
		mockTran := setupTransactionTest()
		mockTran.repo.On("InsertTransaction", mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{nil}, errors.New("transaction error"))
		mockTran.logisticRepo.On("FindLogisticCost", mock.Anything, mock.Anything).Return(&entity.LogisticCost{}, errors.New("address not found"))
		mockTran.productRepo.On("FindProductById", mock.Anything, mock.Anything).Return(&entity.Product{
			ID:    1234,
			Name:  "test",
			Price: 10.00,
			Stock: 10,
		}, nil)
		result, err := mockTran.useCase.NewTransaction(context.TODO(), transaction)
		assert.Equal(t, (*entity.Transaction)(nil), result)
		assert.Error(t, err[0])
	})

}
