package controller

import (
	"context"
	"github.com/Eng21072546/API_maketing/models"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"time"

	//"context"
	"fmt"
	"github.com/Eng21072546/API_maketing/configs"
	"github.com/Eng21072546/API_maketing/response"
	"github.com/gofiber/fiber/v2"
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

func PostProduct(c *fiber.Ctx) error {
	// 1. Parse the request body
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	rand.Seed(time.Now().UnixNano()) // random id product
	product.ID = rand.Intn(100000)

	result, err := response.CreateProduct(product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// 4. Send a successful response with the created product (optional)
	fmt.Println("Create a product")
	return c.Status(http.StatusCreated).JSON(result)
}

func PutProduct(c *fiber.Ctx) error {

	productId := c.Params("id")
	var id int
	fmt.Sscan(productId, &id) // Convert string ID to int

	var productUpdates models.ProductUpdate
	err := c.BodyParser(&productUpdates)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	updateResult, err := response.UpdateProduct(id, productUpdates)

	if err != nil {
		// Handle specific errors like "not found" differently
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// Check for update success (modified count)
	if updateResult.ModifiedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}
	fmt.Println("Update a product")
	return c.Status(http.StatusOK).JSON(updateResult)
}

func DeleteProduct(c *fiber.Ctx) error {

	productId := c.Params("id")
	var id int
	fmt.Sscan(productId, &id) // Convert string ID to int

	deleteResult, err := response.DeleteProduct(id)
	if err != nil {
		// Handle specific errors like "not found" differently
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if deleteResult.DeletedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}
	fmt.Println("Delete a product")
	return c.Status(http.StatusNoContent).JSON(fiber.Map{"deleted": productId}) // No content (204)
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
