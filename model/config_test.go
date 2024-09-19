package model

import (
	"testing"
)

func TestGetBlackListType(t *testing.T) {
	// Mock the LoadWhitelistConfig function
	config := Config{}
	blackListType := config.GetBlackListType()

	if blackListType != 1 {
		t.Errorf("期望黑名单类型为 %d，但实际为 %d", 1, blackListType)
	}
}

func TestSetBlackListType(t *testing.T) {
	// Mock the SaveWhitelistConfig function
	config := Config{}
	config.SetBlackListType(2)

	blackListType := config.GetBlackListType()

	if blackListType != 2 {
		t.Errorf("期望黑名单类型为 %d，但实际为 %d", 2, blackListType)
	}
}

func TestSetSuperAdminPassword(t *testing.T) {
	// Mock the LoadUserConfig and SaveUserConfig functions
	config := Config{}
	err := config.SetSuperAdminPassword("old_password", "new_password")

	if err != nil {
		t.Errorf("设置超级管理员密码时发生错误: %v", err)
	}
}
