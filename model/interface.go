package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/database"
	"regexp"

	"github.com/google/uuid"
)

type Interface struct{}

// 增加接口
func (m *Interface) AddInterface(info database.Interface) (string, error) {
	if err := m.CheckMethod(info.Method); err != nil {
		return "", err
	}

	if err := m.CheckEnabled(info.CacheEnabled); err != nil {
		return "", cron.CreateCustomError(601, "CacheEnabled 值错误, 1是 2否")
	}

	if err := m.CheckEnabled(info.RateLimitEnabled); err != nil {
		return "", cron.CreateCustomError(601, "RateLimitEnabled 值错误, 1是 2否")
	}

	if err := m.CheckEnabled(info.TokenValidationEnabled); err != nil {
		return "", cron.CreateCustomError(601, "是否设置验证token的值设置错误, 1是 2否")
	}

	var existingInterface database.Interface

	if DB.Where("path = ? AND method = ?", info.Path, info.Method).First(&existingInterface); existingInterface.InterfaceId != "" {
		return "", cron.CreateCustomError(601, "接口["+info.Method+"]"+info.Path+" 已存在")
	}

	info.InterfaceId = uuid.New().String()

	err := DB.Create(&info).Error
	return info.InterfaceId, err
}

// 修改修改
func (m *Interface) UpdateInterface(info database.Interface) error {
	if info.InterfaceId == "" {
		return cron.CreateCustomError(601, "InterfaceId 不能为空")
	}

	var existingInterface database.Interface
	if err := DB.Where("interface_id = ?", info.InterfaceId).First(&existingInterface).Error; err != nil {
		return cron.CreateCustomError(601, "数据不存在")
	}

	if err := m.CheckMethod(info.Method); err != nil {
		return err
	}

	if err := m.CheckEnabled(info.CacheEnabled); err != nil {
		return cron.CreateCustomError(601, "CacheEnabled 值错误, 1是 2否")
	}

	if err := m.CheckEnabled(info.RateLimitEnabled); err != nil {
		return cron.CreateCustomError(601, "RateLimitEnabled 值错误, 1是 2否")
	}

	if err := m.CheckEnabled(info.TokenValidationEnabled); err != nil {
		return cron.CreateCustomError(601, "是否设置验证token的值设置错误, 1是 2否")
	}

	var existingInterface2 database.Interface
	if err := DB.Where("interface_id != ? AND path = ? AND method = ?", info.InterfaceId, info.Path, info.Method).First(&existingInterface2).Error; err == nil {
		return cron.CreateCustomError(601, "接口["+info.Method+"]"+info.Path+" 已存在")
	}

	err := DB.Model(&database.Interface{}).Updates(info).Error
	return err
}

// 删除接口
func (m *Interface) DeleteInterface(interfaceId string) error {
	var existingInterface database.Interface
	if err := DB.Where("interface_id = ?", interfaceId).First(&existingInterface).Error; err != nil {
		return cron.CreateCustomError(601, "数据不存在")
	}

	if err := DB.Delete(&existingInterface).Error; err != nil {
		return err
	}

	if err := DB.Where("interface_id = ?", interfaceId).Delete(&database.Params{}).Error; err != nil {
		return err
	}

	return nil
}

// 添加接口参数
func (m *Interface) AddParams(interfaceId string, params database.Params) error {
	if err := m.CheckStringFormat(params.Name); err != nil {
		return err
	}

	if err := m.CheckParamType(params.Type); err != nil {
		return err
	}

	if err := m.CheckEnabled(params.Required); err != nil {
		return cron.CreateCustomError(601, "是否必传参数错误")
	}

	params.InterfaceId = interfaceId
	params.ParamsId = uuid.New().String()
	if err := DB.Create(&params).Error; err != nil {
		return err
	}

	return nil
}

