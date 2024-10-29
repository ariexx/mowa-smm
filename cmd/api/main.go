package main

import (
	"context"
	"fmt"
	"log"
	"mowa-backend/internal/server"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func initializeServer() *server.FiberServer {
	initServer := server.New()
	initServer.RegisterMiddlewares()
	initServer.RegisterFiberRoutes()
	return initServer
}

func startServer(fiberServer *server.FiberServer) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := fiberServer.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}

func gracefulShutdown(fiberServer *server.FiberServer, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := fiberServer.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
	done <- true
}

func main() {
	initServer := initializeServer()

	done := make(chan bool, 1)

	go startServer(initServer)
	go gracefulShutdown(initServer, done)

	<-done
	log.Println("Graceful shutdown complete.")
}
