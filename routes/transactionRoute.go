package routes

import (
	"github.com/Eng21072546/API_maketing/controller"
	"github.com/gofiber/fiber/v2"
)

func TransactionRoute(app *fiber.App, transactionHandler *controller.HttpTransactionHandler) {
	app.Post("/order/calculate", transactionHandler.PostTransaction)
}
