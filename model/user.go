package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/helper"
	"GoEasyApi/structs"
	"time"

	"github.com/google/uuid"
)

type User struct{}

// RegisterUser 注册新用户
func (u *User) RegisterUser(name string, account string, password string, deptId string) error {
	var existingUser structs.User
	if err := DB.Where("account = ?", account).First(&existingUser).Error; err == nil {
		return cron.CreateCustomError(601, "账号已存在")
	}

	userId := uuid.New().String()
	var existingUser2 structs.User
	if err := DB.Where("user_id = ?", userId).First(&existingUser2).Error; err == nil {
		return cron.CreateCustomError(601, "用户ID已存在")
	}

	//生成随机Salt
	salt := helper.GenerateRandomString(6)
	// 生成密码
	hashedPassword := helper.DoubleHashPassword(password, salt)

	return DB.Create(&structs.User{
		UserId:       userId,
		Name:         name,
		Account:      account,
		Password:     hashedPassword,
		Salt:         salt,
		DeptId:       deptId,
		RegisterTime: time.Now(),
		IsValid:      2,
	}).Error
}

// 用户登录
func (u *User) Login(account string, password string) (*structs.User, error) {
	var user structs.User
	if err := DB.Where("account = ?", account).First(&user).Error; err != nil {
		return nil, cron.CreateCustomError(602, "用户不存在")
	}

	if helper.DoubleHashPassword(password, user.Salt) != user.Password {
		return nil, cron.CreateCustomError(602, "密码错误")
	}

	//密码和加密salt 信息 不能输出到前台
	user.Password = ""
	user.Salt = ""

	return &user, nil
}

// ChangeInfo 更新用户信息
func (u *User) ChangeInfo(userId string, name string, deptId string, isValid int) error {
	var user structs.User
	if err := DB.First(&user, "user_id = ?", userId).Error; err != nil {
		return cron.CreateCustomError(602, "用户信息不存在")
	}

	if isValid != 1 && isValid != 2 {
		return cron.CreateCustomError(602, "无效参数")
	}

	return DB.Model(&user).Updates(structs.User{
		Name:    name,
		DeptId:  deptId,
		IsValid: isValid,
	}).Error
}

// UpdateUserPassword 更新用户密码
func (u *User) ChangePassword(userId string, oldPassword string, newPassword string) error {
	var user structs.User
	if err := DB.First(&user, "user_id = ?", userId).Error; err != nil {
		return cron.CreateCustomError(602, "用户信息不存在")
	}

	if helper.DoubleHashPassword(oldPassword, user.Salt) != user.Password {
		return cron.CreateCustomError(602, "旧密码不正确")
	}

	hashedNewPassword := helper.DoubleHashPassword(newPassword, user.Salt)
	return DB.Model(&user).Update("password", hashedNewPassword).Error
}

// UpdateUserPassword 更新用户密码
func (u *User) AdminChangePassword(userId string, password string) error {
	var user structs.User
	if err := DB.First(&user, "user_id = ?", userId).Error; err != nil {
		return cron.CreateCustomError(602, "用户信息不存在")
	}

	hashedNewPassword := helper.DoubleHashPassword(password, user.Salt)
	return DB.Model(&user).Update("password", hashedNewPassword).Error
}

// SetUserValidity 设置用户有效性
func (u *User) SetUserValidity(userId string, isValid int) error {
	if isValid != 1 && isValid != 2 {
		return cron.CreateCustomError(602, "无效参数")
	}

	return DB.Model(&User{}).Where("user_id = ?", userId).Update("is_valid", isValid).Error
}

// GetCurrentUserInfo 获取当前用户信息
func (u *User) GetCurrentUserInfo(userId string) (structs.User, error) {
	var user structs.User
	if err := DB.First(&user, "user_id = ?", userId).Error; err != nil {
		return structs.User{}, err
	}

	//密码和加密salt 信息 不能输出到前台
	user.Password = ""
	user.Salt = ""
	return user, nil
}

// GetUserList 获取用户列表带分页
func (u *User) GetUserList(page int, pageSize int, deptId string, name string, isValid int) ([]structs.User, int64, error) {
	var users []structs.User
	var total int64
	query := DB
	if deptId != "" {
		query = query.Where("dept_id = ?", deptId)
	}
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if isValid == 1 || isValid == 2 {
		query = query.Where("is_valid = ?", isValid)
	}
	if err := query.Find(&users).Offset((page - 1) * pageSize).Limit(pageSize).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	//密码和加密salt 信息 不能输出到前台
	for i := range users {
		users[i].Password = ""
		users[i].Salt = ""
	}

	return users, total, nil
}

// 删除用户信息
func (u *User) DeleteUser(userId string) error {
	return DB.Where("user_id = ?", userId).Delete(&User{}).Error
}
