package useCase

import (
	"context"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository interface {
	FindTransaction(ctx context.Context, id string) (*entity.Transaction, error)
	InsertTransaction(ctx context.Context, transaction *collection.Transaction) (*mongo.InsertOneResult, error)
}
