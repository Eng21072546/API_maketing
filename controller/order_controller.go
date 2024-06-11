package controller

import (
	"fmt"
	"github.com/Eng21072546/API_maketing/models"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"net/http"
	"time"
)

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

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
