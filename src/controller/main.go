package controller

import (
	"github.com/gin-gonic/gin"

	"go-tiny-code/src/service"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(200, service.Hello())
}
