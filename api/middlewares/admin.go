package middlewares

import "github.com/gofiber/fiber/v2"

func IsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		if role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   "forbidden",
				"status":  fiber.StatusForbidden,
				"success": false,
			})
		} else {
			return c.Next()
		}
	}
}
