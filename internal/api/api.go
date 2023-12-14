package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hra42/Go-API/internal/api/ip"
	"github.com/hra42/Go-API/internal/api/time"
)

func StartServer() {
	app := fiber.New(fiber.Config{
		ServerHeader: "X-Forwarded-For",
	})

	app.Get("/", handleIndex)

	// time
	app.Get("/api/v1/time/current", time.HandleCurrent)

	// IPs
	app.Get("/api/v1/ip/current", ip.HandleCurrentIP)

	app.Listen(":3000")
}
