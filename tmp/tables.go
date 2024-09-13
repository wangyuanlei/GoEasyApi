/*
表: client
字段:

	client_id 客户端id
	client_name 客户端名称
	token_validation token验证配置
	encryption 出参入参验证配置
	redis_cache redis缓存配置
	is_open_white_list 是否开启白名单 或者 黑名单
	white_list 白名单
	black_list 黑名单
	certificates https证书配置
*/
package database

type Client struct {
	ClientID        string          `gorm:"primary_key;type:varchar(100)"` // 客户端id
	ClientName      string          `gorm:"type:varchar(100)"`             // 客户端名称
	TokenValidation TokenValidation `gorm:"embedded"`                      // token验证配置
	Encryption      Encryption      `gorm:"embedded"`                      // 出参入参验证配置
	IsOpenWhiteList bool            `gorm:"default:false"`                 // 是否开启白名单 或者 黑名单
	WhiteList       []WhiteList     `gorm:"foreignkey:ClientID"`           // 白名单
	BlackList       []BlackList     `gorm:"foreignkey:ClientID"`           // 黑名单
	Certificates    []Certificate   `gorm:"foreignkey:ClientID"`           // https证书配置
}

type TokenValidation struct {
	Enabled     bool   `gorm:"default:false"`    // 是否启用token验证
	DataSource  string `gorm:"type:varchar(50)"` // 数据源名称
	TableName   string `gorm:"type:varchar(50)"` // 对应的表名
	TokenField  string `gorm:"type:varchar(50)"` // token对应的字段名
	UserIDField string `gorm:"type:varchar(50)"` // 对应的user_id字段名
}

type Encryption struct {
	InputEnabled         bool   `gorm:"default:false"`    // 是否启用入参私钥加密
	OutputEnabled        bool   `gorm:"default:false"`    // 是否启用出参私钥加密
	Algorithm            string `gorm:"type:varchar(50)"` // 使用的加密算法
	KeySize              int    `gorm:"type:int"`         // 密钥长度
	InputPublicKeyPath   string `gorm:"type:text"`        // 入参公钥验证文件的路径
	OutputPrivateKeyPath string `gorm:"type:text"`        // 出参私钥加密文件的路径
}

type RedisCache struct {
	Enabled   bool   `gorm:"default:false"`    // 是否启用Redis缓存
	Host      string `gorm:"type:varchar(50)"` // Redis服务器地址
	Port      int    `gorm:"type:int"`         // Redis服务器端口
	Password  string `gorm:"type:text"`        // Redis密码，如果有的话
	DB        int    `gorm:"type:int"`         // 使用的Redis数据库编号
	KeyPrefix string `gorm:"type:varchar(50)"` // 缓存键前缀
}

type WhiteList struct {
	IP          string `gorm:"type:varchar(15)"` // 允许访问的IP地址
	Description string `gorm:"type:text"`        // IP地址描述或备注
}

type BlackList struct {
	IP          string `gorm:"type:varchar(15)"` // 禁止访问的IP地址
	Description string `gorm:"type:text"`        // IP地址描述或备注
}

type Certificate struct {
	Cert        string `gorm:"type:text"` // 客户端证书内容
	Key         string `gorm:"type:text"` // 客户端私钥内容
	Description string `gorm:"type:text"` // 证书描述或备注
}

/*
	表: database
	字段:
		database_id 数据库源id
		description 数据源描述
		type 数据库类型 mysql/postgresql/oracle/sqlserver/sqlite
		host 数据库地址
		port 数据库端口
		username 数据库用户名
		password 数据库密码
		database_name 数据库名称
		ssl_mode 是否为ssl模式
			ssl_cert 证书路径
			ssl_key 私钥路径
			ssl_ca 证书授权路径
			ssl_capath 证书授权路径
			ssl_cipher 加密方式
			ssl_key_password 私钥密码
*/

type Database struct {
	DatabaseCode   string `gorm:"type:varchar(36);primary_key"` // 数据库源id
	Description    string `gorm:"type:text"`                    // 数据源描述
	Type           string `gorm:"type:varchar(20)"`             // 数据库类型 mysql/postgresql/oracle/sqlserver/sqlite
	Host           string `gorm:"type:varchar(255)"`            // 数据库地址
	Port           int    `gorm:"type:int"`                     // 数据库端口
	Username       string `gorm:"type:varchar(50)"`             // 数据库用户名
	Password       string `gorm:"type:text"`                    // 数据库密码
	DatabaseName   string `gorm:"type:varchar(50)"`             // 数据库名称
	SslMode        bool   `gorm:"type:bool"`                    // 是否为ssl模式
	SslCert        string `gorm:"type:text"`                    // 证书路径
	SslKey         string `gorm:"type:text"`                    // 私钥路径
	SslCa          string `gorm:"type:text"`                    // 证书授权路径
	SslCapath      string `gorm:"type:text"`                    // 证书授权路径
	SslCipher      string `gorm:"type:text"`                    // 加密方式
	SslKeyPassword string `gorm:"type:text"`                    // 私钥密码
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
	InterfaceId      string  `gorm:"type:varchar(36);primary_key"`                                  // 接口编号
	InterfaceName    string  `gorm:"type:varchar(50)"`                                              // 接口名称
	Description      string  `gorm:"type:text"`                                                     // 接口描述
	DatabaseId       string  `gorm:"type:varchar(36)"`                                              // 数据库源id
	ClientId         string  `gorm:"type:varchar(36)"`                                              // 客户端id
	Path             string  `gorm:"type:varchar(255)"`                                             // 接口路径
	Method           string  `gorm:"type:varchar(10)"`                                              // 接口方法
	CacheEnabled     bool    `gorm:"type:bool"`                                                     // 是否启用接口缓存
	CacheTime        int     `gorm:"type:int"`                                                      // 接口缓存时间
	RateLimitEnabled bool    `gorm:"type:bool"`                                                     // 是否启用接口限流
	RateLimitCount   int     `gorm:"type:int"`                                                      // 接口限流次数
	RateLimitTime    int     `gorm:"type:int"`                                                      // 接口限流时间
	SqlContent       string  `gorm:"type:text"`                                                     // 接口sql语句
	ReturnType       string  `gorm:"type:varchar(50)"`                                              // 接口返回类型
	Params           []Param `gorm:"foreignkey:InterfaceCode;association_foreignkey:InterfaceCode"` // 接口参数

}

type Param struct {
	InterfaceId string `gorm:"type:varchar(36)"` // 接口编号
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
		clientId 客户端id
		type 用户类型 1:超级管理员 2:普通用户
*/
type User struct {
	UserId      string `gorm:"type:varchar(36);primary_key"` // 用户id
	Username    string `gorm:"type:varchar(50)"`             // 用户名
	Password    string `gorm:"type:varchar(255)"`            // 用户密码
	Salt        string `gorm:"type:varchar(255)"`            // 盐值
	Description string `gorm:"type:text"`                    // 用户描述
	ClientId    string `gorm:"type:varchar(36)"`             // 客户端id
	Type        int    `gorm:"type:int"`                     // 用户类型 1:超级管理员 2:普通用户
}
