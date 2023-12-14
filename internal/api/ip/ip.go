package ip

import (
	"github.com/gofiber/fiber/v2"
	"net"
)

func HandleCurrentIP(c *fiber.Ctx) error {
	ip := c.IPs()[0]
	return c.JSON(fiber.Map{
		"ip": ip,
	})
}

func HandleCurrentIPv4(c *fiber.Ctx) error {
	// return ipv4 address from request
	ip := c.IPs()[0]
	ipv4 := net.ParseIP(ip).To4()
	return c.JSON(fiber.Map{
		"ip": ipv4,
	})
}

func HandleCurrentIPv6(c *fiber.Ctx) error {
	// return ipv6 address from request
	ip := c.IPs()[0]
	ipv6 := net.ParseIP(ip).To16()
	return c.JSON(fiber.Map{
		"ip": ipv6,
	})
}
