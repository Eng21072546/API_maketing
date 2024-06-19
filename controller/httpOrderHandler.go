package controller

import (
	"fmt"
	"github.com/Eng21072546/API_maketing/useCase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type HttpOrderHandler struct {
	orderUseCase useCase.OrderUseCase
}

func NewHttpOrderHandler(userUseCase useCase.OrderUseCase) *HttpOrderHandler {
	return &HttpOrderHandler{orderUseCase: userUseCase}
}

func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	if err := c.BodyParser(&order); err != nil {
		return err // Handle decoding errors
	}
	order, err := h.orderUseCase.CreateOrder(order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Println("Order ID %d confrim", order.ID)
	return c.Status(http.StatusCreated).JSON(fiber.Map{"order": order}, "orderconfirm")
}
