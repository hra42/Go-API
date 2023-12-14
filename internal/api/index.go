package api

import "github.com/gofiber/fiber/v2"

func handleIndex(c *fiber.Ctx) error {
	return c.Status(404).SendString("404 - No Endpoint specified")
}
