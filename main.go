package main

//加载gin
import (
	"GoEasyApi/libraries"
	"GoEasyApi/route"

	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()
	route.RegisterRoutes(ginServer)

	bindAddress, _ := libraries.GetBind()
	if bindAddress == "" {
		bindAddress = ":8008"
	}

	ginServer.Run(bindAddress)
}
