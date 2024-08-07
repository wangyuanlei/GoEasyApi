package main

//加载gin
import (
	"GoEasyApi/database"
	"net"

	"github.com/gin-gonic/gin"
)

func isPrivateIP(ip string) bool {
	privateIPBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.1/8",
	}

	for _, block := range privateIPBlocks {
		_, ipNet, _ := net.ParseCIDR(block)
		if ipNet.Contains(net.ParseIP(ip)) {
			return true
		}
	}
	return false
}

func main() {

	database.InitMysql()
	ginServer := gin.Default()

	ginServer.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})

	ginServer.GET("/ip", func(ctx *gin.Context) {
		clientIP := ctx.ClientIP()
		if !isPrivateIP(clientIP) {
			ctx.JSON(403, gin.H{"message": "非局域网IP，访问被拒绝"})
			return
		}
		ctx.JSON(200, gin.H{"message": "局域网IP，访问成功"})
	})

	ginServer.GET("/sql", func(ctx *gin.Context) {
		results := database.GetList()
		ctx.JSON(200, gin.H{"data": results})
	})

	ginServer.POST("/sql", func(ctx *gin.Context) {
		var pageInfo struct {
			PageNo   int `json:"pageNo"`
			PageSize int `json:"pageSize"`
		}
		if err := ctx.ShouldBindJSON(&pageInfo); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		results := database.GetListPage(pageInfo.PageNo, pageInfo.PageSize)
		ctx.JSON(200, gin.H{"data": results})
	})

	ginServer.GET("/api/*path", func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		ctx.JSON(200, gin.H{"path": path})
	})

	ginServer.Run(":8008")
}