// 修改参数
func (m *Interface) UpdateParams(interfaceId string, paramsId string, params database.Params) error {
	var existingParam database.Params
	if err := DB.Where("interface_id = ? AND params_id = ?", interfaceId, paramsId).First(&existingParam).Error; err != nil {
		return cron.CreateCustomError(601, "参数不存在")
	}

	if err := m.CheckStringFormat(params.Name); err != nil {
		return err
	}

	if err := m.CheckParamType(params.Type); err != nil {
		return err
	}

	if err := m.CheckEnabled(params.Required); err != nil {
		return cron.CreateCustomError(601, "是否必传参数错误")
	}

	params.InterfaceId = interfaceId
	params.ParamsId = existingParam.ParamsId
	if err := DB.Model(&existingParam).Updates(params).Error; err != nil {
		return err
	}

	return nil
}

// 删除参数
func (m *Interface) DeleteParams(interfaceId string, paramsId string) error {
	var existingParam database.Params
	if err := DB.Where("interface_id = ? AND params_id = ?", interfaceId, paramsId).First(&existingParam).Error; err != nil {
		return cron.CreateCustomError(601, "参数不存在")
	}

	if err := DB.Delete(&existingParam).Error; err != nil {
		return err
	}

	return nil
}

type InterfaceList struct {
	InterfaceId   string `gorm:"type:varchar(32);primary_key"` // 接口编号
	InterfaceName string `gorm:"type:varchar(50)"`             // 接口名称
	Description   string `gorm:"type:text"`                    // 接口描述
	Path          string `gorm:"type:varchar(255)"`            // 接口路径
	Method        string `gorm:"type:varchar(10)"`             // 接口方法
}

// 获得接口列表
func (m *Interface) GetList() ([]InterfaceList, error) {
	/*
		实现逻辑以下逻辑
		1. 获得接口列表.
		2. 数据结构是 database.Interface
		3. 输出的字段根据 InterfaceList 结构来
		3. 根据 Path 字段顺序排序
	*/
	var interfaceList []InterfaceList
	if err := DB.Model(&database.Interface{}).Select("interface_id", "interface_name", "description", "path", "method").Order("path asc").Find(&interfaceList).Error; err != nil {
		return nil, err
	}
	return interfaceList, nil
}

// 获得接口详情
func (m *Interface) GetInfo(InterfaceId string) (database.Interface, error) {
	/*
		现逻辑以下逻辑
		1. 获得接口数据详情
		2. 数据结构是 database.Interface
		3. 同时输出 Params 详细内容
	*/
	var interfaceInfo database.Interface
	if err := DB.Model(&database.Interface{}).Where("interface_id = ?", InterfaceId).First(&interfaceInfo).Error; err != nil {
		return database.Interface{}, err
	}

	var params []database.Params
	if err := DB.Model(&database.Params{}).Where("interface_id = ?", InterfaceId).Find(&params).Error; err != nil {
		return database.Interface{}, err
	}

	interfaceInfo.Params = params
	return interfaceInfo, nil
}

// 检查method 类型 只能是 post 和 get
func (m *Interface) CheckMethod(method string) error {
	if method != "post" && method != "get" {
		return cron.CreateCustomError(601, "Method 只能是 get 或者 post")
	}

	return nil
}

// 检查是否有效的值
func (m *Interface) CheckEnabled(data int) error {

	if data != 1 && data != 2 {
		return cron.CreateCustomError(601, "值设置错误")
	}

	return nil
}

// 检查参数的名称
func (m *Interface) CheckStringFormat(str string) error {
	if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(str) {
		return cron.CreateCustomError(601, "字符串只能包含大小写字母、数字和下划线")
	}
	return nil
}

// 检查参数的类型定义是否准确
func (m *Interface) CheckParamType(paramType string) error {
	if !regexp.MustCompile(`^(string|int|float|bool|date|datetime)$`).MatchString(paramType) {
		return cron.CreateCustomError(601, "参数类型错误")
	}
	return nil
}
