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

	r.GET("/get_black_list_type", controller.CheckAdminLogin, controller.GetBlackListTypeHandler)      //获取黑名单类型
	r.POST("/set_black_list_type", controller.CheckAdminLogin, controller.SetBlackListTypeHandler)     //设置黑名单类型
	r.POST("/set_admin_password", controller.CheckAdminLogin, controller.SetSuperAdminPasswordHandler) //设置超级管理员密码

	r.GET("/get_whilt_list", controller.CheckAdminLogin, controller.GetAllWhiteList)  //获得白名单列表
	r.GET("/get_black_list", controller.CheckAdminLogin, controller.GetAllBlackList)  //获得黑名单列表
	r.POST("/add_whilt_list", controller.CheckAdminLogin, controller.AddWhiteList)    //添加白名单
	r.POST("/add_black_list", controller.CheckAdminLogin, controller.AddBlackList)    //添加黑名单
	r.POST("/del_whilt_list", controller.CheckAdminLogin, controller.DeleteWhiteList) //删除白名单
	r.POST("/del_black_list", controller.CheckAdminLogin, controller.DeleteBlackList) //删除黑名单

	r.GET("/get_db_conf", controller.CheckAdminLogin, controller.GetUserDBConf)   //获得数据库配置
	r.POST("/set_db_conf", controller.CheckAdminLogin, controller.SaveUserDBConf) //保存数据库配置
}
