package helper

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

// HashPassword 函数用于加密密码
func HashPassword(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	firstHash := hex.EncodeToString(hasher.Sum(nil))
	hasher.Reset()
	hasher.Write([]byte(firstHash + "Qwert!@#456"))
	return hex.EncodeToString(hasher.Sum(nil))
}

// 接口请求成功, 返回数据格式:{"code":200,"msg":"","data":[]}
func ApiSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "",
		"data": data,
	})
}

// 接口请求失败, 返回数据格式:{"code":0,"msg":"","data":[]}
func ApiError(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
