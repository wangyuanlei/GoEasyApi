package controller

import (
	"GoEasyApi/database"
	"GoEasyApi/helper"
	"GoEasyApi/model"

	"github.com/gin-gonic/gin"
)

var InterfaceModel = model.Interface{}

type CreateInterface struct {
	InterfaceName          string `gorm:"type:varchar(50)"`  // 接口名称
	Description            string `gorm:"type:text"`         // 接口描述
	Path                   string `gorm:"type:varchar(255)"` // 接口路径
	Method                 string `gorm:"type:varchar(10)"`  // 接口方法
	CacheEnabled           int    `gorm:"type:int"`          // 是否启用接口缓存
	CacheTime              int    `gorm:"type:int"`          // 接口缓存时间
	RateLimitEnabled       int    `gorm:"type:int"`          // 是否启用接口限流
	RateLimitCount         int    `gorm:"type:int"`          // 接口限流次数
	RateLimitTime          int    `gorm:"type:int"`          // 接口限流时间
	SqlContent             string `gorm:"type:text"`         // 接口sql语句
	TokenValidationEnabled int    `gorm:"type:int"`          // 是否启用token验证
	ReturnType             string `gorm:"type:varchar(50)"`  // 接口返回类型
}

// 增加接口
func AddInterface(ctx *gin.Context) {
	var params CreateInterface
	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	InterfaceId, err := InterfaceModel.AddInterface(database.Interface{
		InterfaceName:          params.InterfaceName,
		Description:            params.Description,
		DatabaseId:             "none",
		Path:                   params.Path,
		Method:                 params.Method,
		CacheEnabled:           params.CacheEnabled,
		CacheTime:              params.CacheTime,
		RateLimitEnabled:       params.RateLimitEnabled,
		RateLimitCount:         params.RateLimitCount,
		RateLimitTime:          params.RateLimitTime,
		SqlContent:             params.SqlContent,
		TokenValidationEnabled: params.TokenValidationEnabled,
		ReturnType:             params.ReturnType,
	})
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, InterfaceId)
}
