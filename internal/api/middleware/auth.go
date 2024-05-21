package middleware

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var expectedToken string

func Init() {
	log.Println("Initializing middleware")
	expectedToken = os.Getenv("AUTH_TOKEN")
	if expectedToken == "" {
		panic("AUTH_TOKEN environment variable not set")
	}
	return
}

func AuthRequired(c *fiber.Ctx) error {
	// Check for a token in the headers
	token := c.Get("Authorization")
	if token != expectedToken {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.Next()
}