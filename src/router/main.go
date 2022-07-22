package router

import (
	"go-tiny-code/src/controller"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health.check", controller.Hello)

	routerGroup := router.Group("/v1/api")
	{
		routerGroup.GET("/ping", controller.Hello)
	}

	return router
}
