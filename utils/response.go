package utils

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    data,
	})		
}

func Error(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": message,
	})
}