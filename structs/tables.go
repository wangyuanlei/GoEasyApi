package structs

/*数据库结构表*/

import "time"

//白名单数据结构
type WhiteList struct {
	IP          string `gorm:"type:varchar(15)"` // 允许访问的IP地址
	Description string `gorm:"type:text"`        // IP地址描述或备注
}

//黑名单数据结构
type BlackList struct {
	IP          string `gorm:"type:varchar(15)"` // 禁止访问的IP地址
	Description string `gorm:"type:text"`        // IP地址描述或备注
}

//数据库配置数据结构
type Database struct {
	DatabaseId   string `gorm:"type:varchar(32);primary_key"` // 数据库源id uuid编号
	DatabaseName string `gorm:"type:varchar(32)"`             // 数据库源名称
	Description  string `gorm:"type:text"`                    // 数据源描述
	OrmType      string `gorm:"type:varchar(20)"`             // 数据库连接类型 mysql/postgresql/oracle/sqlserver/sqlite
	Dns          string `gorm:"type:varchar(255)"`            // 数据库地址
}

//接口信息数据结构
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
}

/*
接口参数数据结构
Type: 定义
	string 字符串类型. 支持正则校验
	int 强转换成整型
	float 强转换成浮点类型
	bool 强制设置成 1 是 2 否
	date 日期类型 YYYY-mm-dd
	datetime 日期时间类型 YYYY-mm-dd HH:ii:ss
*/
type Params struct {
	ParamsId    string `gorm:"type:varchar(32);primary_key"` // 接口编号
	InterfaceId string `gorm:"type:varchar(32)"`             // 接口编号
	Name        string `gorm:"type:varchar(50)"`             // 参数名称
	Type        string `gorm:"type:varchar(20)"`             // 参数类型 比如 string/int/float/bool/date/datetime
	Description string `gorm:"type:text"`                    // 参数描述
	Required    int    `gorm:"type:int"`                     // 是否必传
	Default     string `gorm:"type:text"`                    // 默认值
	Example     string `gorm:"type:text"`                    // 示例值
	Regex       string `gorm:"type:text"`                    // 正则表达式
}

//用户数据表结构
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

//用户token表结构
type Token struct {
	Token     string    `gorm:"type:varchar(100);primary_key"` // 登录验证token
	UserId    string    `gorm:"type:varchar(32)"`              // 用户id
	ValidTime time.Time `gorm:"type:datetime"`                 // 有效时间
}
