package controller

import (
	"GoEasyApi/database"
	"GoEasyApi/helper"
	"GoEasyApi/libraries"
	"GoEasyApi/model"

	"github.com/gin-gonic/gin"
)

var DB = libraries.InitDB()
var WhiteListModel = model.WhiteList{}

func GetAllWhiteList(ctx *gin.Context) {
	list, err := WhiteListModel.GetAllWhiteList(DB)
	if err != nil {
		if myErr, ok := err.(*libraries.CustomErrorNew); ok {
			helper.ApiError(ctx, myErr.Code, myErr.Message, nil)
		} else {
			helper.ApiError(ctx, 601, myErr.Error(), nil)
		}
		return
	}

	helper.ApiSuccess(ctx, list)
}

func AddWhiteList(ctx *gin.Contex) {
	var params = &database.WhiteList
	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	list, err := WhiteListModel.AddWhiteList(DB, params)
	if err != nil {
		if myErr, ok := err.(*libraries.CustomErrorNew); ok {
			helper.ApiError(ctx, myErr.Code, myErr.Message, nil)
		} else {
			helper.ApiError(ctx, 601, myErr.Error(), nil)
		}
		return
	}

	helper.ApiSuccess(ctx, list)
}

func GetAllBlackList(ctx *gin.Context) {
	list, err := WhiteListModel.GetAllBlackList(DB)
	if err != nil {
		if myErr, ok := err.(*libraries.CustomErrorNew); ok {
			helper.ApiError(ctx, myErr.Code, myErr.Message, nil)
		} else {
			helper.ApiError(ctx, 601, myErr.Error(), nil)
		}
		return
	}

	helper.ApiSuccess(ctx, list)
}
