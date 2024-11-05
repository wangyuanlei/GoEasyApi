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
		database_name 数据库源名称
		description 数据源描述
		orm_type	数据库连接类型mysql/postgresql/oracle/sqlserver/sqlite
		dns 数据库地址
*/

type Database struct {
	DatabaseId   string `gorm:"type:varchar(32);primary_key"` // 数据库源id uuid编号
	DatabaseName string `gorm:"type:varchar(32)"`             // 数据库源名称
	Description  string `gorm:"type:text"`                    // 数据源描述
	OrmType      string `gorm:"type:varchar(20)"`             // 数据库连接类型 mysql/postgresql/sqlserver/sqlite
	Dns          string `gorm:"type:varchar(255)"`            // 数据库地址
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
	Path                   string `gorm:"type:varchar(255)"`            // 接口路径
	Method                 string `gorm:"type:varchar(10)"`             // 接口方法
	CacheEnabled           int    `gorm:"type:int"`                     // 是否启用接口缓存
	CacheTime              int    `gorm:"type:int"`                     // 接口缓存时间
	RateLimitEnabled       int    `gorm:"type:int"`                     // 是否启用接口限流
	RateLimitCount         int    `gorm:"type:int"`                     // 接口限流次数
	RateLimitTime          int    `gorm:"type:int"`                     // 接口限流时间
	SqlContent             string `gorm:"type:text"`                    // 接口sql语句
	TokenValidationEnabled int    `gorm:"type:int"`                     // 是否启用token验证
	ReturnType             string `gorm:"type:varchar(50)"`             // 接口返回类型
	ReturnValMode          string `gorm:"type:varchar(50)"`             // 接口返回模式
	Params                 []Params
}

/*
1. 数据添加类型端口: ReturnType:insert
   1. 对于mysql 可以返回自增id, 或者uuid. 	ReturnValMode:last_id
   2. 对于没有自增的id 可以指定uuid.		ReturnValMode:uuid
2. 数据修改/删除类型接口: ReturnType:update/delete
   1. 返回 bool 类型, 或者返回 修改后的 数据. ReturnValMode:row / bool
3. 数据单条查询接口: ReturnType:row
   1. 返回单条数据. 列子: {"name":"张三"}
4. 数据多条查询接口: ReturnType:list
   1. 返回多条数据. 列子: [{"name":"张三"},{"name":"李四"}]
*/

type Params struct {
	ParamsId    string `gorm:"type:varchar(32)"` // 接口编号
	InterfaceId string `gorm:"type:varchar(32)"` // 接口编号
	Name        string `gorm:"type:varchar(50)"` // 参数名称
	Type        string `gorm:"type:varchar(20)"` // 参数类型 比如 string/int/float/bool/date/datetime
	Description string `gorm:"type:text"`        // 参数描述
	Required    int    `gorm:"type:int"`         // 是否必传
	Default     string `gorm:"type:text"`        // 默认值
	Example     string `gorm:"type:text"`        // 示例值
	Regex       string `gorm:"type:text"`        // 正则表达式
}

/*
Type: 定义
	string 字符串类型. 支持正则校验
	int 强转换成整型
	float 强转换成浮点类型
	bool 强制设置成 1 是 2 否
	date 日期类型 YYYY-mm-dd
	datetime 日期时间类型 YYYY-mm-dd HH:ii:ss
*/

/*
sql: 系统变量
	_uuid  不需要输入的类型,  自动生成一个uuid值. 然后再sql 中使用
	_nowDate: 不需要输入的类型, 生成日期
	_nowDateTime: 不需要输入的类型: 自动生成日期时间
	_uid: 不需要输入的类型, 当前用户id
*/

/*
表: user
字段:

	user_id 用户id
	name 姓名
	account 账号
	password 用户密码
	dept_id 部门id
	salt 盐值
	register_time 注册时间
	is_valid 是否有效
*/
type User struct {
	UserId       string    `gorm:"type:varchar(32);primary_key"` // 用户id
	Name         string    `gorm:"type:varchar(50)"`             // 姓名
	Account      string    `gorm:"type:varchar(50)"`             // 账号
	Password     string    `gorm:"type:varchar(255)"`            // 用户密码
	DeptId       string    `gorm:"type:varchar(32)"`             // 部门id
	Salt         string    `gorm:"type:varchar(255)"`            // 盐值
	RegisterTime time.Time `gorm:"type:datetime"`                // 注册时间
	IsValid      int       `gorm:"type:int"`                     // 是否有效(1 是 2否)
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
