package routes

import (
	"github.com/Eng21072546/API_maketing/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoute() {
	app := fiber.New()
	app.Get("/product", controller.GetallProducts)
	app.Post("/product", controller.CreateProduct)
	app.Get("/product/:id", controller.GetaProduct)
	app.Put("/product/:id", controller.UpdateProduct)
	app.Delete("/product/:id", controller.DeleteProduct)
	app.Post("/orders", controller.CreateOrder)
	//app.Get("/", func(c *fiber.Ctx) error { return c.SendString("hello") })
	app.Listen(":6000")

}
