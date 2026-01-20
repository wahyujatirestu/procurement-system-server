package routes

import (
	"github.com/wahyujatirestu/simple-procurement-system/controllers"
	"github.com/wahyujatirestu/simple-procurement-system/middleware"

	"github.com/gofiber/fiber/v2"
)

func ItemRoute(r fiber.Router, c *controllers.ItemController, auth middleware.AuthMiddleware) {
	item := r.Group("/items")
	item.Use(auth.RequireToken("admin"))
	item.Post("/", c.Create)
	item.Get("/", c.FindAll)
}
