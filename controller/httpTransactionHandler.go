package controller

import (
	"errors"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/payload"
	"github.com/Eng21072546/API_maketing/useCase"
	"github.com/gofiber/fiber/v2"
)

type HttpTransactionHandler struct {
	transactionUseCase useCase.TransactionUseCase
}

func NewHttpTransactionHandler(transactionUseCase useCase.TransactionUseCase) *HttpTransactionHandler {
	return &HttpTransactionHandler{transactionUseCase: transactionUseCase}
}

func (t HttpTransactionHandler) PostTransaction(c *fiber.Ctx) error {
	transPayload := new(payload.Transaction)
	err := c.BodyParser(transPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.New("Invalid payload")})
	}
	transaction := entity.NewTransaction(transPayload.Address, transPayload.ProductOrder)
	transaction, errList := t.transactionUseCase.NewTransaction(c.Context(), transaction)
	if errList != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errorsToStrings(errList)})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"transaction": transaction})
}
