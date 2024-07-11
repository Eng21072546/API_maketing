package _interface

import (
	"context"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"time"
)

type TransactionRepository interface {
	FindTransaction(ctx context.Context, id string) (*entity.Transaction, error)
	InsertTransaction(ctx context.Context, transaction *collection.Transaction) (*entity.Transaction, error)
	SetTime() time.Time
	SetId() string
}
