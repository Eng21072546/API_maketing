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
	"os"
)

type HttpTransactionHandler struct {
	transactionUseCase _interface.TransactionUseCase
}

func (t *HttpTransactionHandler) newTransactionHandlerLog() (*zap.Logger, func()) {
	logFactory := logger.NewLoggerFactory("./logger/logs")
	logger, loggerClose, err := logFactory.NewLogger()
	if err != nil {
		fmt.Printf("%v", err.Error())
		os.Exit(1)
	}
	//defer loggerClose()
	return logger, loggerClose
}

func NewHttpTransactionHandler(transactionUseCase _interface.TransactionUseCase) *HttpTransactionHandler {
	return &HttpTransactionHandler{transactionUseCase: transactionUseCase}
}

func (t *HttpTransactionHandler) PostTransaction(c *fiber.Ctx) error {
	log, logCancel := t.newTransactionHandlerLog()
	defer logCancel()
	transPayload := new(payload.Transaction)
	err := c.BodyParser(transPayload)
	if err != nil || transPayload.Address == "" || transPayload.ProductOrder == nil {
		log.Error("Invalid transaction req", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.New("invalid request body must have address and product order")})
	}
	transaction := entity.NewTransaction(transPayload.Address, transPayload.ProductOrder)
	transaction, errList := t.transactionUseCase.NewTransaction(c.Context(), transaction)
	if errList != nil && len(errList) != 0 {
		log.Error("Transaction creation error", zap.Error(errList[0]))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errorsToStrings(errList)})
	}
	log.Info("transaction created", zap.String("transaction_id", transaction.ID))
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"transaction": transaction})
}
