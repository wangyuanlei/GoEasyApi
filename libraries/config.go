package libraries

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config 结构体用于保存配置数据
type Config struct {
	Database string `yaml:"database"`
	User     struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"user"`
}

// LoadConfig 函数用于从配置文件加载配置
func LoadConfig() (Config, error) {
	var config Config
	data, err := os.ReadFile("config.yml") // 更新为使用 os.ReadFile
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

// SaveConfig 函数用于将配置保存到配置文件
func SaveConfig(config Config) error {
	data, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}
	err = os.WriteFile("config.yml", data, 0644) // 更新为使用 os.WriteFile
	if err != nil {
		return err
	}
	return nil
}

// LoadDatabaseConfig 函数用于从配置文件加载数据库配置
func LoadDatabaseConfig() (string, error) {
	config, err := LoadConfig()
	if err != nil {
		return "", err
	}
	return config.Database, nil
}

// LoadUserConfig 函数用于从配置文件加载用户配置
func LoadUserConfig() (string, string, error) {
	config, err := LoadConfig()
	if err != nil {
		return "", "", err
	}
	return config.User.Username, config.User.Password, nil
}

// SaveUserConfig 函数用于将用户配置保存到配置文件
func SaveUserConfig(username, password string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}
	config.User.Username = username
	config.User.Password = password
	err = SaveConfig(config)
	if err != nil {
		return err
	}
	return nil
}
