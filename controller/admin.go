package controller

import (
	"GoEasyApi/helper"
	"GoEasyApi/libraries"
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
		if myErr, ok := err.(*libraries.CustomErrorNew); ok {
			helper.ApiError(ctx, myErr.Code, myErr.Message, nil)
		} else {
			helper.ApiError(ctx, 601, myErr.Error(), nil)
		}
		return
	}

	helper.ApiSuccess(ctx, token)
}
