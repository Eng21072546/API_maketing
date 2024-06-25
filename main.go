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

	productRepo := repo.NewMongoProductRepository(Client.Database("market").Collection("product"), Ctx, Cancel)
	productUseCase := useCase.NewProductUseCase(productRepo)
	productHandler := controller.NewHttpProductHandler(productUseCase)

	orderRepo := repo.NewMongoOrderRepository(Client.Database("market").Collection("order"), Ctx, Cancel)
	//orderUseCase := useCase.NewOrderUseCase(orderRepo, productRepo)
	//orderHandler := controller.NewHttpOrderHandler(orderUseCase)

	transactionRepo := repo.NewMongoTransactionRepository(Client, Ctx)
	transactionUseCase := useCase.NewTransactionUseCase(transactionRepo, productRepo, orderRepo)
	transactionHandler := controller.NewHttpTransactionHandler(transactionUseCase)
	app := fiber.New()
	app.Post("/order/calculate", transactionHandler.PostTransaction)
	app.Get("/product", productHandler.GetAllProducts)
	app.Get("product/:id", productHandler.GetProductById)
	app.Post("/product", productHandler.CreateProduct)
	app.Put("/product/:id", productHandler.UpdateProduct)
	app.Delete("/product/:id", productHandler.DeleteProduct)
	//app.Post("/order/:id", orderHandler.CreateOrder)
	//app.Patch("/order/status/:id", orderHandler.PatchOrderStatus)
	err := app.Listen(":6000")
	if err != nil {
		panic(err)
	}
	//repo.Init()
	//routes.UserRoute()
	configs.Close(configs.Client, configs.Ctx, configs.Cancel)
}
