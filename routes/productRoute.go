package routes

import (
	"github.com/Eng21072546/API_maketing/controller"
	"github.com/gofiber/fiber/v2"
)

func ProductRoute(app *fiber.App, productHandler *controller.HttpProductHandler) {

	app.Get("/product", productHandler.GetAllProducts)
	app.Get("product/:id", productHandler.GetProductById)
	app.Post("/product", productHandler.CreateProduct)
	app.Put("/product/:id", productHandler.UpdateProduct)
	app.Delete("/product/:id", productHandler.DeleteProduct)

}
