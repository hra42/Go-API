package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var expectedToken string

func Init() {
	log.Println("Initializing middleware")
	expectedToken = os.Getenv("API_TOKEN")
	if expectedToken == "" {
		panic("AUTH_TOKEN environment variable not set")
	}
	return
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")

	if requiredToken == "" {
		panic("Please set the API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API token required"})
			c.Abort()
			return
		}

		if token != requiredToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
