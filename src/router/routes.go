package router

import (
	"github.com/ealpizr/tse-api/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/id", handlers.FindByIDRequestHandler)
	app.Post("/name", handlers.FindByFullNameRequestHandler)
}
