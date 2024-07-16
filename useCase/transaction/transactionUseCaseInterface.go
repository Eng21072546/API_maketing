package _interface

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
)

type TransactionUseCase interface {
	NewTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, []error)
	FindTransactionById(context context.Context, id string) (*entity.Transaction, error)
}
