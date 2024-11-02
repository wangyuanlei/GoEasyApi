package controller

import (
	"GoEasyApi/helper"
	"GoEasyApi/model"
	"GoEasyApi/structs"

	"github.com/gin-gonic/gin"
)

func MangerLogin(ctx *gin.Context) {
	var loginData structs.AdminLoginParams
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	admin := model.Admin{}
	token, err := admin.Login(loginData.Account, loginData.Password)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, token)
}
