package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/database"
	"GoEasyApi/helper"
	"GoEasyApi/libraries"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Api struct{}

var InterfaceModel = Interface{}
var Api_Model = ApiModel{}

func (m *Api) Get(ctx *gin.Context) (interface{}, error) {
	//获得接口信息
	interfaceInfo, err := InterfaceModel.GetInfoByPath(ctx.Request.URL.Path)
	if err != nil {
		return nil, err
	}

	//判断是否使用token验证
	if interfaceInfo.TokenValidationEnabled == 1 {
		if err := m.CheckUserLogin(ctx); err != nil {
			return nil, err
		}
	}

	//验证参数信息
	params, err := m.CheckParams(ctx, interfaceInfo, "get")
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

	Api_Model.Init()
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
	//获得接口信息
	interfaceInfo, err := InterfaceModel.GetInfoByPath(ctx.Request.URL.Path)
	if err != nil {
		return nil, err
	}

	//判断是否使用token验证
	if interfaceInfo.TokenValidationEnabled == 1 {
		if err := m.CheckUserLogin(ctx); err != nil {
			return nil, err
		}
	}

	//验证参数信息
	params, err := m.CheckParams(ctx, interfaceInfo, "post")
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

	Api_Model.Init()
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

func (m *Api) CheckUserLogin(ctx *gin.Context) error {
	token := ctx.GetHeader("usertoken")
	if token == "" {
		return cron.CreateCustomError(500, "oken 未提交")
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
func (m *Api) GetCacheKeyByParams(ctx *gin.Context, Interface database.Interface) string {

	var paramsText string
	for _, param := range Interface.Params {
		if param.Default != "" {
			paramsText += param.Name + "=" + ctx.DefaultQuery(param.Name, param.Default) + "&"
		} else {
			paramsText += param.Name + "=" + ctx.Query(param.Name) + "&"
		}

	}
	if len(paramsText) > 0 {
		paramsText = paramsText[:len(paramsText)-1] // Remove the trailing '&'
	}

	return "Interface_Data_" + helper.HashMD5(paramsText)
}

// 验证get/post参数
func (m *Api) CheckParams(ctx *gin.Context, Interface database.Interface, method string) (map[string]string, error) {
	var paramsData = make(map[string]string)

	for _, paramItem := range Interface.Params {
		var _param string

		if method == "post" {
			if paramItem.Default != "" {
				_param = ctx.DefaultPostForm(paramItem.Name, paramItem.Default)
			} else {
				_param = ctx.PostForm(paramItem.Name)
			}
		} else {
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
