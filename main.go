package main

import (
	"github.com/Eng21072546/API_maketing/configs"
	"github.com/Eng21072546/API_maketing/controller"
	"github.com/Eng21072546/API_maketing/response"
	"github.com/Eng21072546/API_maketing/routes"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load(".env")
	configs.Connect(os.Getenv("DB_URI"))
	controller.Init()
	response.Init()
	routes.UserRoute()
	configs.Close(configs.Client, configs.Ctx, configs.Cancel)
}
