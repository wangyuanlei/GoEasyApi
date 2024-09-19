package libraries

import (
	"GoEasyApi/helper"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestLoadConfig(t *testing.T) {
	// 创建一个测试配置文件
	config := Config{
		Database:        "test_database1",
		WhitelistConfig: 2,
		User: struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}{
			Username: "test_user1",
			Password: "test_password1",
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("config.yml", data, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("config.yml")

	// 加载配置
	loadedConfig, err := LoadConfig()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("加载的配置: %+v", loadedConfig)

	// 检查加载的配置
	if loadedConfig.Database != config.Database {
		t.Errorf("期望的数据库 %s, 得到的 %s", config.Database, loadedConfig.Database)
	}
	if loadedConfig.WhitelistConfig != config.WhitelistConfig {
		t.Errorf("期望的白名单配置 %d, 得到的 %d", config.WhitelistConfig, loadedConfig.WhitelistConfig)
	}
	if loadedConfig.User.Username != config.User.Username {
		t.Errorf("期望的用户名 %s, 得到的 %s", config.User.Username, loadedConfig.User.Username)
	}
	if loadedConfig.User.Password != config.User.Password {
		t.Errorf("期望的密码 %s, 得到的 %s", config.User.Password, loadedConfig.User.Password)
	}
}

func TestSaveConfig(t *testing.T) {
	// 创建一个测试配置
	config := Config{
		Database:        "test_database22",
		WhitelistConfig: 1,
		User: struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}{
			Username: "test_user11",
			Password: "test_password22",
		},
	}

	// 保存配置
	err := SaveConfig(config)
	if err != nil {
		t.Fatal(err)
	}
	// defer os.Remove("config.yml")

	// 加载保存的配置
	loadedConfig, err := LoadConfig()
	if err != nil {
		t.Fatal(err)
	}

	// 检查加载的配置
	if loadedConfig.Database != config.Database {
		t.Errorf("期望的数据库 %s, 得到的 %s", config.Database, loadedConfig.Database)
	}
	if loadedConfig.WhitelistConfig != config.WhitelistConfig {
		t.Errorf("期望的白名单配置 %d, 得到的 %d", config.WhitelistConfig, loadedConfig.WhitelistConfig)
	}
	if loadedConfig.User.Username != config.User.Username {
		t.Errorf("期望的用户名 %s, 得到的 %s", config.User.Username, loadedConfig.User.Username)
	}
	if loadedConfig.User.Password != config.User.Password {
		t.Errorf("期望的密码 %s, 得到的 %s", config.User.Password, loadedConfig.User.Password)
	}
}

func TestLoadDatabaseConfig(t *testing.T) {
	// 加载数据库配置
	expectedDatabase := "test_database22"
	database, err := LoadDatabaseConfig()
	if err != nil {
		t.Fatal(err)
	}

	// 检查加载的数据库配置
	if database != expectedDatabase {
		t.Errorf("期望的数据库 %s, 得到的 %s", expectedDatabase, database)
	}
}

func TestGetBind(t *testing.T) {
	// 创建一个测试配置文件
	config := Config{
		Database:        "db.sql",
		WhitelistConfig: 1,
		Bind:            "127.0.0.1:8008",
		User: struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}{
			Username: "admin",
			Password: helper.HashPassword("admin"),
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("config.yml", data, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// 加载绑定配置
	bind, err := GetBind()
	if err != nil {
		t.Fatal(err)
	}

	// 检查加载的绑定配置
	if bind != config.Bind {
		t.Errorf("期望的绑定配置 %s, 得到的 %s", config.Bind, bind)
	}
}

func TestSetBind(t *testing.T) {
	// 创建一个测试配置文件
	config := Config{
		Database:        "db.sql",
		WhitelistConfig: 1,
		Bind:            "127.0.0.1:8008",
		User: struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}{
			Username: "admin",
			Password: helper.HashPassword("admin"),
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("config.yml", data, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// 设置新的绑定配置
	newBind := "192.168.1.1:8080"
	err = SetBind(newBind)
	if err != nil {
		t.Fatal(err)
	}

	// 加载新的绑定配置
	updatedBind, err := GetBind()
	if err != nil {
		t.Fatal(err)
	}

	// 检查加载的新的绑定配置
	if updatedBind != newBind {
		t.Errorf("期望的绑定配置 %s, 得到的 %s", newBind, updatedBind)
	}
}

func TestLoadUserConfig(t *testing.T) {
	// 创建一个测试配置文件
	config := Config{
		Database:        "test_database1",
		WhitelistConfig: 2,
		User: struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}{
			Username: "test_user1",
			Password: "test_password1",
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("config.yml", data, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// 加载用户配置
	username, password, err := LoadUserConfig()
	if err != nil {
		t.Fatal(err)
	}

	// 检查加载的用户配置
	if username != config.User.Username {
		t.Errorf("期望的用户名 %s, 得到的 %s", config.User.Username, username)
	}
	if password != config.User.Password {
		t.Errorf("期望的密码 %s, 得到的 %s", config.User.Password, password)
	}
}

func TestSaveUserConfig(t *testing.T) {
	// 创建一个测试配置文件
	config := Config{
		Database:        "test_database1",
		WhitelistConfig: 2,
		User: struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}{
			Username: "test_user",
			Password: "test_password",
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("config.yml", data, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// 加载保存的用户配置
	loadedUsername, loadedPassword, err := LoadUserConfig()
	if err != nil {
		t.Fatal(err)
	}

	// 检查加载的用户配置
	if loadedUsername != config.User.Username {
		t.Errorf("期望的用户名 %s, 得到的 %s", config.User.Username, loadedUsername)
	}
	if loadedPassword != config.User.Password {
		t.Errorf("期望的密码 %s, 得到的 %s", config.User.Password, loadedPassword)
	}
}

func TestLoadWhitelistConfig(t *testing.T) {
	// 创建一个测试配置文件
	config := Config{
		Database:        "test_database1",
		WhitelistConfig: 2,
		User: struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}{
			Username: "test_user",
			Password: "test_password",
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("config.yml", data, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// 加载白名单配置
	whitelistConfig, err := LoadWhitelistConfig()
	if err != nil {
		t.Fatal(err)
	}

	// 检查加载的白名单配置
	if whitelistConfig != config.WhitelistConfig {
		t.Errorf("期望的白名单配置 %d, 得到的 %d", config.WhitelistConfig, whitelistConfig)
	}
}

func TestSaveWhitelistConfig(t *testing.T) {
	// 创建一个测试配置文件
	config := Config{
		Database:        "db.sql",
		WhitelistConfig: 1,
		User: struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}{
			Username: "admin",
			Password: helper.HashPassword("admin"),
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("config.yml", data, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// 加载保存的白名单配置
	loadedWhitelistConfig, err := LoadWhitelistConfig()
	if err != nil {
		t.Fatal(err)
	}

	// 检查加载的白名单配置
	if loadedWhitelistConfig != config.WhitelistConfig {
		t.Errorf("期望的白名单配置 %d, 得到的 %d", config.WhitelistConfig, loadedWhitelistConfig)
	}
}
