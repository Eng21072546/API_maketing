package transaction

import (
	"context"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoTransactionRepository struct {
	client   *mongo.Client
	ctxMongo context.Context
}

func NewMongoTransactionRepository(client *mongo.Client, ctx context.Context) *MongoTransactionRepository {
	return &MongoTransactionRepository{client, ctx}
}

func (t MongoTransactionRepository) InsertTransaction(ctx context.Context, transaction *collection.Transaction) (*entity.Transaction, error) {
	transaction.ID = t.setId()
	transaction.CreatedAt = t.setTime()
	transaction.UpdatedAt = t.setTime()

	_, err := t.client.Database("market").Collection("transaction").InsertOne(t.ctxMongo, transaction)
	if err != nil {
		return nil, err
	}
	result, err := t.FindTransaction(ctx, transaction.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t MongoTransactionRepository) FindTransaction(ctx context.Context, id string) (*entity.Transaction, error) {
	filter := bson.D{{"id", id}}
	transaction := &entity.Transaction{}
	result := t.client.Database("market").Collection("transaction").FindOne(t.ctxMongo, filter)
	err := result.Decode(&transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t MongoTransactionRepository) setTime() time.Time {
	return time.Now().UTC()
}

func (t MongoTransactionRepository) setId() string {
	return uuid.New().String()
}
