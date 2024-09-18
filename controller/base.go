package controller

import (
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
