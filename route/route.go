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
	r.GET("/api/*path", controller.GetApi)                                          //
	r.POST("/api/*path", controller.PostApi)                                        //
	r.POST("/user/register", controller.UserRegister)                               // 用户注册
	r.POST("/user/login", controller.UserLogin)                                     // 用户登入
	r.GET("/user/user_info", controller.CheckUserLogin, controller.GetUserInfoByMe) // 登录用户获得个人信息
	//管理员操作
	r.POST("/manger/login", controller.MangerLogin)                                                //管理员登录
	r.GET("/manger/get_user_info", controller.CheckAdminLogin, controller.GetUserInfo)             //获得用户信息详情
	r.GET("/manger/get_user_list", controller.CheckAdminLogin, controller.GetUserList)             //获得用户信息列表
	r.POST("/manger/set_user_pass", controller.CheckAdminLogin, controller.AdminChangeUserPassord) //管理员设置用户密码
	r.POST("/manger/set_user_info", controller.CheckAdminLogin, controller.ChangeUserInfo)         //管理员设置用户信息

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

	//接口操作
	r.POST("/manger/interface/add", controller.CheckAdminLogin, controller.AddInterface)       //添加接口
	r.POST("/manger/interface/update", controller.CheckAdminLogin, controller.UpdateInterface) //修改接口
	r.GET("/manger/interface/list", controller.CheckAdminLogin, controller.GetList)            //获得接口列表
	r.GET("/manger/interface/info", controller.CheckAdminLogin, controller.GetInfo)            //获得接口详情
	r.POST("/manger/interface/delete", controller.CheckAdminLogin, controller.DeleteInterface) //删除接口
}
