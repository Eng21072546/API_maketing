package response

import (
	"context"
	"fmt"
	"github.com/Eng21072546/API_maketing/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"time"
)

func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	collection := client.Database(dataBase).Collection(col)

	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func InsertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {

	collection := client.Database(dataBase).Collection(col)

	result, err := collection.InsertMany(ctx, docs)
	return result, err
}

func Query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

	collection := client.Database(dataBase).Collection(col)

	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
	return
}

func testInsert() {

	// get Client, Context, CancelFunc and err from connect method.
	client, ctx, _, err := configs.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected")
	}

	// Release resource when main function is returned.
	// defer close(client,ctx)
	// Create a object of type interface to  store
	// the bson values, that  we are inserting into database.
	var document interface{}

	document = bson.D{
		{"rollNo", 175},
		{"maths", 80},
		{"science", 90},
		{"computer", 95},
	}

	// insertOne accepts client , context, database
	// name collection name and an interface that
	// will be inserted into the  collection.
	// insertOne returns an error and a result of
	// insert in a single document into the collection.
	insertOneResult, err := InsertOne(client, ctx, "gfg",
		"marks", document)

	// handle the error
	if err != nil {
		panic(err)
	}

	// print the insertion id of the document,
	// if it is inserted.
	fmt.Println("Result of InsertOne")
	fmt.Println(insertOneResult.InsertedID)

	// Now will be inserting multiple documents into
	// the collection. create  a object of type slice
	// of interface to store multiple  documents
	var documents []interface{}

	// Storing into interface list.
	documents = []interface{}{
		bson.D{
			{"rollNo", 153},
			{"maths", 65},
			{"science", 59},
			{"computer", 55},
		},
		bson.D{
			{"rollNo", 162},
			{"maths", 86},
			{"science", 80},
			{"computer", 69},
		},
	}

	// insertMany insert a list of documents into
	// the collection. insertMany accepts client,
	// context, database name collection name
	// and slice of interface. returns error
	// if any and result of multi document insertion.
	insertManyResult, err := InsertMany(client, ctx, "market",
		"product", documents)

	// handle the error
	if err != nil {
		panic(err)
	}

	fmt.Println("Result of InsertMany")

	// print the insertion ids of the multiple
	// documents, if they are inserted.
	for id := range insertManyResult.InsertedIDs {
		fmt.Println(id)
	}
}

func Queryall() ([]map[string]interface{}, error) {
	// Get Client, Context, CancelFunc and err from connect method.
	client, ctx, _, err := configs.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	var filter, option interface{}

	filter = bson.D{
		// {"maths", bson.D{{"$gt", "*"}}},
		{},
	}
	cursor, err := Query(client, ctx, "market", "product", filter, option)
	// handle the errors.
	if err != nil {
		panic(err)
	}
	var results []bson.D

	// to get bson object  from cursor,
	// returns error if any.
	if err := cursor.All(ctx, &results); err != nil {
		// handle the error
		panic(err)
	}
	var products []map[string]interface{}
	//for cursor.Next(ctx) {
	//	var result map[string]interface{}
	//	err := cursor.Decode(&result) // Decode document into a map
	//	if err != nil {
	//		return nil, fmt.Errorf("failed to find products: %w", err)
	//	}
	//	products = append(products, result)
	//}
	// printing the result of query.
	fmt.Println("Query Result")
	products = convertBSONToMaps(results)
	defer cursor.Close(ctx) // Close the cursor after the function finishes

	return products, nil
}

func convertBSONToMaps(documents []bson.D) []map[string]interface{} {
	var results []map[string]interface{}
	for _, doc := range documents {
		result := make(map[string]interface{})
		for _, elem := range doc {
			result[elem.Key] = elem.Value
		}
		results = append(results, result)
	}
	return results
}

func CheckStock(ctx context.Context, collection *mongo.Collection, productID int, quantity int) (bool, error) {
	//client, ctx, _, err := configs.Connect("mongodb://localhost:27017")
	//if err != nil {
	//	panic(err)
	//}
	var result struct {
		DesiredField int `bson:"stock"` // Replace with actual field name and type
	}
	var filter interface{}
	// 1. Build the filter to find the product by ID
	filter = bson.M{"id": productID}

	// 2. Project only the "stock" field (optional, improve performance)
	//projection := bson.D{{"project", bson.M{"id": 0, "stock": 1}}} // Replace "desiredField" with your actual field name
	// Use FindOne instead of Find

	err := collection.FindOne(ctx, filter).Decode(&result)

	// Access the retrieved value in the "result" map using "desiredField" key
	//Query(client, ctx, "market", "product", filter, projection)
	//collection.FindOne(ctx, filter, options.Find().SetProjection(&projection)).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, fmt.Errorf("product with ID %d not found", productID)
		}
		return false, fmt.Errorf("error finding product: %w", err)
	}
	// 4. Check if stock is available and sufficient
	stock := result.DesiredField
	//if !ok {
	//	return false, fmt.Errorf("invalid stock value for product ID %d", productID)
	//}
	if stock < quantity {
		return false, fmt.Errorf("insufficient stock for product ID %d, only %d available", productID, stock)
	}

	// 5. If all good, return true and nil error
	return true, nil
}
