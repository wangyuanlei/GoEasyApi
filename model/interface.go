package model

import (
	"GoEasyApi/database"
	"GoEasyApi/libraries"
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
		return "", libraries.CreateCustomError(601, "CacheEnabled 值错误, 1是 2否")
	}

	if err := m.CheckEnabled(info.RateLimitEnabled); err != nil {
		return "", libraries.CreateCustomError(601, "RateLimitEnabled 值错误, 1是 2否")
	}

	if err := m.CheckEnabled(info.TokenValidationEnabled); err != nil {
		return "", libraries.CreateCustomError(601, "是否设置验证token的值设置错误, 1是 2否")
	}

	var existingInterface database.Interface
	if err := DB.Where("path = ? AND method = ?", info.Path, info.Method).First(&existingInterface).Error; err == nil {
		return "", libraries.CreateCustomError(601, "接口路径和方法已存在")
	}

	info.InterfaceId = uuid.New().String()

	err := DB.Create(&info).Error

	if len(info.Params) > 0 { //如果参数不为空. 保存数据
		m.SaveParams(info.InterfaceId, info.Params)
	}

	return info.InterfaceId, err
}

// 修改修改
func (m *Interface) UpdateInterface(info database.Interface) error {
	if info.InterfaceId == "" {
		return libraries.CreateCustomError(601, "InterfaceId 不能为空")
	}

	var existingInterface database.Interface
	if err := DB.Where("interface_id = ?", info.InterfaceId).First(&existingInterface).Error; err != nil {
		return libraries.CreateCustomError(601, "数据不存在")
	}

	if err := m.CheckMethod(info.Method); err != nil {
		return err
	}

	if err := m.CheckEnabled(info.CacheEnabled); err != nil {
		return libraries.CreateCustomError(601, "CacheEnabled 值错误, 1是 2否")
	}

	if err := m.CheckEnabled(info.RateLimitEnabled); err != nil {
		return libraries.CreateCustomError(601, "RateLimitEnabled 值错误, 1是 2否")
	}

	if err := m.CheckEnabled(info.TokenValidationEnabled); err != nil {
		return libraries.CreateCustomError(601, "是否设置验证token的值设置错误, 1是 2否")
	}

	var existingInterface2 database.Interface
	if err := DB.Where("interface_id != ? AND path = ? AND method = ?", info.InterfaceId, info.Path, info.Method).First(&existingInterface2).Error; err == nil {
		return libraries.CreateCustomError(601, "接口路径和方法已存在")
	}

	err := DB.Model(&database.Interface{}).Updates(info).Error
	if len(info.Params) > 0 {
		m.SaveParams(info.InterfaceId, info.Params)
	} else {
		//如果空 则情况一次参数表
		if err := DB.Where("interface_id = ?", info.InterfaceId).Delete(&database.Params{}).Error; err != nil {
			return err
		}
	}

	return err
}

// 删除接口
func (m *Interface) DeleteInterface(interfaceId string) error {
	var existingInterface database.Interface
	if err := DB.Where("interface_id = ?", interfaceId).First(&existingInterface).Error; err != nil {
		return libraries.CreateCustomError(601, "数据不存在")
	}

	if err := DB.Delete(&existingInterface).Error; err != nil {
		return err
	}

	if err := DB.Where("interface_id = ?", interfaceId).Delete(&database.Params{}).Error; err != nil {
		return err
	}

	return nil
}

// 保存接口的参数数据
func (m *Interface) SaveParams(interfaceId string, params []database.Params) error {
	existingParamsIds := []string{}
	for _, param := range params {

		if err := m.CheckStringFormat(param.Name); err != nil {
			return err
		}

		if err := m.CheckParamType(param.Type); err != nil {
			return err
		}

		if err := m.CheckEnabled(param.Required); err != nil {
			return libraries.CreateCustomError(601, "是否必传参数错误")
		}

		param.InterfaceId = interfaceId
		var existingParam database.Params
		if err := DB.Where("params_id = ? AND interface_id = ?", param.ParamsId, interfaceId).First(&existingParam).Error; err == nil {
			// Update existing param
			if err := DB.Save(&param).Error; err != nil {
				return err
			}
		} else {
			// Create new param
			param.ParamsId = uuid.New().String()
			if err := DB.Create(&param).Error; err != nil {
				return err
			}

		}
		existingParamsIds = append(existingParamsIds, param.ParamsId)
	}

	//数据库内 该接口下 不存在的参数删除掉
	if err := DB.Where("interface_id = ? AND params_id NOT IN (?)", interfaceId, existingParamsIds).Delete(&database.Params{}).Error; err != nil {
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
		return libraries.CreateCustomError(601, "Method 只能是 get 或者 post")
	}

	return nil
}

// 检查是否有效的值
func (m *Interface) CheckEnabled(data int) error {

	if data != 1 && data != 2 {
		return libraries.CreateCustomError(601, "值设置错误")
	}

	return nil
}

// 检查参数的名称
func (m *Interface) CheckStringFormat(str string) error {
	if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(str) {
		return libraries.CreateCustomError(601, "字符串只能包含大小写字母、数字和下划线")
	}
	return nil
}

// 检查参数的类型定义是否准确
func (m *Interface) CheckParamType(paramType string) error {
	if !regexp.MustCompile(`^(string|int|float|bool|date|datetime)$`).MatchString(paramType) {
		return libraries.CreateCustomError(601, "参数类型错误")
	}
	return nil
}
