package repo

import (
	"context"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogisticCostRepo interface {
	FindLogisticCost(ctx context.Context, address string) (*entity.LogisticCost, error)
	InsertLogisticCost(ctx context.Context, cost collection.LogisticCost) (*mongo.InsertOneResult, error)
}
