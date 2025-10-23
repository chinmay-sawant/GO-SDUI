package api

import (
	"sdui/internal/services"

	"github.com/gin-gonic/gin"
)

// Routes is an interface capable of registering UI routes onto a gin RouterGroup.
// This allows the registration behaviour to be swapped or mocked in tests.
type Routes interface {
	Register(rg *gin.RouterGroup)
}

// DefaultRoutes is the package-provided implementation of Routes.
// It holds a UIService which is used by the handlers it registers.
type DefaultRoutes struct {
	svc services.UIService
}

// NewDefaultRoutes constructs a DefaultRoutes instance using the provided UIService.
func NewDefaultRoutes(svc services.UIService) Routes {
	return &DefaultRoutes{svc: svc}
}

// Register registers all UI related routes on the provided router group using the
// service attached to the DefaultRoutes receiver.
func (dr *DefaultRoutes) Register(rg *gin.RouterGroup) {
	rg.GET("/nav", NavHandler(dr.svc))
	rg.GET("/home", HomeHandler(dr.svc))
	rg.GET("/products", ProductsHandler(dr.svc))
	rg.GET("/profile", ProfileHandler(dr.svc))
	rg.GET("/list-detail", ListDetailHandler(dr.svc))
	rg.GET("/detail/:id", DetailHandler(dr.svc))
	rg.GET("/forms", FormsHandler(dr.svc))
	rg.GET("/articles", ArticlesHandler(dr.svc))
	// Combined overview page assembled from multiple sub-screens
	rg.GET("/combined", CombinedHandler(dr.svc))
}

// RegisterRoutes preserves the original helper signature while delegating to
// the DefaultRoutes implementation.
func RegisterRoutes(rg *gin.RouterGroup, svc services.UIService) {
	NewDefaultRoutes(svc).Register(rg)
}

// RegisterRoutesDefault preserves the old signature by wiring the default service.
func RegisterRoutesDefault(rg *gin.RouterGroup) {
	RegisterRoutes(rg, services.NewDefaultUIService())
}
