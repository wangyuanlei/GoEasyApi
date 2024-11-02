package controller

import (
	"GoEasyApi/helper"
	"GoEasyApi/model"
	"GoEasyApi/structs"

	"github.com/gin-gonic/gin"
)

var DataBaseModel = model.DataBase{}

// 获得数据库配置
func GetUserDBConf(ctx *gin.Context) {
	conf, err := DataBaseModel.GetUserDBConf()
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, conf)
}

// 保存数据库配置
func SaveUserDBConf(ctx *gin.Context) {
	var params structs.CreateDatabaseParams

	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	err := DataBaseModel.UpdateUserDBConf(params.Name, params.Description, params.OrmType, params.Dns)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, true)
}
