package transaction

import (
	"context"
	"errors"
	"fmt"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/repo/logistic"
	"github.com/Eng21072546/API_maketing/repo/order"
	"github.com/Eng21072546/API_maketing/repo/product"
	"github.com/Eng21072546/API_maketing/repo/transaction"
)

type transactionUseCaseImpl struct {
	transactionRepo transaction.TransactionRepository
	productRepo     product.ProductRepository
	orderRepo       order.OrderRepository
	logisticRepo    logistic.LogisticCostRepo
}

func NewTransactionUseCase(transactionRepo transaction.TransactionRepository, productRepo product.ProductRepository, orderRepo order.OrderRepository, logisticRepo logistic.LogisticCostRepo) TransactionUseCase {
	return &transactionUseCaseImpl{transactionRepo, productRepo, orderRepo, logisticRepo}
}

func (t transactionUseCaseImpl) FindTransactionById(ctx context.Context, id string) (*entity.Transaction, error) {
	return t.transactionRepo.FindTransaction(ctx, id)
}

func (t transactionUseCaseImpl) NewTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, []error) {
	var errList []error
	_, err := t.logisticRepo.FindLogisticCost(ctx, transaction.Address)
	if err != nil {
		errList = append(errList, errors.New("address not found"))
	}
	if transaction.ProductOrder == nil {
		errList = append(errList, errors.New("product order empty"))
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
	transaction.TotalPrice, err = t.calculatePrice(ctx, transaction)
	transaction.Amount = len(transaction.ProductOrder)

	result, err := t.transactionRepo.InsertTransaction(ctx, collection.NewTransaction(transaction))
	if err != nil {
		errList = append(errList, err)
	}
	if len(errList) > 0 {
		return nil, errList
	}
	return result, errList
}

func (t transactionUseCaseImpl) calculatePrice(ctx context.Context, transaction *entity.Transaction) (float64, error) {
	var totalPrice float64
	//var bill []string
	logisticCost, err := t.logisticRepo.FindLogisticCost(ctx, (transaction.Address))
	if err != nil {
		return 0, err
	}
	totalPrice += logisticCost.Cost
	for _, productOrder := range transaction.ProductOrder {
		product, _ := t.productRepo.FindProductById(ctx, productOrder.ProductID)
		productPrice := product.Price * float64(productOrder.Quantity)
		//bill = append(bill, product.Name, " ", strconv.FormatFloat(product.Price, 'f', 2, 64), "  ", string(productOrder.Quantity), " ", strconv.FormatFloat(productPrice, 'f', 2, 64))
		totalPrice += productPrice
	}
	return totalPrice, nil
}
