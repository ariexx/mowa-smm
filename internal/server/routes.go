package server

import (
	"mowa-backend/injector"

	"github.com/gofiber/fiber/v2"
)

func (f *FiberServer) RegisterFiberRoutes() {
	f.App.Get("/", f.DefaultHandler)

	f.App.Get("/health", f.healthHandler)

	api := f.App.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			admin := v1.Group("/admin")
			{
				adminController := injector.InitializeAdminController()
				adminController.Route(admin)
			}
			users := v1.Group("/users")
			{
				userController := injector.InitializeUserController()
				userController.Route(users)
			}
		}
	}
}

func (f *FiberServer) DefaultHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Welcome",
		"version": "1.0.0",
		"status":  "up",
	}
	return c.JSON(resp)
}

func (f *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(f.db.Health())
}
