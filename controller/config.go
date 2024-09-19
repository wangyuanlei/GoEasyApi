package controller

import (
	"GoEasyApi/helper"
	"GoEasyApi/libraries"
	"GoEasyApi/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBlackListTypeHandler 获取黑名单类型的处理器
func GetBlackListTypeHandler(ctx *gin.Context) {
	config := model.Config{}
	blackListType := config.GetBlackListType()
	helper.ApiSuccess(ctx, blackListType)
}

// SetBlackListTypeHandler 设置黑名单类型的处理器
func SetBlackListTypeHandler(ctx *gin.Context) {
	var param struct {
		BlackListType string `json:"blackListType"`
	}

	if err := ctx.ShouldBindJSON(&param); err != nil {
		helper.ApiError(ctx, 601, "无效的数据", nil)
		return
	}

	blackListType, _ := strconv.Atoi(param.BlackListType)
	config := model.Config{}

	_err := config.SetBlackListType(blackListType)
	if _err != nil {
		if myErr, ok := _err.(*libraries.CustomErrorNew); ok {
			helper.ApiError(ctx, myErr.Code, myErr.Message, nil)
		} else {
			helper.ApiError(ctx, 601, myErr.Error(), nil)
		}
	}

	helper.ApiSuccess(ctx, true)
}

// SetSuperAdminPasswordHandler 设置超级管理员密码的处理器
func SetSuperAdminPasswordHandler(ctx *gin.Context) {
	var passwordData struct {
		OldPassword string `json:"oldPass"`
		NewPassword string `json:"newPass"`
	}

	if err := ctx.ShouldBindJSON(&passwordData); err != nil {
		helper.ApiError(ctx, 601, "无效的数据", nil)
		return
	}

	config := model.Config{}
	err := config.SetSuperAdminPassword(passwordData.OldPassword, passwordData.NewPassword)
	if err != nil {
		if myErr, ok := err.(*libraries.CustomErrorNew); ok {
			helper.ApiError(ctx, myErr.Code, myErr.Message, nil)
		} else {
			helper.ApiError(ctx, 601, myErr.Error(), nil)
		}
		return
	}

	helper.ApiSuccess(ctx, true)
}
