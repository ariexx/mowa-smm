package server

import (
	"github.com/gofiber/fiber/v2"
)

func (f *FiberServer) RegisterFiberRoutes() {
	f.App.Get("/", f.HelloWorldHandler)

	f.App.Get("/health", f.healthHandler)

}

func (f *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}
	return c.JSON(resp)
}

func (f *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(f.db.Health())
}
