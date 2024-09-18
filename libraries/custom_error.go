package libraries

/*
错误类 CustomError, 继承系统error
下面是错误编码定义列表
500: 账号和密码错误
501: 验证失败,请重新登录
502: token 未提交
601: 自定义错误内容的错误
*/
type CustomErrorNew struct { // Renamed struct
	Code    int
	Message string
}

func (e *CustomErrorNew) Error() string {
	return e.Message
}

func CreateCustomError(code int, message string) *CustomErrorNew {
	return &CustomErrorNew{
		Code:    code,
		Message: message,
	}
}
