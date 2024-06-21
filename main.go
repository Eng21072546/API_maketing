package main

import (
	"github.com/Eng21072546/API_maketing/configs"
	"github.com/Eng21072546/API_maketing/controller"
	"github.com/Eng21072546/API_maketing/repo"
	"github.com/gofiber/fiber/v2"

	//"github.com/Eng21072546/API_maketing/routes"
	"github.com/Eng21072546/API_maketing/useCase"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load(".env")
	Client, Ctx, Cancel, Err := configs.Connect(os.Getenv("DB_URI"))
	if Err != nil {
		panic(Err)
	}
	orderRepo := repo.NewMongoOrderRepository(Client.Database("market").Collection("order"), Ctx, Cancel)
	orderUsecase := useCase.NewOrderUseCase(orderRepo)
	orderHandler := controller.NewHttpOrderHandler(orderUsecase)

	productRepo := repo.NewMongoProductRepository(Client.Database("market").Collection("product"), Ctx, Cancel)
	productUseCase := useCase.NewProductUseCase(productRepo)
	productHandler := controller.NewHttpProductHandler(productUseCase)
	app := fiber.New()
	app.Post("/order", orderHandler.CreateOrder)
	app.Get("/product", productHandler.GetAllProducts)
	app.Get("product/:id", productHandler.GetProductById)
	app.Post("/product", productHandler.CreateProduct)
	app.Put("/product/:id", productHandler.UpdateProduct)
	app.Delete("/product/:id", productHandler.DeleteProduct)
	app.Listen(":6000")
	//repo.Init()
	//routes.UserRoute()
	configs.Close(configs.Client, configs.Ctx, configs.Cancel)
}
