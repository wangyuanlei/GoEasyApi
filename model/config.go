package model

import (
	"GoEasyApi/helper"
	"GoEasyApi/libraries"
)

type Config struct{}

// 获得名单类型
func (c *Config) GetBlackListType() int {
	whitelistConfig, err := libraries.LoadWhitelistConfig()
	if err != nil {
		return 0 // 或者根据需要返回一个默认值或处理错误
	}
	return whitelistConfig
}

// 设置使用名单类型 0-不使用 1-使用黑名单 2-使用白名单
func (c *Config) SetBlackListType(blackListType int) error {
	err := libraries.SaveWhitelistConfig(blackListType)
	if err != nil {
		return err
	}
	return nil
}

// 设置超级管理员密码
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
