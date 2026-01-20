package routes

import (
	"github.com/wahyujatirestu/simple-procurement-system/controllers"
	"github.com/wahyujatirestu/simple-procurement-system/middleware"

	"github.com/gofiber/fiber/v2"
)

func PurchasingRoute(r fiber.Router, c *controllers.PurchasingController, auth middleware.AuthMiddleware)  {
	purchase := r.Group("/purchases")
	purchase.Use(auth.RequireToken())
	purchase.Post("/", c.Create)
	purchase.Get("/", c.FindAll)
}