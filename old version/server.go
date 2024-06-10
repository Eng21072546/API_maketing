package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// connectToDB connects to the MongoDB database and returns the collection handle
func connectToDB(ctx context.Context, uri string) (*mongo.Collection, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI())
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	return client.Database("Market").Collection("products"), nil
}

func getProduct_fiber(c *fiber.Ctx) error {

	return c.JSON(fiber.StatusOK, product)
}
func main() {
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("return /")
		return c.SendString("Hello, World!")
	})
	app.Get("/:ID", func(c *fiber.Ctx) error {
		fmt.Println("return ID")
		return c.SendString("ID " + c.Params("ID"))
	})
	app.Get("/api/*", func(c *fiber.Ctx) error {
		fmt.Println("return information")
		return c.SendString("API " + c.Params("*"))
	})
	app.Listen(":3000")
}
