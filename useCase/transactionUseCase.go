package useCase

import (
	"context"
	"errors"
	"fmt"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type TransactionUseCase interface {
	NewTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, []error)
	FindTransactionById(context context.Context, id string) (*entity.Transaction, error)
}

type transactionUseCaseImpl struct {
	transactionRepo TransactionRepository
	productRepo     ProductRepository
	orderRepo       OrderRepository
}

func NewTransactionUseCase(transactionRepo TransactionRepository, productRepo ProductRepository, orderRepo OrderRepository) TransactionUseCase {
	return &transactionUseCaseImpl{transactionRepo, productRepo, orderRepo}
}

func (t transactionUseCaseImpl) FindTransactionById(ctx context.Context, id string) (*entity.Transaction, error) {
	return t.transactionRepo.FindTransaction(ctx, id)
}

func (t transactionUseCaseImpl) NewTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, []error) {
	var errList []error
	err := entity.CheckAddress(transaction.Address)
	if err != nil {
		errList = append(errList, err)
	}
	for _, productOrder := range transaction.ProductOrder {
		_, err = t.productRepo.FindProductById(ctx, productOrder.ProductID)
		if err != nil {
			errList = append(errList, errors.New(fmt.Sprintf("product ID %d not found", productOrder.ProductID)))
		}
	}
	if len(errList) > 0 { // if product in the ProductOrder list is not match in DB, Stop and return []ERR
		return nil, errList
	}
	transaction.ID = uuid.New().String()
	transaction.TotalPrice = t.CalculatePrice(ctx, transaction)
	transaction.Amount = len(transaction.ProductOrder)
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	_, err = t.transactionRepo.InsertTransaction(ctx, collection.NewTransaction(transaction))
	if err != nil {
		errList = append(errList, err)
	}
	if len(errList) > 0 {
		return nil, errList
	}
	return transaction, errList
}

func (t transactionUseCaseImpl) CalculatePrice(ctx context.Context, transaction *entity.Transaction) float64 {
	var totalPrice float64
	var bill []string
	logisticPrice, _ := entity.LogisticCost(transaction.Address)
	totalPrice += logisticPrice
	for _, productOrder := range transaction.ProductOrder {
		product, _ := t.productRepo.FindProductById(ctx, productOrder.ProductID)
		productPrice := product.Price * float64(productOrder.Quantity)
		bill = append(bill, product.Name, " ", strconv.FormatFloat(product.Price, 'f', 2, 64), "  ", string(productOrder.Quantity), " ", strconv.FormatFloat(productPrice, 'f', 2, 64))
		totalPrice += productPrice
	}
	return totalPrice
}
