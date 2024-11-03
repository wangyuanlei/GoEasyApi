package controller

import (
	"GoEasyApi/helper"
	"GoEasyApi/model"
	"GoEasyApi/structs"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBlackListTypeHandler 获取黑名单类型
func GetBlackListTypeHandler(ctx *gin.Context) {
	config := model.Config{}
	blackListType := config.GetBlackListType()
	helper.ApiSuccess(ctx, blackListType)
}

// SetBlackListTypeHandler 设置黑名单类型
func SetBlackListTypeHandler(ctx *gin.Context) {
	var param structs.BlackListTypeParams

	if err := ctx.ShouldBindJSON(&param); err != nil {
		helper.ApiError(ctx, 601, "无效的数据", nil)
		return
	}

	blackListType, _ := strconv.Atoi(param.BlackListType)
	config := model.Config{}

	fmt.Println("blackListType", blackListType)

	_err := config.SetBlackListType(blackListType)
	if _err != nil {
		ShowModelError(ctx, _err)
		return
	}

	helper.ApiSuccess(ctx, true)
}

// SetSuperAdminPasswordHandler 设置超级管理员密码
func SetSuperAdminPasswordHandler(ctx *gin.Context) {
	var passwordData structs.AdminPasswordParams
	if err := ctx.ShouldBindJSON(&passwordData); err != nil {
		helper.ApiError(ctx, 601, "无效的数据", nil)
		return
	}

	config := model.Config{}
	err := config.SetSuperAdminPassword(passwordData.OldPassword, passwordData.NewPassword)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, true)
}
