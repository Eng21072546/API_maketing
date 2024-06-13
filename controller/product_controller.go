package controller

import (
	"context"
	"github.com/Eng21072546/API_maketing/models"
	"go.mongodb.org/mongo-driver/mongo"
	//"context"
	"fmt"
	"github.com/Eng21072546/API_maketing/configs"
	"github.com/Eng21072546/API_maketing/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	//"time"
)

var ctx context.Context
var cancel context.CancelFunc
var client *mongo.Client
var err error

func Init() {
	fmt.Printf("Init product controller varable")
	ctx = configs.Ctx
	cancel = configs.Cancel
	client = configs.Client
	err = configs.Err
	if err != nil {
		fmt.Println(err)
	}
	//defer cancel()
	//defer client.Disconnect(ctx)
}

func GetallProducts(c *fiber.Ctx) error {
	fmt.Printf("Get all products")
	products, err := response.Queryall()
	if err != nil {
		// Handle error (e.g., return error response)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": products})
}

func GetaProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id int
	fmt.Sscan(idStr, &id) // Convert string ID to int

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	product, err := response.GetProduct(id)
	if err != nil { // Handle "not found" error differently
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	fmt.Println("Get a product")
	return c.Status(http.StatusOK).JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	fmt.Printf("Create a product")
	// 1. Parse the request body
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	//// 2. Connect to MongoDB (assuming you have a Connect function defined elsewhere)
	//client, ctx, cancel, err := configs.Connect("mongodb://localhost:27017")
	//if err != nil {
	//	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	//}
	//defer cancel()
	//defer client.Disconnect(ctx)

	// 3. Insert the product into the database
	collection := client.Database("market").Collection("product")
	_, err = collection.InsertOne(ctx, product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// 4. Send a successful response with the created product (optional)
	return c.Status(http.StatusCreated).JSON(product)
}
func UpdateProduct(c *fiber.Ctx) error {
	fmt.Printf("Update a product")
	// 1. Get product ID from the request path (adjust based on your API design)
	productId := c.Params("id")
	var id int
	fmt.Sscan(productId, &id) // Convert string ID to int

	// 2. Parse the request body for updates
	var productUpdates map[string]interface{}
	err := c.BodyParser(&productUpdates)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	//// 3. Connect to MongoDB (assuming you have a Connect function defined elsewhere)
	//client, ctx, cancel, err := configs.Connect("mongodb://localhost:27017")
	//if err != nil {
	//	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	//}
	//defer cancel()
	//defer client.Disconnect(ctx)

	// 4. Build the update filter and document
	filter := bson.M{"id": id}
	update := bson.D{{"$set", productUpdates}} // Update specific fields

	// 5. Update the product in the database
	collection := client.Database("market").Collection("product")
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		// Handle specific errors like "not found" differently
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// 6. Check for update success (modified count)
	if updateResult.ModifiedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}

	// 7. Optional: Fetch the updated product (consider performance implications)
	var updatedProduct models.Product
	err = collection.FindOne(ctx, filter).Decode(&updatedProduct)
	if err != nil {
		// Handle potential error here (consider logging)
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "product updated"}) // Send basic success response
	}

	// 8. Send a successful response with the updated product (optional)
	return c.Status(http.StatusOK).JSON(updatedProduct)
}
func DeleteProduct(c *fiber.Ctx) error {
	fmt.Printf("Delete a product")
	// 1. Get product ID from the request path (adjust based on your API design)
	productId := c.Params("id")
	var id int
	fmt.Sscan(productId, &id) // Convert string ID to int

	// 2. Connect to MongoDB (assuming you have a Connect function defined elsewhere)
	//client, ctx, cancel, err := configs.Connect("mongodb://localhost:27017")
	//if err != nil {
	//	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	//}
	//defer cancel()
	//defer client.Disconnect(ctx)

	// 3. Build the delete filter
	filter := bson.M{"id": id}

	// 4. Delete the product from the database
	collection := client.Database("market").Collection("product")
	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		// Handle specific errors like "not found" differently
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// 5. Check for delete success (deleted count)
	if deleteResult.DeletedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}
	// 6. Send a successful response (consider returning a success message or no content)
	return c.Status(http.StatusNoContent).JSON(fiber.Map{"Delete": productId}) // No content (204)
}

func PatchStock(c *fiber.Ctx) error {
	productId := c.Params("id")
	var id int
	fmt.Sscan(productId, &id)
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	err = response.UpdateStock(id, product.Stock)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	product, err = response.GetProduct(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(product)
}
