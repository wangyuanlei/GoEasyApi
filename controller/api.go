package controller

import (
	"GoEasyApi/helper"
	"GoEasyApi/model"

	"github.com/gin-gonic/gin"
)

var _ApiModel = model.Api{}

func GetApi(ctx *gin.Context) {
	// path := ctx.Request.URL.Path
	// ctx.JSON(200, gin.H{"path": path})
	data, err := _ApiModel.Get(ctx)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, data)
}

func PostApi(ctx *gin.Context) {
	// path := ctx.Request.URL.Path
	// ctx.JSON(200, gin.H{"path": path})
	data, err := _ApiModel.Post(ctx)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, data)
}
