package main

import (
	"fmt"
	"stonks-api/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting Stonks API...")
	godotenv.Load()
	app := fiber.New(fiber.Config{
		ReadBufferSize: 4096 * 3,
	})

	basePath := "/api"

	app.Get(fmt.Sprintf("%s/users/health", basePath), services.GetUserHealth)
	app.Get(fmt.Sprintf("%s/stonks/health", basePath), services.GetStocksHealth)
	app.Listen(":4545")
}
