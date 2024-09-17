package main

//加载gin
import (
	"GoEasyApi/route"

	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()
	route.RegisterRoutes(ginServer)

	ginServer.Run(":8008")
}
