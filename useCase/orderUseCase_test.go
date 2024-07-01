package useCase

import (
	"context"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

// Mock implement Order Repository
type mockOrderRepo struct {
	insertFunc func(ctx context.Context, order collection.Order) (*mongo.InsertOneResult, error)
	findFunc   func(ctx context.Context, orderId string) (*entity.Order, error)
	updateFunc func(ctx context.Context, orderID string, newStatus entity.Status) (err error)
}

func (o *mockOrderRepo) InsertOrder(ctx context.Context, order collection.Order) (*mongo.InsertOneResult, error) {
	return o.insertFunc(ctx, order)
}

func (o mockOrderRepo) FindOrderById(ctx context.Context, orderId string) (*entity.Order, error) {
	return o.findFunc(ctx, orderId)
}

func (o mockOrderRepo) UpdateOrderStatus(ctx context.Context, orderID string, newStatus entity.Status) error {
	return o.updateFunc(ctx, orderID, newStatus)
}

// mock implement Product Repository
type mockProductRepo struct {
	insertFunc        func(ctx context.Context, product *entity.Product) (*mongo.InsertOneResult, error)
	findFunc          func(ctx context.Context, productId int) (*entity.Product, error)
	findAllFunc       func(ctx context.Context) (*[]entity.Product, error)
	updateFunc        func(ctx context.Context, id int, updateDocument bson.M) (*mongo.UpdateResult, error)
	deleteFunc        func(ctx context.Context, productId int) (*mongo.DeleteResult, error)
	updateStockFunc   func(ctx context.Context, productID int, quantity int) error
	checkStockFunc    func(ctx context.Context, productID int, quantity int) error
	decreaseStockFunc func(ctx context.Context, productOrder []entity.ProductOrder) error
}

func (p *mockProductRepo) InsertProduct(ctx context.Context, product *entity.Product) (*mongo.InsertOneResult, error) {
	return p.insertFunc(ctx, product)
}
func (p *mockProductRepo) FindProductById(ctx context.Context, productId int) (*entity.Product, error) {
	return p.findFunc(ctx, productId)
}

func (p *mockProductRepo) FindAllProducts(ctx context.Context) (*[]entity.Product, error) {
	return p.findAllFunc(ctx)
}

func (p *mockProductRepo) UpdateProduct(ctx context.Context, id int, updateDocument bson.M) (*mongo.UpdateResult, error) {
	return p.updateFunc(ctx, id, updateDocument)
}

func (p *mockProductRepo) DeleteProductById(ctx context.Context, productId int) (*mongo.DeleteResult, error) {
	return p.deleteFunc(ctx, productId)
}

func (p *mockProductRepo) UpdateStock(ctx context.Context, productID int, quantity int) error {
	return p.updateStockFunc(ctx, productID, quantity)
}

func (p *mockProductRepo) CheckStock(ctx context.Context, productID int, quantity int) error {
	return p.checkStockFunc(ctx, productID, quantity)
}
func (p *mockProductRepo) DecreaseStock(ctx context.Context, productOrders []entity.ProductOrder) error {
	return p.decreaseStockFunc(ctx, productOrders)
}

// mock implement Transaction Repository
type mockTransactionRepo struct {
	findFunc   func(ctx context.Context, id string) (*entity.Transaction, error)
	insertFunc func(ctx context.Context, transaction *collection.Transaction) (*mongo.InsertOneResult, error)
}

func (t *mockTransactionRepo) FindTransaction(ctx context.Context, transactionId string) (*entity.Transaction, error) {
	return t.findFunc(ctx, transactionId)
}

func (t *mockTransactionRepo) InsertTransaction(ctx context.Context, transaction *collection.Transaction) (*mongo.InsertOneResult, error) {
	return t.insertFunc(ctx, transaction)
}

func TestInsertOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orderRepo := &mockOrderRepo{
			insertFunc: func(ctx context.Context, order collection.Order) (*mongo.InsertOneResult, error) {
				return &mongo.InsertOneResult{}, nil
			},
			findFunc: func(ctx context.Context, orderId string) (*entity.Order, error) {
				return &entity.Order{}, nil
			},
			updateFunc: func(ctx context.Context, orderID string, newStatus entity.Status) error {
				return nil
			},
		}
		productRepo := &mockProductRepo{
			insertFunc: func(ctx context.Context, product *entity.Product) (*mongo.InsertOneResult, error) {
				return &mongo.InsertOneResult{}, nil
			},
			findFunc: func(ctx context.Context, productId int) (*entity.Product, error) {
				return &entity.Product{}, nil
			},
			updateFunc: func(ctx context.Context, id int, updateDocument bson.M) (*mongo.UpdateResult, error) {
				return &mongo.UpdateResult{}, nil
			},
			deleteFunc: func(ctx context.Context, productId int) (*mongo.DeleteResult, error) {
				return &mongo.DeleteResult{}, nil
			},
			updateStockFunc: func(ctx context.Context, productID int, quantity int) error {
				return nil
			},
			checkStockFunc: func(ctx context.Context, productID int, quantity int) error {
				return nil
			},
			decreaseStockFunc: func(ctx context.Context, product []entity.ProductOrder) error {
				return nil
			},
		}
		transactionRepo := &mockTransactionRepo{
			insertFunc: func(ctx context.Context, transaction *collection.Transaction) (*mongo.InsertOneResult, error) {
				return &mongo.InsertOneResult{}, nil
			},
			findFunc: func(ctx context.Context, id string) (*entity.Transaction, error) {
				return &entity.Transaction{}, nil
			},
		}

		orderUseCase := NewOrderUseCase(orderRepo, productRepo, transactionRepo)
		_, err := orderUseCase.NewOrder(context.Background(), &entity.Order{
			ID:            "as8923d67f45g23h23jk",
			CustomerName:  "Jame",
			Status:        1,
			TransactionId: "as8923d67f45g23h23jkf34dg",
			Transaction:   nil,
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
		})
		assert.NoError(t, err[0])
	})
}
