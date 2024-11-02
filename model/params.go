package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/helper"
	"GoEasyApi/structs"
)

// 参数检验
func (m *Interface) ParamsVerify(params structs.Params) error {
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
func (m *Interface) CheckParamRequired(params structs.Params) error {
	if params.Name == "" {
		return cron.CreateCustomError(601, "name 必传参数未提交")
	}
	if params.Type == "" {
		return cron.CreateCustomError(601, "type 必传参数未提交")
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
