package api

import (
	"net/http"

	"sdui/internal/services"

	"github.com/gin-gonic/gin"
)

// Adapter handlers: convert services.UIService results into HTTP responses.
func NavHandler(svc services.UIService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusOK, svc.Nav()) }
}

func HomeHandler(svc services.UIService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusOK, svc.Home()) }
}

func ProductsHandler(svc services.UIService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusOK, svc.Products()) }
}

func ProfileHandler(svc services.UIService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusOK, svc.Profile()) }
}

func ListDetailHandler(svc services.UIService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusOK, svc.ListDetail()) }
}

func DetailHandler(svc services.UIService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if comps, ok := svc.Detail(id); ok {
			c.JSON(http.StatusOK, comps)
		} else {
			c.JSON(http.StatusNotFound, comps)
		}
	}
}

func FormsHandler(svc services.UIService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusOK, svc.Forms()) }
}

func ArticlesHandler(svc services.UIService) gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusOK, svc.Articles()) }
}
