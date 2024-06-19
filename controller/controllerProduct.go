package controller

import (
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"time"
	//"context"
	"fmt"
	"github.com/Eng21072546/API_maketing/repo"
	"github.com/gofiber/fiber/v2"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	//"time"
)

func GetProducts(c *fiber.Ctx) error {

	productsList, err := repo.GetAllProduct()
	if err != nil {
		// Handle error (e.g., return error repo)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	fmt.Println("Get all products")
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": productsList})
}

func GetaProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id int
	fmt.Sscan(idStr, &id) // Convert string ID to int
	product, err := repo.GetProduct(id)
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
	bodyBytes := c.Body()
	if len(bodyBytes) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "empty request body"})
	}

	// 1. Parse the request body
	var product entity.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	rand.Seed(time.Now().UnixNano()) // random id product
	product.ID = rand.Intn(100000)

	result, err := repo.CreateProduct(product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// 4. Send a successful repo with the created product (optional)
	fmt.Println("Create a product")
	return c.Status(http.StatusCreated).JSON(result)
}

func PutProduct(c *fiber.Ctx) error {

	bodyBytes := c.Body()

	if len(bodyBytes) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "empty request body"})
	}

	productId := c.Params("id")
	var id int
	fmt.Sscan(productId, &id) // Convert string ID to int

	var productUpdates entity.ProductUpdate
	err := c.BodyParser(&productUpdates)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	validate := validator.New()
	err = validate.Struct(productUpdates)
	if err != nil {
		// Handle validation errors
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()}) // Provide specific error details
	}

	updateResult, err := repo.UpdateProduct(id, productUpdates)

	if err != nil {
		// Handle specific errors like "not found" differently
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// Check for update success (modified count)
	if updateResult.ModifiedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "product not found Or it doesn't change"})
	}
	fmt.Println("Update a product")
	return c.Status(http.StatusOK).JSON(updateResult)
}

func DeleteProduct(c *fiber.Ctx) error {

	productId := c.Params("id")
	var id int
	fmt.Sscan(productId, &id) // Convert string ID to int

	deleteResult, err := repo.DeleteProduct(id)
	if err != nil {
		// Handle specific errors like "not found" differently
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if deleteResult.DeletedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}
	fmt.Println("Delete a product")
	return c.Status(http.StatusOK).JSON(fiber.Map{"deleted": deleteResult}) // No content (204)
}

func PatchStock(c *fiber.Ctx) error {
	bodyBytes := c.Body()
	if len(bodyBytes) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "empty request body"})
	}
	productId := c.Params("id")
	var id int
	fmt.Sscan(productId, &id)
	var product entity.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	err = repo.UpdateStock(id, product.Stock)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	product, err = repo.GetProduct(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(product)
}
