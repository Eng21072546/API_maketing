package repo

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/useCase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOrderRepository struct {
	collection *mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewMongoOrderRepository(collection *mongo.Collection, ctx2 context.Context, cancelFunc context.CancelFunc) useCase.OrderRepository {
	return &MongoOrderRepository{
		collection: collection,
		ctx:        ctx2,
		cancel:     cancelFunc}
}

func (m *MongoOrderRepository) InsertOrder(order entity.Order) (*mongo.InsertOneResult, error) {

	//err := entity.CheckAddress(order)
	//if err != nil {
	//	errList = append(errList, err.Error())
	//}
	// Generate a random order ID (replace with a more robust ID generation mechanism if needed)
	//id will in range 10000-100000

	// Validate product availability in future (implementation not shown here)
	//for _, productorder := range order.ProductList {
	//	bool, err := CheckStock(productorder.ProductID, productorder.Quantity)
	//	if bool != true {
	//		fmt.Println(err)
	//		errorMessage := err.Error()             // Convert err to string
	//		errList = append(errList, errorMessage) // Append string to errList
	//	}
	//}

	//collection := client.Database("market").Collection("order")
	result, err := m.collection.InsertOne(m.ctx, order)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MongoOrderRepository) FindOrderById(id string) (*entity.Order, error) {
	result := m.collection.FindOne(m.ctx, bson.M{"id": id})
	order := new(entity.Order)
	err := result.Decode(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (m *MongoOrderRepository) UpdateOrderStatus(orderID string, newStatus entity.Status) (err error) {
	//Build the filter to identify the order
	filter := bson.M{"id": bson.M{"$eq": orderID}} // Replace "_id" if your order uses a different identifier

	// Update document with the new status
	update := bson.M{"$set": bson.M{"Status": newStatus}}

	// Update the order status
	_, err = m.collection.UpdateOne(m.ctx, filter, update)
	if err != nil {
		return err // Handle errors appropriately (e.g., logging, returning specific error codes)
		// }

	}
	return nil
}
