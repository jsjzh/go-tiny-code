package main

import "go-tiny-code/src/router"

func main() {
	r := router.InitializeRouter()

	r.Run("127.0.0.1:7001")
}
