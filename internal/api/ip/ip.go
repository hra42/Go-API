package ip

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func GetIP(c *gin.Context) {
	ip := getClientIP(c.Request)
	c.JSON(http.StatusOK, gin.H{"ip": ip})
}

func getClientIP(req *http.Request) string {
	ip := req.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(req.RemoteAddr)
	}
	return ip
}
