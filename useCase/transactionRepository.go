package useCase

import (
	"context"
	"github.com/Eng21072546/API_maketing/collection"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository interface {
	FindTransaction(ctx context.Context, id string) (*collection.Transaction, error)
	InsertTransaction(ctx context.Context, transaction *collection.Transaction) (*mongo.InsertOneResult, error)
}
