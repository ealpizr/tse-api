package main

import (
	"os"
	"github.com/ealpizr/tse-api/src/router"
	"github.com/gofiber/fiber/v2"
)

// @title		API Docs
// @BasePath 	/

func main() {
	port := os.Getenv("PORT")
	
	app := fiber.New()
	router.RegisterRoutes(app)
	app.Listen(":" + port)
}
