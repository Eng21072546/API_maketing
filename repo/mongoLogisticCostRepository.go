package repo

import (
	"context"
	"fmt"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoLogisticCostRepo struct {
	client   *mongo.Client
	ctxMongo context.Context
}

func NewMongoLogisticRepository(client *mongo.Client, ctx context.Context) LogisticCostRepo {
	return &MongoLogisticCostRepo{client, ctx}
}

// InsertLogisticCost this function for use with CRUD in the future.
func (l MongoLogisticCostRepo) InsertLogisticCost(ctx context.Context, cost collection.LogisticCost) (*mongo.InsertOneResult, error) {
	result, err := l.client.Database("market").Collection("logisticCost").InsertOne(l.ctxMongo, cost)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (l MongoLogisticCostRepo) FindLogisticCost(ctx context.Context, address string) (*entity.LogisticCost, error) {
	filter := bson.D{{"address", address}}
	result := l.client.Database("market").Collection("logisticCost").FindOne(l.ctxMongo, filter)
	logisticCost := &entity.LogisticCost{}
	err := result.Decode(logisticCost)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		fmt.Println(err.Error())
		return nil, err
	}

	return logisticCost, nil
}
