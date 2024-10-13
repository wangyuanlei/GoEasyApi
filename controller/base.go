package controller

import (
	"GoEasyApi/cron"
	"GoEasyApi/helper"
	"GoEasyApi/model"

	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	helper.ApiSuccess(ctx, "ok")
}

// 验证管理员登录的中间件
func CheckAdminLogin(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		helper.ApiError(ctx, 502, "Token 未提交", nil)
		ctx.Abort()
		return
	}

	admin := model.Admin{}
	_, err := admin.ValidateLogin(token)

	if err != nil {
		helper.ApiError(ctx, 501, "验证失败", nil)
		ctx.Abort()
		return
	}

	ctx.Next()
}

// 检查用户是否登录
func CheckUserLogin(ctx *gin.Context) {
	token := ctx.GetHeader("usertoken")
	if token == "" {
		helper.ApiError(ctx, 502, "用户 Token 未提交", nil)
		ctx.Abort()
		return
	}
	// userModel := model.User{}
	tokenModel := model.Token{}
	userId, err := tokenModel.GetTokenInfo(token)
	if err != nil {
		ShowModelError(ctx, err)
		ctx.Abort()
		return
	}

	//token 有效时间延续2小时
	tokenModel.TokenExtendTime(userId, token)

	ctx.Set("USERID", userId)
	ctx.Next()
}

func ShowModelError(ctx *gin.Context, err error) {
	if err != nil {
		if myErr, ok := err.(*cron.CustomErrorNew); ok {
			helper.ApiError(ctx, myErr.Code, myErr.Message, nil)
		} else {
			helper.ApiError(ctx, 601, err.Error(), nil)
		}
	}
}
