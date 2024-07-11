package controller

import (
	"errors"
	"fmt"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/logger"
	"github.com/Eng21072546/API_maketing/payload"
	"github.com/Eng21072546/API_maketing/useCase/interface"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func (h HttpOrderHandler) newOrderHandlerLog() (*zap.Logger, func()) {
	logFactory := logger.NewLoggerFactory("./logger/logs")
	logger, loggerClose, err := logFactory.NewLogger()
	if err != nil {
		fmt.Printf("%v", err.Error())
		os.Exit(1)
	}
	//defer loggerClose()
	return logger, loggerClose
}

type HttpOrderHandler struct {
	orderUseCase _interface.OrderUseCase
}

func NewHttpOrderHandler(userUseCase _interface.OrderUseCase) *HttpOrderHandler {
	return &HttpOrderHandler{orderUseCase: userUseCase}
}

func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	logger, loggerClose := h.newOrderHandlerLog()
	defer loggerClose()

	var orderReq payload.Order
	if err := c.BodyParser(&orderReq); err != nil || orderReq.TransactionId == "" || orderReq.CustomerName == "" {
		logger.Warn("Invalid user_name or transaction_id", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": errors.New("Invalid request body mush have customerName and transaction")})
	}
	orderEntity := entity.Order{CustomerName: orderReq.CustomerName, TransactionId: orderReq.TransactionId}

	order, err := h.orderUseCase.NewOrder(c.Context(), &orderEntity)
	if len(err) != 0 {
		logger.Error("Cannot CreateNewOrder ")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": errorsToStrings(err)})
	}
	logger.Info("CreateNewOrder success", zap.Any("order_id", order.ID))
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"OrderCreate": order})

}

func (h *HttpOrderHandler) PatchOrderStatus(c *fiber.Ctx) error {
	logger, loggerClose := h.newOrderHandlerLog()
	defer loggerClose()
	id := c.Params("id")
	result, err := h.orderUseCase.PatchOrderStatus(c.Context(), id)
	if err != nil {
		logger.Warn("Cannot PatchOrderStatus", zap.Any("id", id), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	logger.Info("PatchOrderStatus success", zap.Any("order_id", result.ID))
	return c.Status(http.StatusOK).JSON(fiber.Map{"order": result})
}
