package model

import (
	"GoEasyApi/database"
	"GoEasyApi/helper"
	"GoEasyApi/libraries"
	"time"
)

type User struct{}

// RegisterUser 注册新用户
func (u *User) RegisterUser(name string, account string, password string, deptId string) error {
	var existingUser database.User
	if err := DB.Where("account = ?", account).First(&existingUser).Error; err == nil {
		return libraries.CreateCustomError(601, "账号已存在")
	}

	//生成随机Salt
	salt := helper.GenerateRandomString(6)
	// 生成密码
	hashedPassword := helper.DoubleHashPassword(password, salt)

	return DB.Create(&database.User{
		Name:         name,
		Account:      account,
		Password:     hashedPassword,
		Salt:         salt,
		DeptId:       deptId,
		RegisterTime: time.Now(),
		IsValid:      1,
	}).Error
}

// UpdateUserInfo 更新用户信息
func (u *User) UpdateUserInfo(userId string, name string, deptId string) error {
	var user database.User
	if err := DB.First(&user, "user_id = ?", userId).Error; err != nil {
		return libraries.CreateCustomError(602, "用户信息不存在")
	}
	return DB.Model(&user).Updates(database.User{
		Name:   name,
		DeptId: deptId,
	}).Error
}

// UpdateUserPassword 更新用户密码
func (u *User) UpdateUserPassword(userId string, oldPassword string, newPassword string) error {
	var user database.User
	if err := DB.First(&user, "user_id = ?", userId).Error; err != nil {
		return libraries.CreateCustomError(602, "用户信息不存在")
	}

	if helper.DoubleHashPassword(oldPassword, user.Salt) != user.Password {
		return libraries.CreateCustomError(602, "旧密码不正确")
	}

	hashedNewPassword := helper.DoubleHashPassword(newPassword, user.Salt)
	return DB.Model(&user).Update("password", hashedNewPassword).Error
}

// SetUserValidity 设置用户有效性
func (u *User) SetUserValidity(userId string, isValid bool) error {
	return DB.Model(&User{}).Where("user_id = ?", userId).Update("is_valid", isValid).Error
}

// GetCurrentUserInfo 获取当前用户信息
func (u *User) GetCurrentUserInfo(userId string) (database.User, error) {
	var user database.User
	if err := DB.First(&user, "user_id = ?", userId).Error; err != nil {
		return database.User{}, err
	}
	return user, nil
}

// GetUserList 获取用户列表带分页
func (u *User) GetUserList(page int, pageSize int, deptId string, name string, isValid int) ([]database.User, int64, error) {
	var users []database.User
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
	return users, total, nil
}
