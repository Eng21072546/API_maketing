package transaction

import (
	"context"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
)

type TransactionRepository interface {
	FindTransaction(ctx context.Context, id string) (*entity.Transaction, error)
	InsertTransaction(ctx context.Context, transaction *collection.Transaction) (*entity.Transaction, error)
}
