package controller

import (
	"github.com/Eng21072546/API_maketing/configs"
	"github.com/Eng21072546/API_maketing/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateOrder(c *fiber.Ctx) error {
	// 1. Parse the request body
	var order models.Order
	err := c.BodyParser(&order)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	// 2. Connect to MongoDB (assuming you have a Connect function defined elsewhere)
	client, ctx, cancel, err := configs.Connect("mongodb://localhost:27017")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer cancel()
	defer client.Disconnect(ctx)
	//for productId,quntity := range order.ProductList{
	//	response.CheckStock(ctx,client.Database("market").Collection("product"),productId,quntity)
	//}

	collection := client.Database("market").Collection("order")
	// 3. Insert the product into the database

	_, err = collection.InsertOne(ctx, order)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// 4. Send a successful response with the created product (optional)
	return c.Status(http.StatusCreated).JSON(fiber.Map{"order": order, "data": "orderconfirm"})

}
