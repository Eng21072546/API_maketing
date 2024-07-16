package useCase

import (
	"github.com/Eng21072546/API_maketing/useCase/order"
	"github.com/Eng21072546/API_maketing/useCase/product"
	"github.com/Eng21072546/API_maketing/useCase/transaction"
)

func NewMockUseCase() (*product.MockProductUseCase, *transaction.MockTransactionUseCase, *order.MockOrderUseCase) {
	return new(product.MockProductUseCase), new(transaction.MockTransactionUseCase), new(order.MockOrderUseCase)
}
