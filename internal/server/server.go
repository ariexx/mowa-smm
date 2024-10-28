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
			ServerHeader:  "mowa-backend",
			AppName:       "mowa-backend",
			StrictRouting: true,
		}),

		db: database.New(),
	}

	return server
}

func (f *FiberServer) RegisterMiddlewares() {
	f.App.Use(f.compress(), f.cors(), f.helmet(), f.idempotency(), f.limiter(), f.logger(), f.requestid(), f.recover())
}
