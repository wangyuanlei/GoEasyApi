package structs

/*外部参数输入结构*/
type AdminLoginParams struct {
	Account  string `json:"account"` //账号密码
	Password string `json:"pass"`    //密码
}

//黑名单类型结构
type BlackListTypeParams struct {
	BlackListType string `json:"black_list_type"`
}

//黑名单ip表结构
type BlackListIpParams struct {
	Ip          string `json:"ip"`
	Description string `json:"description"`
}

//白名单类型结构
type WhiteListTypeParams struct {
	WhiteListType string `json:"white_list_type"`
}

//白名单ip表结构
type WhiteListIpParams struct {
	Ip          string `json:"ip"`
	Description string `json:"description"`
}

//管理员密码修改结构
type AdminPasswordParams struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

//数据库创建结构
type CreateDatabaseParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OrmType     string `json:"orm_type"`
	Dns         string `json:"dns"`
}

//用户信息注册结构
type RegisterParams struct {
	Name     string `json:"name"`
	Account  string `json:"account"`
	Password string `json:"password"`
	DeptId   string `json:"deptId"`
}

//用户信息登录结构
type UserLoginParams struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

//用户信息修改结构
type UpdateUserParams struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	DeptId string `json:"dept_id"`
}

//用户信息修改密码结构
type UpdateUserPasswordParams struct {
	UserId      string `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

//管理员设置用户密码结构
type SetUserPasswordParams struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}
