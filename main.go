package main

//加载gin
import (
	"GoEasyApi/libraries"
	"GoEasyApi/model"
	"GoEasyApi/route"

	"github.com/gin-gonic/gin"
)

// 初始化model
func init() {
	model.InitDB() //连接数据库
	model.CreateNewDBFile()
}

func main() {
	ginServer := gin.Default()
	route.RegisterRoutes(ginServer)

	bindAddress, _ := libraries.GetBind()
	if bindAddress == "" {
		bindAddress = ":8008"
	}

	ginServer.Run(bindAddress)
}
