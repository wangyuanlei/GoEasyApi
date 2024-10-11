package controller

import (
	"GoEasyApi/database"
	"GoEasyApi/helper"
	"GoEasyApi/model"

	"github.com/gin-gonic/gin"
)

var InterfaceModel = model.Interface{}

// 增加接口
func AddInterface(ctx *gin.Context) {
	var params database.Interface
	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}
	InterfaceId, err := InterfaceModel.AddInterface(params)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, InterfaceId)
}
