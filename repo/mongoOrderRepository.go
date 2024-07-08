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

type MongoOrderRepository struct {
	client   *mongo.Client
	ctxMongo context.Context
}

func NewMongoOrderRepository(client *mongo.Client, ctx context.Context) OrderRepository {
	return &MongoOrderRepository{client, ctx}
}

func (m *MongoOrderRepository) InsertOrder(ctx context.Context, order collection.Order) (*entity.Order, error) {

	_, err := m.client.Database("market").Collection("order").InsertOne(m.ctxMongo, order)
	if err != nil {
		return nil, err
	}
	result, err := m.FindOrderById(ctx, order.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MongoOrderRepository) FindOrderById(ctx context.Context, id string) (*entity.Order, error) {
	filter := bson.D{{"id", id}}
	result := m.client.Database("market").Collection("order").FindOne(m.ctxMongo, filter)
	order := new(entity.Order)
	err := result.Decode(order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return order, nil
}

func (m *MongoOrderRepository) UpdateOrderStatus(ctx context.Context, orderID string, newStatus entity.Status) (*mongo.UpdateResult, error) {
	//Build the filter to identify the order
	filter := bson.M{"id": bson.M{"$eq": orderID}} // Replace "_id" if your order uses a different identifier

	// Update document with the new status
	update := bson.M{"$set": bson.M{"status": newStatus, "UpdatedAt": m.SetTime()}}

	// Update the order status
	result, err := m.client.Database("market").Collection("order").UpdateOne(m.ctxMongo, filter, update)

	if err != nil {
		return result, err // Handle errors appropriately (e.g., logging, returning specific error codes)
		// }

	}
	return result, nil
}

func (m *MongoOrderRepository) SetTime() time.Time {
	return time.Now().UTC()
}

func (m *MongoOrderRepository) SetId() string {
	return uuid.New().String()
}
