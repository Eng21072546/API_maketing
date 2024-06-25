package controller

import (
	"errors"
	"github.com/Eng21072546/API_maketing/payload"
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

//func (h *HttpOrderHandler) CalculatePriceOrder(c *fiber.Ctx) error {
//	fmt.Println("CalculatePriceOrder")
//	var orderReq payload.Order
//	if err := c.BodyParser(&orderReq); err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(errors.New("error1"))
//	}
//	order := entity.Order{Address: orderReq.Address, ProductList: orderReq.ProductList}
//	order, err := h.orderUseCase.CalculateOrderPrice(order)
//	fmt.Println("CalculatePriceOrder", order, err)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(errors.New("error2"))
//	}
//	fmt.Println("Order ID %d confrim", order.ID)
//	return c.Status(http.StatusCreated).JSON(fiber.Map{"order": order}, "orderconfirm")
//}

func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	var orderReq payload.Order
	if err := c.BodyParser(&orderReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": errors.New("Invalid request body")})
	}
	orderEntity, err := h.orderUseCase.NewOrderEntity(c.Context(), orderReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": errorsToStrings(err)})
	}
	orderEntity, err = h.orderUseCase.NewOrder(c.Context(), orderEntity)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": errorsToStrings(err)})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"OrderCreate": orderEntity})
}

func (h *HttpOrderHandler) PatchOrderStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.orderUseCase.PatchOrderStatus(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"order": result})
}
