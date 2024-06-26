package routes

import (
	"github.com/Eng21072546/API_maketing/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoute() {
	app := fiber.New()
	app.Get("/product", controller.GetProducts)
	app.Post("/product", controller.PostProduct)
	app.Get("/product/:id", controller.GetaProduct)
	app.Put("/product/:id", controller.PutProduct)
	app.Delete("/product/:id", controller.DeleteProduct)
	app.Patch("product/stock/:id", controller.PatchStock)
	app.Post("/order", controller.CreateOrder)
	app.Get("/order/:id", controller.GetOrder)
	app.Patch("/order/status/:id", controller.UpdateStatus)
	app.Post("/order/calculate/:id", controller.GetOrderPrice)

	//app.Get("/", func(c *fiber.Ctx) error { return c.SendString("hello") })
	app.Listen(":6000")

}
