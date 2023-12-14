package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hra42/Go-API/internal/api/ip"
	"github.com/hra42/Go-API/internal/api/time"
)

func StartServer() {
	app := fiber.New()

	app.Get("/", handleIndex)

	// time
	app.Get("/api/v1/time/current", time.HandleCurrent)

	// IPs
	app.Get("/api/v1/ip/current", ip.HandleCurrentIP)
	app.Get("/api/v1/ip/current_v4", ip.HandleCurrentIPv4)
	app.Get("/api/v1/ip/current_v6", ip.HandleCurrentIPv6)

	app.Listen(":3000")
}
