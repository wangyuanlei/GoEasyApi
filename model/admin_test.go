package model

import (
	"GoEasyApi/libraries"
	"fmt"
	"testing"
)

func TestAdmin_Login(t *testing.T) {
	// 准备测试数据
	admin := &Admin{}
	username, password := "testUser", "testPassword"

	// 设置配置
	libraries.SaveUserConfig(username, password)

	// 测试登录
	token, err := admin.Login(username, password)
	if err != nil {
		t.Fatalf("登录失败: %v", err)
	}

	// 验证 token 是否有效
	if token == "" {
		t.Fatal("生成的 token 为空")
	}
}

func TestAdmin_ValidateLogin(t *testing.T) {
	admin := &Admin{}
	username, password := "testUser", "testPassword"

	// 设置配置
	libraries.SaveUserConfig(username, password)

	// 测试登录以获取 token
	token, err := admin.Login(username, password)
	if err != nil {
		t.Fatalf("登录失败: %v", err)
	}

	fmt.Print("token:", token)

	// 测试验证登录
	validatedUsername, err := admin.ValidateLogin(token)
	if err != nil {
		t.Fatalf("验证登录失败: %v", err)
	}

	// 验证返回的用户名是否正确
	if validatedUsername != username {
		t.Fatalf("验证的用户名不匹配: got %v, want %v", validatedUsername, username)
	}
}
