package ip

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientIP(t *testing.T) {
	tests := []struct {
		name       string
		header     string
		remoteAddr string
		expectedIP string
	}{
		{
			name:       "X-Forwarded-For header present",
			header:     "203.0.113.195",
			remoteAddr: "192.0.2.1:12345",
			expectedIP: "203.0.113.195",
		},
		{
			name:       "X-Forwarded-For header absent",
			header:     "",
			remoteAddr: "192.0.2.1:12345",
			expectedIP: "192.0.2.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.header != "" {
				req.Header.Set("X-Forwarded-For", tt.header)
			}
			req.RemoteAddr = tt.remoteAddr

			ip := getClientIP(req)
			assert.Equal(t, tt.expectedIP, ip)
		})
	}
}
