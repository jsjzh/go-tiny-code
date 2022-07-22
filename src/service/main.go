package service

import "github.com/gin-gonic/gin"

func Hello() map[string]any {
	return gin.H{
		"message": "hello",
	}
}
