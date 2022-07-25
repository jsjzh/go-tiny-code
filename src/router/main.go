package router

import (
	"go-tiny-code/src/controller"
	"go-tiny-code/src/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", middleware.SayHello(), controller.Hello)

	routerGroup := router.Group("/v1/api/user")
	{
		routerGroup.POST("", controller.CreateUser)
		routerGroup.DELETE("", controller.DeleteUser)
		routerGroup.PUT("", controller.UpdateUser)
		routerGroup.GET("", controller.ReadUser)
	}

	return router
}
