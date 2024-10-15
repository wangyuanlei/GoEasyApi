package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/database"
	"GoEasyApi/helper"

	"github.com/google/uuid"
)

// 参数检验
func (m *Interface) ParamsVerify(params database.Params) error {
	if err := m.CheckParamRequired(params); err != nil {
		return err
	}

	if err := helper.CheckStringFormat(params.Name); err != nil {
		return err
	}

	if err := m.CheckParamType(params.Type); err != nil {
		return err
	}

	if err := helper.CheckEnabled(params.Required); err != nil {
		return cron.CreateCustomError(601, "传入的值无效")
	}
	return nil
}

// 检查参数是否必传
func (m *Interface) CheckParamRequired(params database.Params) error {
	if params.Name == "" {
		return cron.CreateCustomError(601, "name 必传参数未提交")
	}
	if params.Type == "" {
		return cron.CreateCustomError(601, "type 必传参数未提交")
	}

	return nil
}

// 检查参数数据是否存在, 存在则返回存在的数据
func (m *Interface) ParamsExist(interfaceId string, paramsId string) (database.Params, error) {
	var existingParam database.Params
	if err := DB.Where("interface_id = ? AND params_id = ?", interfaceId, paramsId).First(&existingParam).Error; err != nil {
		return existingParam, cron.CreateCustomError(601, "接口参数不存在")
	}

	return existingParam, nil
}

// 添加接口参数(单个)
func (m *Interface) AddParams(interfaceId string, params database.Params) error {
	if err := m.ParamsVerify(params); err != nil {
		return err
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
	if err := m.ParamsVerify(params); err != nil {
		return err
	}

	existingParam, err := m.ParamsExist(interfaceId, paramsId)
	if err != nil {
		return err
	}

	params.InterfaceId = existingParam.InterfaceId
	params.ParamsId = existingParam.ParamsId
	if err := DB.Model(&existingParam).Updates(params).Error; err != nil {
		return err
	}

	return nil
}

// 删除参数 (单个)
func (m *Interface) DeleteParams(interfaceId string, paramsId string) error {
	existingParam, err := m.ParamsExist(interfaceId, paramsId)
	if err != nil {
		return err
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

		if err := m.ParamsVerify(param); err != nil {
			return err
		}

		param.InterfaceId = interfaceId

		if param.ParamsId == "" {
			// Create new param
			param.ParamsId = uuid.New().String()
			if err := DB.Create(&param).Error; err != nil {
				return err
			}
		} else {
			_, err := m.ParamsExist(param.InterfaceId, param.ParamsId)
			if err != nil {
				return err
			}
			// Update existing param
			if err := DB.Save(&param).Error; err != nil {
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
