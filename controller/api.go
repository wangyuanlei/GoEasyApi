package controller

import "github.com/gin-gonic/gin"

func GetApi(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	ctx.JSON(200, gin.H{"path": path})
}

func PostApi(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	ctx.JSON(200, gin.H{"path": path})
}
