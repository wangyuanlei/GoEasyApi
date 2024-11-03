package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/helper"
	"GoEasyApi/libraries"
	"fmt"
)

type Config struct{}

// 获得名单类型
// GetBlackListType 返回黑名单类型
//
// 该函数从配置中读取黑名单类型并返回。
//
// 如果加载白名单配置时发生错误，则返回0（或根据需要返回一个默认值或处理错误）。
func (c *Config) GetBlackListType() int {
	whitelistConfig, err := libraries.LoadWhitelistConfig()
	if err != nil {
		return 0 // 或者根据需要返回一个默认值或处理错误
	}
	return whitelistConfig
}

// 设置使用名单类型 0-不使用 1-使用黑名单 2-使用白名单
// SetBlackListType 用于设置黑名单类型
//
// 参数：
//
//	blackListType int - 黑名单类型
//
// 返回值：
//
//	error - 错误信息，如果成功则为nil
//
// 功能说明：
//
//	调用libraries.SaveWhitelistConfig函数，将传入的黑名单类型保存到配置文件中
//	如果保存成功，则返回nil；否则返回错误信息
func (c *Config) SetBlackListType(blackListType int) error {
	err := libraries.SaveWhitelistConfig(blackListType)
	if err != nil {
		return err
	}
	return nil
}

// 设置超级管理员密码
// SetSuperAdminPassword 函数用于设置超级管理员的密码
//
// 参数：
//
//	c *Config：Config 类型的指针，表示配置信息
//	oldpass string：旧密码，用于验证当前密码是否正确
//	newpass string：新密码，用于设置超级管理员的新密码
//
// 返回值：
//
//	error：返回错误信息，如果设置成功则返回 nil
func (c *Config) SetSuperAdminPassword(oldpass string, newpass string) error {
	username, password, err := libraries.LoadUserConfig()
	if err != nil {
		return err
	}

	fmt.Println("oldpass:", oldpass)
	fmt.Println("oldpass:", helper.HashPassword(oldpass))
	fmt.Println("password:", password)
	// 验证旧密码是否正确
	if helper.HashPassword(oldpass) != password {
		return cron.CreateCustomError(500, "旧密码错误")
	}
	// 检查旧密码和新密码是否相同
	if oldpass == newpass {
		return cron.CreateCustomError(500, "新密码不能与旧密码相同")
	}

	// 检查新密码长度是否小于6位
	if len(newpass) < 6 {
		return cron.CreateCustomError(500, "新密码长度不能小于6位")
	}

	// 检查新密码是否包含数字、字母和符号
	hasNumber := false
	hasLetter := false
	hasSymbol := false
	for _, char := range newpass {
		switch {
		case char >= '0' && char <= '9':
			hasNumber = true
		case (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z'):
			hasLetter = true
		case (char >= '!' && char <= '/') || (char >= ':' && char <= '@') || (char >= '[' && char <= '`') || (char >= '{' && char <= '~'):
			hasSymbol = true
		}
	}

	if !hasNumber || !hasLetter || !hasSymbol {
		return cron.CreateCustomError(500, "新密码必须包含数字、字母和符号")
	}

	// 保存配置
	err = libraries.SaveUserConfig(username, newpass)
	if err != nil {
		return err
	}

	return nil
}
