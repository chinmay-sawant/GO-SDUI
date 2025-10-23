package api

import "github.com/gin-gonic/gin"

// RegisterRoutes registers all UI related routes on the provided router group.
func RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/nav", NavHandler)
	rg.GET("/home", HomeHandler)
	rg.GET("/products", ProductsHandler)
	rg.GET("/profile", ProfileHandler)
	rg.GET("/list-detail", ListDetailHandler)
	rg.GET("/detail/:id", DetailHandler)
	rg.GET("/forms", FormsHandler)
	rg.GET("/articles", ArticlesHandler)
}
