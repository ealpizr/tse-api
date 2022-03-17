package main

import (
	"github.com/ealpizr/tse-api/src/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.RegisterRoutes(app)
	app.Listen(":80")
}
