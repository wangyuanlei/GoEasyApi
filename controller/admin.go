package controller

import (
	"GoEasyApi/helper"
	"GoEasyApi/model"

	"github.com/gin-gonic/gin"
)

func MangerLogin(ctx *gin.Context) {
	var loginData struct {
		Username string `json:"account"`
		Password string `json:"pass"`
	}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	admin := model.Admin{}
	token, err := admin.Login(loginData.Username, loginData.Password)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, token)
}
