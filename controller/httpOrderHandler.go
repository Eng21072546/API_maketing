package controller

import (
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

//func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
//	id := c.Params("id")
//	transaction, err := h.orderUseCase.GetOrderTransaction(id)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(errors.New("transaction not found"))
//	}
//	order, err2 := h.orderUseCase.CreateOrder(transaction)
//	if err2 != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(err2)
//	}
//	return c.Status(http.StatusCreated).JSON(fiber.Map{"order": order})
//}

func (h *HttpOrderHandler) PatchOrderStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.orderUseCase.PatchOrderStatus(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"order": result})
}
