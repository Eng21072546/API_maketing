package repo

import (
	"github.com/Eng21072546/API_maketing/repo/logistic"
	"github.com/Eng21072546/API_maketing/repo/order"
	"github.com/Eng21072546/API_maketing/repo/product"
	"github.com/Eng21072546/API_maketing/repo/transaction"
)

func NewMockRepository() (*product.MockProductRepo, *transaction.MockTransactionRepo, *order.MockOrderRepo, *logistic.MockLogisticRepo) {
	return new(product.MockProductRepo), new(transaction.MockTransactionRepo), new(order.MockOrderRepo), new(logistic.MockLogisticRepo)
}
