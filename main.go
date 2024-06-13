package main

import (
	"os"

	"github.com/IIITManjeet/shortify/database"
	"github.com/IIITManjeet/shortify/handlers"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New(fiber.Config{})
	database.InitDB()
	handlers.InitRouter(app)
	app.Listen(os.Getenv("PORT"))
}
