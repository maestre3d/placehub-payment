package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/maestre3d/placehub-payment/internals/shared/domain/util"
	"log"
	"net/http"
)

type UserHandler struct {
	mux    *gin.RouterGroup
	logger *log.Logger
}

func NewUserHandler(logger *log.Logger, mux *gin.RouterGroup) (*UserHandler, error) {
	logger.Printf("Executing new User Handler")
	userHandler := &UserHandler{
		mux:    mux,
		logger: logger,
	}

	userHandler.setRoutes()

	return userHandler, nil
}

func (u *UserHandler) setRoutes() *gin.RouterGroup {
	users := u.mux.Group("/user")
	{
		users.GET("/", u.GetUser)
	}

	return users
}

func (u *UserHandler) GetUser(c *gin.Context) {
	u.logger.Printf("Got a request")
	c.JSON(http.StatusOK, &util.Response{Message: "Hello there from Gin Users"})
}
