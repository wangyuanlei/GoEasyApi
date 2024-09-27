package model

import (
	"GoEasyApi/database"
	"GoEasyApi/libraries"
	"regexp"

	"github.com/google/uuid"
)

type Interface struct{}

// 增加接口
func (m *Interface) AddInterface(info database.Interface) error {
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

	var existingInterface database.Interface
	if err := DB.Where("path = ? AND method = ?", info.Path, info.Method).First(&existingInterface).Error; err == nil {
		return libraries.CreateCustomError(601, "接口路径和方法已存在")
	}

	info.InterfaceId = uuid.New().String()

	err := DB.Create(&info).Error

	if len(info.Params) > 0 { //如果参数不为空. 保存数据
		m.SaveParams(info.InterfaceId, info.Params)
	}

	return err
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

//获得接口详情

//获得接口列表

// 检查method 类型
func (m *Interface) CheckMethod(method string) error {
	if method != "post" && method != "get" {
		return libraries.CreateCustomError(601, "Method 只能是 get 或者 post")
	}

	return nil
}

func (m *Interface) CheckEnabled(data int) error {

	if data != 1 && data != 2 {
		return libraries.CreateCustomError(601, "值设置错误")
	}

	return nil
}

func (m *Interface) CheckStringFormat(str string) error {
	if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(str) {
		return libraries.CreateCustomError(601, "字符串只能包含大小写字母、数字和下划线")
	}
	return nil
}

func (m *Interface) CheckParamType(paramType string) error {
	if !regexp.MustCompile(`^(string|int|float|bool|date|datetime)$`).MatchString(paramType) {
		return libraries.CreateCustomError(601, "参数类型错误")
	}
	return nil
}
