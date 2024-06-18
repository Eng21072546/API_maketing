package controller

import (
	"fmt"
	"github.com/Eng21072546/API_maketing/models"
	"github.com/Eng21072546/API_maketing/repo"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var order models.Order

func CreateOrder(c *fiber.Ctx) error {

	if err := c.BodyParser(&order); err != nil {
		return err // Handle decoding errors
	}
	order, err := repo.CreateOrder(order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Println("Order ID %d confrim", order.ID)
	return c.Status(http.StatusCreated).JSON(fiber.Map{"order": order}, "orderconfirm")
}

func GetOrder(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id int
	fmt.Sscan(idStr, &id) // Convert string ID to int
	order, err := repo.GetOrder(id)
	if err != nil {
		// Handle "not found" error differently
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "order not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"order": order}, "order request")
}

func UpdateStatus(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id int
	fmt.Sscan(idStr, &id)
	order, err := repo.GetOrder(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "order not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	currStatus := order.Status
	var newStatus models.Status
	if currStatus == models.New {
		err := repo.DecreaseStock(order) // Decrease stock when
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		newStatus = models.Paid
	} else if currStatus == models.Paid {
		newStatus = models.Processing
	} else if currStatus == models.Processing {
		newStatus = models.Done
	} else {
		newStatus = models.Done
	}
	err = repo.PatchOrderStatus(id, newStatus) //update status
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	order, err = repo.GetOrder(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "order not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	fmt.Println("Order ID %d confrim ", order.ID, " Status --> ", order.Status)

	return c.Status(http.StatusOK).JSON(fiber.Map{"order": order})
}

func GetOrderPrice(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id int
	fmt.Sscan(idStr, &id)
	order, err := repo.GetOrder(id)
	if err != nil {
		// Handle "not found" error differently
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "order not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"total price": repo.CalculateOrderPrice(order)})
}
