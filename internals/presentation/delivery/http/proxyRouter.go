package server

import (
	"github.com/gin-gonic/gin"
	"github.com/maestre3d/placehub-payment/internals/presentation/delivery/http/handler"
)

func NewProxyRouter(mux *gin.Engine) *gin.RouterGroup {
	return mux.Group("/api/v1")
}

func Register(u *handler.UserHandler) {
	// User Handler set routes
	// u.SetRoutes()
}
