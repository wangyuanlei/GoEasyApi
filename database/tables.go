package database

import (
	"time"
)

type WhiteList struct {
	IP          string `gorm:"type:varchar(15)"` // 允许访问的IP地址
	Description string `gorm:"type:text"`        // IP地址描述或备注
}

type BlackList struct {
	IP          string `gorm:"type:varchar(15)"` // 禁止访问的IP地址
	Description string `gorm:"type:text"`        // IP地址描述或备注
}

/*
	表: database
	字段:
		database_id 数据库源id
		description 数据源描述
		orm_type	数据库连接类型
		type 数据库类型 mysql/postgresql/oracle/sqlserver/sqlite
		host 数据库地址
		port 数据库端口
		username 数据库用户名
		password 数据库密码
		database_name 数据库名称
*/

type Database struct {
	DatabaseCode string `gorm:"type:varchar(32);primary_key"` // 数据库源id
	Description  string `gorm:"type:text"`                    // 数据源描述
	OrmType      string `gorm:"type:varchar(20)"`             // 数据库连接类型
	Type         string `gorm:"type:varchar(20)"`             // 数据库类型 mysql/postgresql/oracle/sqlserver/sqlite
	Host         string `gorm:"type:varchar(255)"`            // 数据库地址
	Port         int    `gorm:"type:int"`                     // 数据库端口
	Username     string `gorm:"type:varchar(50)"`             // 数据库用户名
	Password     string `gorm:"type:text"`                    // 数据库密码
	DatabaseName string `gorm:"type:varchar(50)"`             // 数据库名称
}

/*
表: interface
字段:

	interface_id 接口编号
	interface_name 接口名称
	database_id 数据库源id
	client_id 客户端id
	description 接口描述
	path 接口路径，例如 /api/users
	method 接口方法，例如 GET/POST
	cache_enabled 是否启用接口缓存
	cache_time 接口缓存时间，单位秒
	rate_limit_enabled 是否启用接口限流
	rate_limit_count 接口限流次数
	rate_limit_time 接口限流时间，单位秒
	sql_content 接口sql语句
	token_validation_enabled 是否启用token验证
	return_type 接口返回类型
			json对象 比如{"test":1},
			json二维数组  比如[{"test":1}],
			分页内容  比如{"total":100,"list":[{"test":1}]},
			布尔类型  比如true, false,
			新增id 比如1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	params 接口参数
		name 参数名称
		type 参数类型
		description 参数描述
		required 是否必传
		default 默认值
		example 示例值
		regex 正则表达式
*/
type Interface struct {
	InterfaceId            string `gorm:"type:varchar(32);primary_key"` // 接口编号
	InterfaceName          string `gorm:"type:varchar(50)"`             // 接口名称
	Description            string `gorm:"type:text"`                    // 接口描述
	DatabaseId             string `gorm:"type:varchar(32)"`             // 数据库源id
	ClientId               string `gorm:"type:varchar(32)"`             // 客户端id
	Path                   string `gorm:"type:varchar(255)"`            // 接口路径
	Method                 string `gorm:"type:varchar(10)"`             // 接口方法
	CacheEnabled           bool   `gorm:"type:bool"`                    // 是否启用接口缓存
	CacheTime              int    `gorm:"type:int"`                     // 接口缓存时间
	RateLimitEnabled       bool   `gorm:"type:bool"`                    // 是否启用接口限流
	RateLimitCount         int    `gorm:"type:int"`                     // 接口限流次数
	RateLimitTime          int    `gorm:"type:int"`                     // 接口限流时间
	SqlContent             string `gorm:"type:text"`                    // 接口sql语句
	TokenValidationEnabled bool   `gorm:"type:bool"`                    // 是否启用token验证
	ReturnType             string `gorm:"type:varchar(50)"`             // 接口返回类型
}

type Params struct {
	InterfaceId string `gorm:"type:varchar(32)"` // 接口编号
	Name        string `gorm:"type:varchar(50)"` // 参数名称
	Type        string `gorm:"type:varchar(20)"` // 参数类型 比如 string/int/float/bool/datetime
	Description string `gorm:"type:text"`        // 参数描述
	Required    bool   `gorm:"type:bool"`        // 是否必传
	Default     string `gorm:"type:text"`        // 默认值
	Example     string `gorm:"type:text"`        // 示例值
	Regex       string `gorm:"type:text"`        // 正则表达式
}

/*
表: user
字段:

	user_id 用户id
	username 用户名
	password 用户密码
	salt 盐值
	description 用户描述
	type 用户类型 1:超级管理员 2:普通用户
*/
type User struct {
	UserId      string `gorm:"type:varchar(32);primary_key"` // 用户id
	Username    string `gorm:"type:varchar(50)"`             // 用户名
	Password    string `gorm:"type:varchar(255)"`            // 用户密码
	Salt        string `gorm:"type:varchar(255)"`            // 盐值
	Description string `gorm:"type:text"`                    // 用户描述
	Type        int    `gorm:"type:int"`                     // 用户类型 1:超级管理员 2:普通用户
}

/*
表: token
字段:

	token 登录验证token
	user_id 用户id
	valid_time 有效时间
*/
type Token struct {
	Token     string    `gorm:"type:varchar(100);primary_key"` // 登录验证token
	UserId    string    `gorm:"type:varchar(32)"`              // 用户id
	ValidTime time.Time `gorm:"type:datetime"`                 // 有效时间
}
