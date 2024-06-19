package repo

import (
	"context"
	"fmt"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateProduct(product entity.Product) (*mongo.InsertOneResult, error) {
	collection := client.Database("market").Collection("product")
	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetProduct(productID int) (entity.Product, error) {
	var product entity.Product
	filter := bson.M{"id": productID}
	collection := client.Database("market").Collection("product")
	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func GetAllProduct() (products []entity.Product, err error) {
	var productsList []entity.Product
	collection := client.Database("market").Collection("product")
	cursor, err := collection.Find(context.Background(), bson.M{}) // Empty bson.M{} matches all documents
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	err = cursor.All(context.Background(), &productsList)
	if err != nil {
		return nil, err
	}
	return productsList, nil
}

func UpdateProduct(id int, productUpdates entity.ProductUpdate) (*mongo.UpdateResult, error) {

	filter := bson.M{"id": id}
	update := bson.D{{"$set", productUpdates}} // Update specific fields
	collection := client.Database("market").Collection("product")
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return updateResult, err
	}
	return updateResult, nil
}

func DeleteProduct(id int) (*mongo.DeleteResult, error) {

	filter := bson.M{"id": id}
	collection := client.Database("market").Collection("product")
	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return deleteResult, err
	}
	return deleteResult, nil
}

func UpdateStock(productID int, quantity int) error {
	collection := client.Database("market").Collection("product")
	filter := bson.M{"id": productID}
	update := bson.M{"$set": bson.M{"stock": quantity}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func CheckStock(productID int, quantity int) (bool, error) {

	var result struct {
		DesiredField int `bson:"stock"` // Replace with actual field name and type
	}
	var filter interface{}
	filter = bson.M{"id": productID}

	collection := client.Database("market").Collection("product")

	err := collection.FindOne(ctx, filter).Decode(&result) // Access the retrieved value in the "result" map using "desiredField" key

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, fmt.Errorf("product with ID %d not found", productID)
		}
		return false, fmt.Errorf("error finding product: %w", err)
	}
	// 4. Check if stock is available and sufficient
	stock := result.DesiredField

	if stock < quantity {
		return false, fmt.Errorf("insufficient stock for product ID %d, only %d available", productID, stock)
	}

	return true, nil
}
