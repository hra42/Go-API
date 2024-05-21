package ip

import (
	"github.com/gofiber/fiber/v2"
)

// return the current user IPv4 IP of the request
func HandleCurrentIP(c *fiber.Ctx) error {
	// get the current user IPs
	ips := c.IPs()
	var ip string

	if len(ips) > 0 {
		ip = ips[0]
	} else {
		// fallback to direct IP if no forwarded IPs are found
		ip = c.IP()
	}

	// json return of the IP
	return c.JSON(fiber.Map{
		"ip": ip,
	})
}
