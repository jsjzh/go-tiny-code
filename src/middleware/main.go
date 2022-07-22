package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SayHello() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("hello")
		ctx.Next()
	}
}
