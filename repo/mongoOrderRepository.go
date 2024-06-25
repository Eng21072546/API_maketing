package repo

import (
	"context"
	"fmt"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/useCase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOrderRepository struct {
	client   *mongo.Client
	ctxMongo context.Context
}

func NewMongoOrderRepository(client *mongo.Client, ctx context.Context) useCase.OrderRepository {
	return &MongoOrderRepository{client, ctx}
}

func (m *MongoOrderRepository) InsertOrder(ctx context.Context, order collection.Order) (*mongo.InsertOneResult, error) {
	result, err := m.client.Database("market").Collection("order").InsertOne(m.ctxMongo, order)
	if err != nil {
		return nil, err
	}
	fmt.Println("Saved order", result)
	return result, nil
}

func (m *MongoOrderRepository) FindOrderById(ctx context.Context, id string) (*entity.Order, error) {
	result := m.client.Database("market").Collection("order").FindOne(m.ctxMongo, bson.M{"id": id})
	order := new(entity.Order)
	err := result.Decode(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (m *MongoOrderRepository) UpdateOrderStatus(ctx context.Context, orderID string, newStatus entity.Status) (err error) {
	//Build the filter to identify the order
	filter := bson.M{"id": bson.M{"$eq": orderID}} // Replace "_id" if your order uses a different identifier

	// Update document with the new status
	update := bson.M{"$set": bson.M{"Status": newStatus}}

	// Update the order status
	_, err = m.client.Database("market").Collection("order").UpdateOne(m.ctxMongo, filter, update)
	if err != nil {
		return err // Handle errors appropriately (e.g., logging, returning specific error codes)
		// }

	}
	return nil
}
