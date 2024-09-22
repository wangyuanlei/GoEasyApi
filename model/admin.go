package model

/*
当前项目名称: GoEasyApi
写管理员操作类.使用中文注释. 并且写上实现过程
libraries/cache.go 是 缓存类
libraries/custom_error.go 错误类
libraries/config.go 配置读取类
需要实现以下功能:
	1. 登录验证功能.
		入参: username , password
		实现过程:
			1. 从配置 获得 	username 和 password
			2. pass = md5(md5(配置password) + 'Qwert!@#456')
			3. 入参 username 和 配置 username 对比, pass 和 配置 password 对比
			4. 如果对比失败. 使用 custom_error 错误类报错
			5.如果对比成功, 生成一个 uuid 作为 token 返回
			6. uuid 存储到 cache类



	2. 验证登录
		入参: token
		实现过程:
			1. 验证 token 是否在 cache 缓存中存在, 不存在, 验证失败.
			2. 如果验证成功. 有效期改为当前时间 + 2小时
*/
import (
	"GoEasyApi/helper"
	"GoEasyApi/libraries"
	"time"

	"github.com/google/uuid"
)

type Admin struct{}

// 登录验证功能
// Login 实现了Admin结构体的登录方法
//
// 参数：
//
//	username string - 登录用户名
//	password string - 登录密码
//
// 返回值：
//
//	string - 登录成功后返回的token
//	error - 登录失败时返回的错误信息
func (a *Admin) Login(username, password string) (string, error) {
	// 从配置获得 username 和 password
	configUsername, configPassword, err := libraries.LoadUserConfig()
	if err != nil {
		return "", libraries.CreateCustomError(601, "配置错误") // 更新为直接从 libraries 调用 NewCustomError
	}

	pass := helper.HashPassword(password)
	// fmt.Print("pass: ", pass)
	// 对比 username 和 password
	if username != configUsername || pass != configPassword {
		return "", libraries.CreateCustomError(500, "账号或者密码错误")
	}

	// 生成 uuid 作为 token
	token := uuid.New().String()

	// uuid 存储到 cache 类
	libraries.AddCache(token, username, 2*time.Hour)

	return token, nil
}

// 验证登录
// ValidateLogin 验证登录 token 是否有效
//
// 参数：
//
//	a *Admin - Admin 结构体指针
//	token string - 登录 token
//
// 返回值：
//
//	string - 用户名，如果验证失败则为空字符串
//	error - 错误信息，如果验证成功则为 nil
func (a *Admin) ValidateLogin(token string) (string, error) {
	// 验证 token 是否在 cache 缓存中存在
	_userName, IsExists := libraries.GetCache(token)
	if !IsExists {
		return "", libraries.CreateCustomError(601, "验证失败，token 不存在")
	}

	// 类型断言将 _userName 转换为 string
	username, ok := _userName.(string)
	if !ok {
		return "", libraries.CreateCustomError(500, "用户名类型错误")
	}

	// 如果验证成功，有效期改为当前时间 + 2小时
	libraries.UpdateCache(token, _userName, 2*time.Hour)

	return username, nil
}
