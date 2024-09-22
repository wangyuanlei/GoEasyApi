package controller

import (
	"GoEasyApi/helper"
	"GoEasyApi/model"

	"github.com/gin-gonic/gin"
)

var WhiteListModel = model.WhiteList{}

func GetAllWhiteList(ctx *gin.Context) {
	list, err := WhiteListModel.GetAllWhiteList()
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, list)
}

func AddWhiteList(ctx *gin.Context) {
	var params struct {
		Ip          string `json:"ip"`
		Description string `json:"description"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	err := WhiteListModel.AddWhiteList(params.Ip, params.Description)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, true)
}

func DeleteWhiteList(ctx *gin.Context) {
	var params struct {
		Ip string `json:"ip"`
	}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	err := WhiteListModel.DeleteWhiteList(params.Ip)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, true)
}

func GetAllBlackList(ctx *gin.Context) {
	list, err := WhiteListModel.GetAllBlackList()
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, list)
}

func AddBlackList(ctx *gin.Context) {
	var params struct {
		Ip          string `json:"ip"`
		Description string `json:"description"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	err := WhiteListModel.AddBlackList(params.Ip, params.Description)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, true)
}

func DeleteBlackList(ctx *gin.Context) {
	var params struct {
		Ip string `json:"ip"`
	}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	err := WhiteListModel.DeleteBlackList(params.Ip)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, true)
}
