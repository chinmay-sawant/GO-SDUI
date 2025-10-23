package main

import (
	"net/http"

	"sdui/internal/api"
	"sdui/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Simple CORS middleware
	r.Use(middleware.CORSMiddleware())

	apiRouter := r.Group("/api/ui")
	api.RegisterRoutes(apiRouter)

	// Simple health
	r.GET("/health", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	_ = r.Run(":8080")
}
