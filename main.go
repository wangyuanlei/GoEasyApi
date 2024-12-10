package main

//加载gin
import (
	"GoEasyApi/libraries"
	"GoEasyApi/model"
	"GoEasyApi/route"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 初始化model
func init() {
	model.InitDB()          //连接数据库
	model.CreateNewDBFile() //创建系统数据库

}

func main() {
	ginServer := gin.Default()
	// 自定义请求头中设置的用户的真实Ip地址
	ginServer.TrustedPlatform = "Client-IP"

	// ginServer.StaticFS("/static", http.Dir("./manger/dist")) //加载静态文件
	ginServer.StaticFS("/static", http.Dir("./vue/dist")) //加载静态文件
	ginServer.StaticFile("/favicon.ico", "./static/favicon.ico")

	route.RegisterRoutes(ginServer)
	bindAddress, _ := libraries.GetBind()
	if bindAddress == "" {
		bindAddress = ":8008"
	}

	ginServer.Run(bindAddress)
}
