package routes

import (
	"github.com/Eng21072546/API_maketing/controller"
	"github.com/gofiber/fiber/v2"
)

func OrderRoute(app *fiber.App, orderHandler *controller.HttpOrderHandler) {
	app.Post("/order", orderHandler.CreateOrder)
	app.Patch("/order/status/:id", orderHandler.PatchOrderStatus)
}
