package server

import (
	"github.com/gin-gonic/gin"
)

func NewProxyRouter(mux *gin.Engine) *gin.RouterGroup {
	return mux.Group("/api/v1")
}
