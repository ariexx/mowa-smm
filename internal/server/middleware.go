package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"time"
)

func (f *FiberServer) compress() fiber.Handler {
	return compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	})
}

func (f *FiberServer) cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*", //TODO: Change this to the actual origin
		AllowHeaders: "Origin, Content-Type, Accept",
	})
}

func (f *FiberServer) helmet() fiber.Handler {
	return helmet.New()
}

func (f *FiberServer) idempotency() fiber.Handler {
	return idempotency.New()
}

func (f *FiberServer) limiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max: 10,
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   "Too many requests",
				"status":  fiber.StatusTooManyRequests,
				"success": false,
			})
		},
		Expiration: 1 * time.Minute,
	})
}

func (f *FiberServer) logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${locals:requestid} - [${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	})
}

func (f *FiberServer) recover() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			c.JSON(fiber.Map{
				"success": false,
				"status":  fiber.StatusInternalServerError,
				"error":   e.(error).Error(),
			})
		},
	})
}

func (f *FiberServer) requestid() fiber.Handler {
	return requestid.New()
}
