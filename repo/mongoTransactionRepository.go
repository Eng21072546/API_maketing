package repo

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

func (t MongoTransactionRepository) InsertTransaction(ctx context.Context, transaction *collection.Transaction) (*mongo.InsertOneResult, error) {
	transaction.ID = t.SetId()
	transaction.CreatedAt = t.SetTime()
	transaction.UpdatedAt = t.SetTime()
	return t.client.Database("market").Collection("transaction").InsertOne(t.ctxMongo, transaction)
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

func (t MongoTransactionRepository) SetTime() time.Time {
	return time.Now().UTC()
}

func (t MongoTransactionRepository) SetId() string {
	return uuid.New().String()
}
