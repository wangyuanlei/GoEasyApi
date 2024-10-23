package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/database"
	"GoEasyApi/helper"
	"GoEasyApi/libraries"
	"strconv"
	"strings"
	"time"

	"github.com/cengsin/oracle"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Api struct{}

var InterfaceModel = Interface{}
var UserDB *gorm.DB

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
	params, err := m.CheckGetParams(ctx, interfaceInfo)
	if err != nil {
		return nil, err
	}

	//判断是否开启缓存
	if interfaceInfo.CacheEnabled == 1 {
		cacheKey := m.GetCacheKeyByParams(ctx, interfaceInfo)
		_CacheData, IsExists := libraries.GetCache(cacheKey)
		if IsExists { //存在缓存 从缓存中获取
			return _CacheData, nil
		}
	}

	//获得要执行的sql
	sql, err := m.GetInterfaceSql(params, interfaceInfo.SqlContent)
	if err != nil {
		return nil, err
	}

	if interfaceInfo.CacheEnabled == 1 { //开启缓存 写入缓存

	}

	return interfaceInfo, nil

}

// func (m *Api) Post(ctx *gin.Context) (interface{}, error) {
// 	path := ctx.Request.URL.Path
// }

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

// 验证参数
func (m *Api) CheckGetParams(ctx *gin.Context, Interface database.Interface) (map[string]string, error) {
	var paramsData = make(map[string]string)

	for _, paramItem := range Interface.Params {
		var _param string

		if paramItem.Default != "" {
			_param = ctx.DefaultQuery(paramItem.Name, paramItem.Default)
		} else {
			_param = ctx.Query(paramItem.Name)
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

// 根据参数 和 接口sql 获得sql 语句
func (m *Api) GetInterfaceSql(params map[string]string, sql_content string) (string, error) {

	//根据params 组织sql 语句
	for k, v := range params {
		sql_content = strings.Replace(sql_content, "{{"+k+"}}", v, -1)
	}

	return sql_content, nil
}

// 获得数据库连接
func (m *Api) InitUserDB() (*gorm.DB, error) {
	var DBModel = DataBase{}
	//获得数据库连接配置
	dbConfig, err := DBModel.GetUserDBConf()
	if err != nil {
		return nil, err
	}

	//获得数据库连接句柄
	if dbConfig.OrmType == "mysql" {
		//dsn := "username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		UserDB, err = gorm.Open(mysql.Open(dbConfig.Dns), &gorm.Config{})
		if err != nil {
			panic("failed to connect database:")
		}
	} else if dbConfig.OrmType == "postgresql" {
		//dsn := "host=localhost user=your_username password=your_password dbname=your_db port=5432 sslmode=disable"
		UserDB, err = gorm.Open(postgres.Open(dbConfig.Dns), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	} else if dbConfig.OrmType == "sqlite" {
		//dsn := "test.db"
		UserDB, err = gorm.Open(sqlite.Open(dbConfig.Dns), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	} else if dbConfig.OrmType == "sqlserver" {
		//dsn := "sqlserver://username:password@localhost:1433?database=dbname"
		UserDB, err = gorm.Open(sqlserver.Open(dbConfig.Dns), &gorm.Config{})
		if err != nil {
			panic("failed to connect to database")
		}
	} else if dbConfig.OrmType == "oracle" {
		//dsn := "system/oracle@127.0.0.1:1521/XE"
		UserDB, err = gorm.Open(oracle.Open(dbConfig.Dns), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	} else {
		return nil, cron.CreateCustomError(602, "数据库类型错误")
	}

	return UserDB, nil

}

func (m *Api) run(ctx *gin.Context, Interface database.Interface) (interface{}, error) {

}
