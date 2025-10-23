package middleware

import "github.com/gin-gonic/gin"

// CorsProvider allows supplying a CORS handler; implementations can be swapped for tests.
type CorsProvider interface {
	Handler() gin.HandlerFunc
}

// DefaultCors is the default implementation used by the application.
type DefaultCors struct{}

// Handler returns the Gin middleware that sets standard CORS headers.
func (d *DefaultCors) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// CORSMiddleware keeps the original helper function for compatibility.
func CORSMiddleware() gin.HandlerFunc {
	return (&DefaultCors{}).Handler()
}

// NewDefaultCors returns a CorsProvider instance with default behaviour.
func NewDefaultCors() CorsProvider { return &DefaultCors{} }
