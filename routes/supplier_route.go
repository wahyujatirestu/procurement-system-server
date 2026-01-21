package routes

import (
	"github.com/wahyujatirestu/simple-procurement-system/controllers"
	"github.com/wahyujatirestu/simple-procurement-system/middleware"

	"github.com/gofiber/fiber/v2"
)

func SupplierRoute(r fiber.Router, c *controllers.SupplierController, auth middleware.AuthMiddleware) {
	supplier := r.Group("/suppliers")
	supplier.Use(auth.RequireToken())
	supplier.Post("/", c.Create)
	supplier.Get("/", c.FindAll)
	supplier.Get("/:id", c.FindById)
	supplier.Put("/:id", c.Update)
	supplier.Delete("/:id", c.Delete)
}
