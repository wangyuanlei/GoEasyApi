package model

import (
	"GoEasyApi/helper"
	"GoEasyApi/libraries"
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

	// 验证旧密码是否正确
	if helper.HashPassword(oldpass) != password {
		return libraries.CreateCustomError(500, "旧密码错误")
	}

	// 设置新密码
	newPassword := helper.HashPassword(newpass)

	// 保存配置
	err = libraries.SaveUserConfig(username, newPassword)
	if err != nil {
		return err
	}

	return nil
}
