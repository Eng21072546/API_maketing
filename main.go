package main

import (
	"fmt"
	"github.com/Eng21072546/API_maketing/configs"
	"github.com/Eng21072546/API_maketing/controller"
	"github.com/Eng21072546/API_maketing/repo"
	"github.com/Eng21072546/API_maketing/routes"
	"github.com/gofiber/fiber/v2"
	//"github.com/Eng21072546/API_maketing/routes"
	"github.com/Eng21072546/API_maketing/useCase"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load(".env")
	fmt.Println(os.Getenv("MONGODB_URI"))
	Client, Ctx, Cancel, Err := configs.Connect(os.Getenv("MONGODB_URI"))
	if Err != nil {
		panic(Err)
	}

	transactionRepo := repo.NewMongoTransactionRepository(Client, Ctx)
	productRepo := repo.NewMongoProductRepository(Client, Ctx)
	orderRepo := repo.NewMongoOrderRepository(Client, Ctx)

	productUseCase := useCase.NewProductUseCase(productRepo)
	productHandler := controller.NewHttpProductHandler(productUseCase)

	orderUseCase := useCase.NewOrderUseCase(orderRepo, productRepo, transactionRepo)
	orderHandler := controller.NewHttpOrderHandler(orderUseCase)

	transactionUseCase := useCase.NewTransactionUseCase(transactionRepo, productRepo, orderRepo)
	transactionHandler := controller.NewHttpTransactionHandler(transactionUseCase)

	app := fiber.New()

	routes.ProductRoute(app, productHandler)
	routes.OrderRoute(app, orderHandler)
	routes.TransactionRoute(app, transactionHandler)

	err := app.Listen(":6000")

	//log.Debug("starting server at http://localhost:" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	//repo.Init()
	//routes.UserRoute()
	configs.Close(Client, Ctx, Cancel)
}
