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
	// ginServer.TrustedPlatform = "Client-IP"
	// err := ginServer.SetTrustedProxies([]string{"192.168.1.1", "10.0.0.1", "127.0.0.1"})
	// if err != nil {
	// 	log.Fatal("设置信任的代理失败")
	// }
	ginServer.StaticFS("/static", http.Dir("./static")) //加载静态文件
	ginServer.StaticFile("/favicon.ico", "./static/favicon.ico")

	route.RegisterRoutes(ginServer)
	bindAddress, _ := libraries.GetBind()
	if bindAddress == "" {
		bindAddress = ":8008"
	}

	ginServer.Run(bindAddress)
}
