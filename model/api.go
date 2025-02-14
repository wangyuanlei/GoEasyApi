package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/helper"
	"GoEasyApi/libraries"
	"GoEasyApi/structs"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Api struct{}

var InterfaceModel = Interface{}
var Api_Model = ApiModel{}

func (m *Api) Get(ctx *gin.Context) (interface{}, error) {

	//判断名单
	if err := m.CheckIp(ctx); err != nil {
		return nil, err
	}

	//获得接口信息
	interfaceInfo, err := InterfaceModel.GetInfoByPath(ctx.Request.URL.Path, "get")
	if err != nil {
		return nil, err
	}

	//判断是否限流
	if err := m.CheckRateLimit(ctx.ClientIP(), interfaceInfo); err != nil {
		return nil, err
	}

	//判断是否使用token验证
	if interfaceInfo.TokenValidationEnabled == 1 {
		if err := m.CheckUserLogin(ctx); err != nil {
			return nil, err
		}
	}

	//验证参数信息
	params, err := m.CheckParams(ctx, interfaceInfo)
	if err != nil {
		return nil, err
	}

	//判断是否开启缓存
	cacheKey := m.GetCacheKeyByParams(ctx, interfaceInfo)
	if interfaceInfo.CacheEnabled == 1 {
		_CacheData, IsExists := libraries.GetCache(cacheKey)
		if IsExists { //存在缓存 从缓存中获取
			return _CacheData, nil
		}
	}
	Api_Model.InitApiDB()
	data, err := Api_Model.Run(interfaceInfo, params)
	if err != nil {
		return nil, err
	}

	if interfaceInfo.ReturnType == "insert" ||
		interfaceInfo.ReturnType == "update" ||
		interfaceInfo.ReturnType == "delete" {
		//这几种类型不错缓存
	} else {
		if interfaceInfo.CacheEnabled == 1 { //开启缓存 写入缓存
			libraries.AddCache(cacheKey, data, time.Duration(interfaceInfo.CacheTime)*time.Second)
		}
	}

	return data, nil
}

func (m *Api) Post(ctx *gin.Context) (interface{}, error) {
	//判断名单
	if err := m.CheckIp(ctx); err != nil {
		return nil, err
	}

	//获得接口信息
	interfaceInfo, err := InterfaceModel.GetInfoByPath(ctx.Request.URL.Path, "post")
	if err != nil {
		return nil, err
	}

	//判断是否限流
	if err := m.CheckRateLimit(ctx.ClientIP(), interfaceInfo); err != nil {
		return nil, err
	}

	//判断是否使用token验证
	if interfaceInfo.TokenValidationEnabled == 1 {
		if err := m.CheckUserLogin(ctx); err != nil {
			return nil, err
		}
	}

	//验证参数信息
	params, err := m.CheckParams(ctx, interfaceInfo)
	if err != nil {
		return nil, err
	}

	//判断是否开启缓存
	cacheKey := m.GetCacheKeyByParams(ctx, interfaceInfo)
	if interfaceInfo.CacheEnabled == 1 {
		_CacheData, IsExists := libraries.GetCache(cacheKey)
		if IsExists { //存在缓存 从缓存中获取
			return _CacheData, nil
		}
	}
	Api_Model.InitApiDB()
	data, err := Api_Model.Run(interfaceInfo, params)
	if err != nil {
		return nil, err
	}

	if interfaceInfo.ReturnType == "insert" ||
		interfaceInfo.ReturnType == "update" ||
		interfaceInfo.ReturnType == "delete" {
		//这几种类型不错缓存
	} else {
		if interfaceInfo.CacheEnabled == 1 { //开启缓存 写入缓存
			libraries.AddCache(cacheKey, data, time.Duration(interfaceInfo.CacheTime)*time.Second)
		}
	}

	return data, nil
}

// 限流验证
func (m *Api) CheckRateLimit(ip string, interfaceInfo structs.Interface) error {
	if interfaceInfo.RateLimitEnabled == 1 {
		cacheKey := "RateLimit_" + ip + "_" + interfaceInfo.Path + "_" + interfaceInfo.Method
		// 获取当前请求次数
		count, _ := libraries.GetCache(cacheKey)
		if count == nil {
			count = 0
		}

		// 检查是否超过限流次数
		if count.(int) >= interfaceInfo.RateLimitCount {
			return cron.CreateCustomError(429, "请求次数超过限制，请稍后再试")
		}

		// 更新请求次数
		libraries.AddCache(cacheKey, count.(int)+1, time.Duration(interfaceInfo.RateLimitTime)*time.Second)
	}

	return nil
}

func (m *Api) CheckUserLogin(ctx *gin.Context) error {
	token := ctx.GetHeader("usertoken")
	if token == "" {
		return cron.CreateCustomError(500, "token 未提交")
	}

	tokenModel := Token{}
	userId, err := tokenModel.GetTokenInfo(token)
	if err != nil {
		return err
	}

	//token 有效时间延续2小时
	tokenModel.TokenExtendTime(userId, token)
	ctx.Set("USERID", userId)
	return nil
}

