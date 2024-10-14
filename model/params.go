package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/database"
	"GoEasyApi/helper"

	"github.com/google/uuid"
)

// 添加接口参数(单个)
func (m *Interface) AddParams(interfaceId string, params database.Params) error {
	if err := helper.CheckStringFormat(params.Name); err != nil {
		return err
	}

	if err := m.CheckParamType(params.Type); err != nil {
		return err
	}

	if err := helper.CheckEnabled(params.Required); err != nil {
		return cron.CreateCustomError(601, "是否必传参数错误")
	}

	params.InterfaceId = interfaceId
	params.ParamsId = uuid.New().String()
	if err := DB.Create(&params).Error; err != nil {
		return err
	}

	return nil
}

// 修改参数 (单个)
func (m *Interface) UpdateParams(interfaceId string, paramsId string, params database.Params) error {
	var existingParam database.Params
	if err := DB.Where("interface_id = ? AND params_id = ?", interfaceId, paramsId).First(&existingParam).Error; err != nil {
		return cron.CreateCustomError(601, "参数不存在")
	}

	if err := helper.CheckStringFormat(params.Name); err != nil {
		return err
	}

	if err := helper.CheckParamItem(params.Type, "string|int|float|bool|date|datetime"); err != nil {
		return err
	}

	if err := helper.CheckEnabled(params.Required); err != nil {
		return cron.CreateCustomError(601, "是否必传参数错误")
	}

	params.InterfaceId = interfaceId
	params.ParamsId = existingParam.ParamsId
	if err := DB.Model(&existingParam).Updates(params).Error; err != nil {
		return err
	}

	return nil
}

// 删除参数 (单个)
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

// 保存接口的参数数据 (多个批量更新)
func (m *Interface) SaveParams(interfaceId string, params []database.Params) error {
	existingParamsIds := []string{}
	for _, param := range params {

		if err := helper.CheckStringFormat(param.Name); err != nil {
			return err
		}

		if err := m.CheckParamType(param.Type); err != nil {
			return err
		}

		if err := helper.CheckEnabled(param.Required); err != nil {
			return cron.CreateCustomError(601, "是否必传参数错误")
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

// 检查参数的类型定义是否准确
func (m *Interface) CheckParamType(paramType string) error {
	if err := helper.CheckParamItem(paramType, "string|int|float|bool|date|datetime"); err != nil {
		return err
	}

	return nil
}
