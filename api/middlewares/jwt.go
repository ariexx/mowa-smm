package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap/zapcore"
	"mowa-backend/api/services"
)

func JwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwt := services.NewJwtService()
		logService := services.NewLoggerService()
		// Get the JWT token from the request header
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"status":  fiber.StatusUnauthorized,
				"success": false,
			})
		}

		// Validate the token
		isValid := jwt.ValidateToken(token)
		if !isValid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"status":  fiber.StatusUnauthorized,
				"success": false,
			})
		}

		// get the user from the token
		user, err := jwt.GetUserFromToken(token)
		if err != nil {
			logService.PrintStdout(c.Context(), zapcore.ErrorLevel, "JwtMiddleware.GetUserFromToken", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"status":  fiber.StatusUnauthorized,
				"success": false,
			})
		}

		// Set the user in the context
		c.Locals("user", user)

		// Continue stack
		return c.Next()
	}
}
