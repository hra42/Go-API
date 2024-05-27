package ip

import (
    "github.com/gofiber/fiber/v2"
    "net"
)

// return the current user IPv4 IP of the request
func HandleCurrentIP(c *fiber.Ctx) error {
    // get the current user IPs
    ips := c.IPs()
    var ip string

    // helper function to check if IP is IPv4
    isIPv4 := func(ip string) bool {
        parsedIP := net.ParseIP(ip)
        return parsedIP != nil && parsedIP.To4() != nil
    }

    if len(ips) > 0 {
        // iterate over IPs and find the first IPv4
        for _, candidateIP := range ips {
            if isIPv4(candidateIP) {
                ip = candidateIP
                break
            }
        }
    }

    // fallback to direct IP if no forwarded IPv4 IPs are found
    if ip == "" {
        remoteAddr := c.Context().RemoteAddr().String()
        if remoteAddr != "" {
            remoteIP, _, err := net.SplitHostPort(remoteAddr)
            if err == nil && isIPv4(remoteIP) {
                ip = remoteIP
            } else if isIPv4(remoteAddr) {
                ip = remoteAddr
            }
        }
    }

    if ip == "" || ip == "0.0.0.0" {
        // handle the case when no IPv4 address is found or IP is "0.0.0.0"
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "No IPv4 address found",
        })
    }

    // json return of the IP
    return c.JSON(fiber.Map{
        "ip": ip,
    })
}