package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hra42/Go-API/internal/api/ip"
	"github.com/hra42/Go-API/internal/api/middleware"
)

func StartServer() {
	router := gin.Default()
	middleware.Init()
	router.Use(middleware.TokenAuthMiddleware())

	router.GET("/", index)
	router.GET("/ip", ip.GetIP)

	router.Run(":8080")
}
