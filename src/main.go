package main

import (
	"os"
	"github.com/ealpizr/tse-api/src/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.GetEnv("PORT")
	
	app := fiber.New()
	router.RegisterRoutes(app)
	app.Listen(":80")
}
