package helper

import (
	"GoEasyApi/cron"
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strconv"
	"strings"

	"math/rand"
	"time"

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

// DoubleHashPassword 函数用于对密码进行双重加密
func DoubleHashPassword(pass, salt string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	firstHash := hex.EncodeToString(hasher.Sum(nil))
	hasher.Reset()
	hasher.Write([]byte(firstHash + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}

func HashMD5(val string) string {
	hasher := md5.New()
	hasher.Write([]byte(val))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GenerateRandomString 生成一个随机的6位字符串
func GenerateRandomString(length int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
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

// 判断是否是正常的ip地址.
func IsValidIP(ip string) bool {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}
	for _, item := range parts {
		if !strings.HasPrefix(item, "0") {
			num, err := strconv.Atoi(item)
			if err != nil || num < 0 || num > 255 {
				return false
			}
		}
	}
	return true
}

func InArray(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func CheckParamItem(item string, itemList string) error {
	regText := "^(" + itemList + ")$"
	if !regexp.MustCompile(regText).MatchString(item) {

		return cron.CreateCustomError(601, "参数类型错误:"+item+"不是有效的值, 只能是以下类型"+itemList)
	}

	return nil
}

// 检查method 类型 只能是 post 和 get
func CheckMethod(method string) error {
	if method != "post" && method != "get" {
		return cron.CreateCustomError(601, "Method 只能是 get 或者 post")
	}

	return nil
}

// 检查是否有效的值
func CheckEnabled(data int) error {
	if data != 1 && data != 2 {
		return cron.CreateCustomError(601, "值设置错误")
	}

	return nil
}

// 检查参数的名称
func CheckStringFormat(str string) error {
	if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(str) {
		return cron.CreateCustomError(601, "字符串只能包含大小写字母、数字和下划线")
	}
	return nil
}

// 正则表达式验证
func MatchString(reg string, str string) bool {
	return regexp.MustCompile(reg).MatchString(str)
}
