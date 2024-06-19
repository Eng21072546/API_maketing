package repo

import (
	"context"
	"fmt"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/useCase"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"time"
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

func (m *MongoOrderRepository) SaveOrder(order entity.Order) (entity.Order, []string) {
	var errList []string
	//err := entity.CheckAddress(order)
	//if err != nil {
	//	errList = append(errList, err.Error())
	//}
	// Generate a random order ID (replace with a more robust ID generation mechanism if needed)
	rand.Seed(time.Now().UnixNano())
	order.ID = 10000 + rand.Intn(90001) //id will in range 10000-100000

	// Validate product availability in future (implementation not shown here)
	//for _, productorder := range order.ProductList {
	//	bool, err := CheckStock(productorder.ProductID, productorder.Quantity)
	//	if bool != true {
	//		fmt.Println(err)
	//		errorMessage := err.Error()             // Convert err to string
	//		errList = append(errList, errorMessage) // Append string to errList
	//	}
	//}
	if len(errList) != 0 {
		fmt.Println("Order ID %d NOT confrim", order.ID)
		return order, errList
	} else {
		var status = entity.New
		order.Status = status // Set order status  Enum
		//collection := client.Database("market").Collection("order")
		_, err = m.collection.InsertOne(ctx, order)
		if err != nil {
			return order, errList // Handle insertion errors
		}
	}

	return order, nil
}
