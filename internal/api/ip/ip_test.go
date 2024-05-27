package ip

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gofiber/fiber/v2"
    "github.com/stretchr/testify/assert"
)

func TestHandleCurrentIP(t *testing.T) {
    // Initialize the Fiber app
    app := fiber.New()

    // Register the route
    app.Get("/current/ip", HandleCurrentIP)

    tests := []struct {
        name           string
        ipHeader       string
        remoteAddr     string
        expectedStatus int
        expectedBody   string
    }{
        {
            name:           "No IP addresses",
            ipHeader:       "",
            remoteAddr:     "",
            expectedStatus: http.StatusNotFound,
            expectedBody:   `{"error":"No IPv4 address found"}`,
        },
        {
            name:           "Only IPv4 address",
            ipHeader:       "203.0.113.195",
            remoteAddr:     "",
            expectedStatus: http.StatusOK,
            expectedBody:   `{"ip":"203.0.113.195"}`,
        },
        {
            name:           "Only IPv6 address",
            ipHeader:       "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
            remoteAddr:     "",
            expectedStatus: http.StatusNotFound,
            expectedBody:   `{"error":"No IPv4 address found"}`,
        },
        {
            name:           "Both IPv4 and IPv6 addresses",
			ipHeader:       "2001:0db8:85a3:0000:0000:8a2e:0370:7334, 203.0.113.195",
            remoteAddr:     "",
            expectedStatus: http.StatusOK,
            expectedBody:   `{"ip":"203.0.113.195"}`,
        },
        {
            name:           "No forwarded IPs, direct IPv4",
            ipHeader:       "",
            remoteAddr:     "203.0.113.195:1234",
            expectedStatus: http.StatusOK,
            expectedBody:   `{"ip":"203.0.113.195"}`,
        },
        {
            name:           "No forwarded IPs, direct IPv6",
            ipHeader:       "",
            remoteAddr:     "[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:1234",
            expectedStatus: http.StatusNotFound,
            expectedBody:   `{"error":"No IPv4 address found"}`,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Create a request
            req := httptest.NewRequest(http.MethodGet, "/current/ip", nil)
            if tt.ipHeader != "" {
                req.Header.Set("X-Forwarded-For", tt.ipHeader)
            }
            if tt.remoteAddr != "" {
                req.RemoteAddr = tt.remoteAddr
            }

            // Create a response recorder
            resp, err := app.Test(req, -1)
            assert.NoError(t, err)

            // Check the status code
            assert.Equal(t, tt.expectedStatus, resp.StatusCode)

            // Check the response body
            body := make([]byte, resp.ContentLength)
            resp.Body.Read(body)
            assert.JSONEq(t, tt.expectedBody, string(body))
        })
    }
}