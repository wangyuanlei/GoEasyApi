package main

//加载gin
import (
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()

	ginServer.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})

	ginServer.GET("/api/*path", func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		ctx.JSON(200, gin.H{"path": path})
	})

	ginServer.Run(":8008")
}
