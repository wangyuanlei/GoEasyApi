package controller

import (
	"GoEasyApi/helper"
	"GoEasyApi/model"
	"GoEasyApi/structs"

	"github.com/gin-gonic/gin"
)

var InterfaceModel = model.Interface{}

// 增加接口
func AddInterface(ctx *gin.Context) {
	var params structs.Interface
	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	params.DatabaseId = "none"
	InterfaceId, err := InterfaceModel.AddInterface(params)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, InterfaceId)
}

// 增加接口
func UpdateInterface(ctx *gin.Context) {
	var params structs.Interface
	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	params.DatabaseId = "none"

	err := InterfaceModel.UpdateInterface(params)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, params.Id)
}

// 获得所有接口
func GetList(ctx *gin.Context) {
	interfaceList, err := InterfaceModel.GetList()
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, interfaceList)
}

// 获得接口详情
func GetInfo(ctx *gin.Context) {
	interface_id, ok := ctx.GetQuery("id")
	if !ok {
		helper.ApiError(ctx, 601, "参数 id 未提交", nil)
		return
	}

	interfaceList, err := InterfaceModel.GetInfo(interface_id)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, interfaceList)
}

func DeleteInterface(ctx *gin.Context) {
	interface_id, ok := ctx.GetQuery("id")
	if !ok {
		helper.ApiError(ctx, 601, "参数 id 未提交", nil)
		return
	}

	err := InterfaceModel.DeleteInterface(interface_id)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, true)
}
