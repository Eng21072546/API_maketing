package controller

import (
	"fmt"
	"github.com/Eng21072546/API_maketing/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"net/http"
	"time"
)

var order models.Order

func CreateOrder(c *fiber.Ctx) error {

	// Extract user name (assuming it's sent in the request body)
	if err := c.BodyParser(&order); err != nil {
		return err // Handle decoding errors
	}

	// Alternatively, extract user name from authentication context (if using oIDC)
	// user := c.Locals("user").(*oidc.UserInfo) // Assuming user info is stored in "user" context key
	// order.CustomerName = user.PreferredUsername

	// Generate a random order ID (replace with a more robust ID generation mechanism if needed)
	rand.Seed(time.Now().UnixNano())
	order.ID = rand.Intn(100000) // Example format
	fmt.Println("Order ID %d confrim", order.ID)

	// Validate product availability in future (implementation not shown here)
	// ...

	// Set order status (optional)
	order.Status = "pending" // Example status

	// Insert order into MongoDB
	collection := client.Database("market").Collection("order") // Replace with your database and collection names
	_, err = collection.InsertOne(ctx, order)
	if err != nil {
		return err // Handle insertion errors
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"order": order}, "orderconfirm")
}

func GetOrder(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id int
	fmt.Sscan(idStr, &id) // Convert string ID to int
	filter := bson.M{"id": id}
	collection := client.Database("market").Collection("order")
	err := collection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		// Handle "not found" error differently
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"order": order}, "order request")
}
