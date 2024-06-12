package controller

import (
	"fmt"
	"github.com/Eng21072546/API_maketing/models"
	"github.com/Eng21072546/API_maketing/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"net/http"
	"time"
)

var order models.Order

func CreateOrder(c *fiber.Ctx) error {
	var errList []string
	// Extract user name (assuming it's sent in the request body)
	if err := c.BodyParser(&order); err != nil {
		return err // Handle decoding errors
	}
	collection := client.Database("market").Collection("order")
	// Alternatively, extract user name from authentication context (if using oIDC)
	// user := c.Locals("user").(*oidc.UserInfo) // Assuming user info is stored in "user" context key
	// order.CustomerName = user.PreferredUsername

	// Generate a random order ID (replace with a more robust ID generation mechanism if needed)
	rand.Seed(time.Now().UnixNano())
	order.ID = rand.Intn(100000) // Example format

	// Validate product availability in future (implementation not shown here)
	for _, productorder := range order.ProductList {
		bool, err := response.CheckStock(ctx, client.Database("market").Collection("product"), productorder.ProductID, productorder.Quantity)
		if bool != true {
			fmt.Println(err)
			errorMessage := err.Error()             // Convert err to string
			errList = append(errList, errorMessage) // Append string to errList
		}
	}
	if len(errList) != 0 {
		//fmt.Println(len(errList))
		fmt.Println("Order ID %d NOT confrim", order.ID)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": errList}, "orderNOTconfirm")
	} else {

		var status = models.New
		// Set order status (optional)
		order.Status = status //Enum
		//fmt.Println("Order Status %s", status)

		// Insert order into MongoDB
		_, err = collection.InsertOne(ctx, order)
		if err != nil {
			return err // Handle insertion errors
		}
		fmt.Println("Order ID %d confrim", order.ID)
		return c.Status(http.StatusCreated).JSON(fiber.Map{"order": order}, "orderconfirm")
	}
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
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "order not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"order": order}, "order request")
}
