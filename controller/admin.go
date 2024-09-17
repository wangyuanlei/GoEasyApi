package controller

import (
	"GoEasyApi/model"

	"github.com/gin-gonic/gin"
)

func MangerLogin(ctx *gin.Context) {
	var loginData struct {
		Username string `json:"account"`
		Password string `json:"pass"`
	}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(400, gin.H{"error": "请求数据格式错误"})
		return
	}

	admin := model.Admin{}
	token, err := admin.Login(loginData.Username, loginData.Password)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}