// 获得参数拼接的字符
func (m *Api) GetCacheKeyByParams(ctx *gin.Context, Interface structs.Interface) string {

	var paramsText string
	if Interface.Method == "get" {
		for _, param := range Interface.Params {
			if param.Default != "" {
				paramsText += param.Name + "=" + ctx.DefaultQuery(param.Name, param.Default) + "&"
			} else {
				paramsText += param.Name + "=" + ctx.Query(param.Name) + "&"
			}

		}
	} else if Interface.Method == "post" {
		for _, param := range Interface.Params {
			if param.Default != "" {
				paramsText += param.Name + "=" + ctx.DefaultPostForm(param.Name, param.Default) + "&"
			} else {
				paramsText += param.Name + "=" + ctx.PostForm(param.Name) + "&"
			}

		}
	}

	if len(paramsText) > 0 {
		paramsText = paramsText[:len(paramsText)-1] // Remove the trailing '&'
	}

	fmt.Println("cacheKey:", Interface.Path+Interface.Method+paramsText)
	return "Interface_Data_" + helper.HashMD5(Interface.Path+Interface.Method+paramsText)
}

// 验证get/post参数
func (m *Api) CheckParams(ctx *gin.Context, Interface structs.Interface) (map[string]string, error) {
	var paramsData = make(map[string]string)

	for _, paramItem := range Interface.Params {
		var _param string

		if Interface.Method == "post" {
			if paramItem.Default != "" {
				_param = ctx.DefaultPostForm(paramItem.Name, paramItem.Default)
			} else {
				_param = ctx.PostForm(paramItem.Name)
			}
		} else if Interface.Method == "get" {
			if paramItem.Default != "" {
				_param = ctx.DefaultQuery(paramItem.Name, paramItem.Default)
			} else {
				_param = ctx.Query(paramItem.Name)
			}
		}

		//判断是否必填
		if paramItem.Required == 1 && _param == "" {
			return nil, cron.CreateCustomError(500, "参数["+paramItem.Name+"]不能为空")
		}

		//正则表达式验证
		if paramItem.Type == "string" && paramItem.Regex != "" {
			match := helper.MatchString(paramItem.Regex, _param)
			if !match {
				return nil, cron.CreateCustomError(500, "参数["+paramItem.Name+"]格式不正确")
			}
		} else if paramItem.Type == "int" {
			_, err := strconv.Atoi(_param)
			if err != nil {
				return nil, cron.CreateCustomError(500, "参数["+paramItem.Name+"]必须是整数")
			}
		} else if paramItem.Type == "float" {
			_, err := strconv.ParseFloat(_param, 64)
			if err != nil {
				return nil, cron.CreateCustomError(500, "参数["+paramItem.Name+"]必须是浮点数")
			}
		} else if paramItem.Type == "bool" {
			if _param != "1" && _param != "2" {
				return nil, cron.CreateCustomError(500, "参数["+paramItem.Name+"]必须是(1是, 2否)")
			}
		} else if paramItem.Type == "date" {
			_, err := time.Parse("2006-01-02", _param)
			if err != nil {
				return nil, cron.CreateCustomError(500, "参数["+paramItem.Name+"]日期格式必须是(YYYY-MM-DD)")
			}
		} else if paramItem.Type == "datetime" {
			_, err := time.Parse("2006-01-02 15:04:05", _param)
			if err != nil {
				return nil, cron.CreateCustomError(500, "参数["+paramItem.Name+"]日期时间格式必须是(YYYY-MM-DD HH:mm:ss)")
			}
		}

		paramsData[paramItem.Name] = _param
	}
	return paramsData, nil
}

// 判断是否是白名单或者黑名单内的ip
func (m *Api) CheckIp(ctx *gin.Context) error {
	//获得名单类型
	var ConfigModel = Config{}
	var WhiteListModel = WhiteList{}
	ListType := ConfigModel.GetBlackListType()

	//获得当前ip
	var CurrentIp = ctx.ClientIP()

	if ListType == 0 {
		return nil
	} else if ListType == 1 {
		//黑名单

		BlackList, err := WhiteListModel.GetAllBlackList()
		if err != nil {
			return cron.CreateCustomError(500, "获取黑名单失败")
		}

		for _, ip := range BlackList {
			if ip.IP == CurrentIp {
				return cron.CreateCustomError(500, "当前ip已被禁止访问")
			}
		}

	} else if ListType == 2 {
		//白名单
		var WhiteListModel = WhiteList{}

		WhiteList, err := WhiteListModel.GetAllWhiteList()
		if err != nil {
			return cron.CreateCustomError(500, "获取白名单失败")
		}

		for _, ip := range WhiteList {
			if ip.IP == CurrentIp {
				return nil
			}
		}

		return cron.CreateCustomError(500, "当前ip不允许访问")
	}

	return nil
}
