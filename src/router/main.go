package router

import (
	"go-tiny-code/src/controller"
	"go-tiny-code/src/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", middleware.SayHello(), controller.Hello)

	routerGroup := router.Group("/v1/api")
	{
		routerGroup.GET("/ping", controller.Hello)
	}

	return router
}
