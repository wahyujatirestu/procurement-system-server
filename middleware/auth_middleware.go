package middleware

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"github.com/wahyujatirestu/simple-procurement-system/utils/services"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) fiber.Handler
}

type authMiddleware struct {
	jwtService services.JwtService
}

func NewAuthMiddleware(jwtService services.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}

func (a *authMiddleware) RequireToken(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		claim, err := a.jwtService.VerifyToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		user := models.User{
			ID: claim.UserId,
			Role: claim.Role,
		}

		if len(roles) > 0 {
			valid := false
			for _, role := range roles {
				if role == user.Role {
					valid = true
					break
				}
			}

			if !valid {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "Forbidden",
				})
			}
		}
		c.Locals("user", user)
		return c.Next()
	}
}