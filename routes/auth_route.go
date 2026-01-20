package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wahyujatirestu/simple-procurement-system/controllers"
)

func AuthRoutes(r fiber.Router, ac *controllers.AuthController)  {
	auth := r.Group("/auth")

	auth.Post("/register", ac.Register)
	auth.Post("/login", ac.Login)
}