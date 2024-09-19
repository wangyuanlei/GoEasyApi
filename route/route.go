package route

//加载gin
import (
	"GoEasyApi/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	//健康接口
	r.GET("/health", controller.Health)

	r.GET("/api/*path", controller.GetApi)
	r.POST("/api/*path", controller.PostApi)

	r.POST("/manger_login", controller.MangerLogin) //管理员登录

	r.GET("/get_black_list_type", controller.CheckAdminLogin, controller.GetBlackListTypeHandler)  //获取黑名单类型
	r.POST("/set_black_list_type", controller.CheckAdminLogin, controller.SetBlackListTypeHandler) //设置黑名单类型

	r.POST("/set_admin_password", controller.CheckAdminLogin, controller.SetSuperAdminPasswordHandler) //设置超级管理员密码
}
