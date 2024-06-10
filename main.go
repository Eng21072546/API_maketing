package main

import (
	"fmt"
	"github.com/Eng21072546/API_maketing/configs"
	"github.com/Eng21072546/API_maketing/controller"
	"github.com/Eng21072546/API_maketing/routes"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	fmt.Println("Listening on port:", port)
	configs.Connect(os.Getenv("DB_URI"))
	controller.Init()
	routes.UserRoute()
}
