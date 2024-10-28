package server

import (
	"github.com/gofiber/fiber/v2"

	"mowa-backend/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "mowa-backend",
			AppName:      "mowa-backend",
		}),

		db: database.New(),
	}

	return server
}
