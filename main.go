package main

import (
	"go-tiny-code/config"
	"go-tiny-code/src/model"
	"go-tiny-code/src/router"
	"go-tiny-code/src/shared"
)

func main() {
	c := config.InitializeConfig()
	r := router.InitializeRouter()
	model.InitializeMysql(c)

	ip := shared.GetLocalIpAddress(int(c.Dev.Port))

	r.Run(ip)
}
