package route

//加载gin
import (
	"GoEasyApi/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	//健康接口
	r.GET("/health", controller.Health)
	//用户接口
	r.GET("/api/*path", controller.GetApi)        //
	r.POST("/api/*path", controller.PostApi)      //
	r.POST("/user/register", controller.Register) // 用户注册
	//管理员操作
	r.POST("/manger/login", controller.MangerLogin)                                                   //管理员登录
	r.GET("/manger/get_user_info", controller.CheckAdminLogin, controller.GetUserInfo)                //获得用户信息详情
	r.GET("/manger/get_user_list", controller.CheckAdminLogin, controller.GetUserList)                //获得用户信息列表
	r.POST("/manger/reset_pass", controller.CheckAdminLogin, controller.SetSuperAdminPasswordHandler) //设置超级管理员密码
	//名单类型配置
	r.GET("/manger/list/get_type", controller.CheckAdminLogin, controller.GetBlackListTypeHandler)  //获取黑名单类型
	r.POST("/manger/list/set_type", controller.CheckAdminLogin, controller.SetBlackListTypeHandler) //设置黑名单类型
	//白名单操作
	r.GET("/manger/whilt_list/get_list", controller.CheckAdminLogin, controller.GetAllWhiteList) //获得白名单列表
	r.POST("/manger/whilt_list/add", controller.CheckAdminLogin, controller.AddWhiteList)        //添加白名单
	r.POST("/manger/whilt_list/del", controller.CheckAdminLogin, controller.DeleteWhiteList)     //删除白名单
	//黑名单操作
	r.GET("/manger/black_list/get_list", controller.CheckAdminLogin, controller.GetAllBlackList) //获得黑名单列表
	r.POST("/manger/black_list/add", controller.CheckAdminLogin, controller.AddBlackList)        //添加黑名单
	r.POST("/manger/black_list/del", controller.CheckAdminLogin, controller.DeleteBlackList)     //删除黑名单
	//数据库配置操作
	r.GET("/manger/db/get_conf", controller.CheckAdminLogin, controller.GetUserDBConf)   //获得数据库配置
	r.POST("/manger/db/set_conf", controller.CheckAdminLogin, controller.SaveUserDBConf) //保存数据库配置

}
