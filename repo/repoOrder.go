package repo

import (
	"fmt"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"strconv"
	"time"
	//"time"
)

func PatchOrderStatus(orderID int, newStatus entity.Status) (err error) {
	// Access the collection for orders
	collection := client.Database("market").Collection("order")

	// Build the filter to identify the order
	filter := bson.M{"id": bson.M{"$eq": orderID}} // Replace "_id" if your order uses a different identifier

	// Update document with the new status
	update := bson.M{"$set": bson.M{"status": newStatus}}

	// Update the order status
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err // Handle errors appropriately (e.g., logging, returning specific error codes)
	}

	return nil // Indicate successful update
}

func GetOrder(orderID int) (entity.Order, error) {
	var order entity.Order
	filter := bson.M{"id": orderID}
	collection := client.Database("market").Collection("order")
	err := collection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.Order{}, fmt.Errorf("order with ID %d not found", orderID)
		}
		return entity.Order{}, err
	}
	return order, nil
}
func CalculateOrderPrice(order entity.Order) float64 {
	var totalPrice float64
	var bill []string
	logisticPrice, _ := entity.LogisticCost(order)
	totalPrice += logisticPrice
	for _, productOrder := range order.ProductList {
		product, _ := GetProduct(productOrder.ProductID)
		productPrice := product.Price * float64(productOrder.Quantity)
		bill = append(bill, product.Name, " ", strconv.FormatFloat(product.Price, 'f', 2, 64), "  ", string(productOrder.Quantity), " ", strconv.FormatFloat(productPrice, 'f', 2, 64))
		totalPrice += productPrice
	}
	return totalPrice
}

func DecreaseStock(order entity.Order) error {
	for _, productOrder := range order.ProductList {
		product, _ := GetProduct(productOrder.ProductID)
		currenStock := product.Stock
		newStock := currenStock - productOrder.Quantity
		err := UpdateStock(productOrder.ProductID, newStock)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateOrder(order entity.Order) (entity.Order, []string) {
	var errList []string
	err := entity.CheckAddress(order)
	if err != nil {
		errList = append(errList, err.Error())
	}
	// Generate a random order ID (replace with a more robust ID generation mechanism if needed)
	rand.Seed(time.Now().UnixNano())
	order.ID = 10000 + rand.Intn(90001) //id will in range 10000-100000

	// Validate product availability in future (implementation not shown here)
	for _, productorder := range order.ProductList {
		bool, err := CheckStock(productorder.ProductID, productorder.Quantity)
		if bool != true {
			fmt.Println(err)
			errorMessage := err.Error()             // Convert err to string
			errList = append(errList, errorMessage) // Append string to errList
		}
	}
	if len(errList) != 0 {
		fmt.Println("Order ID %d NOT confrim", order.ID)
		return order, errList
	} else {
		var status = entity.New
		order.Status = status // Set order status  Enum
		collection := client.Database("market").Collection("order")
		_, err = collection.InsertOne(ctx, order)
		if err != nil {
			return order, errList // Handle insertion errors
		}
	}

	return order, nil
}
