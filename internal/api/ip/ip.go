package ip

import (
	"github.com/gofiber/fiber/v2"
)

func HandleCurrentIP(c *fiber.Ctx) error {
	ip := c.IPs()[0]
	return c.JSON(fiber.Map{
		"ip": ip,
	})
}
