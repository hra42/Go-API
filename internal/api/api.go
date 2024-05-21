package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hra42/Go-API/internal/api/ip"
	"github.com/hra42/Go-API/internal/api/middleware"
	"github.com/hra42/Go-API/internal/api/time"
)

func StartServer() {
	log.Println("Starting server")
	app := fiber.New(fiber.Config{
		ServerHeader: "X-Forwarded-For",
	})

	middleware.Init()

	app.Get("/", handleIndex)

	// time
	app.Get("/api/v1/current/time",middleware.AuthRequired, time.HandleCurrent)

	// IPs
	app.Get("/api/v1/current/ip/", middleware.AuthRequired, ip.HandleCurrentIP)

	// start the server
	app.Listen(":3000")
}
