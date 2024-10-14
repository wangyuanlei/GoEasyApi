package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/database"
	"GoEasyApi/helper"

	"github.com/google/uuid"
)

type Interface struct{}

func (m *Interface) InterfaceVerify(info database.Interface) error {
	if err := helper.CheckMethod(info.Method); err != nil {
		return err
	}

	if err := helper.CheckEnabled(info.CacheEnabled); err != nil {
		return cron.CreateCustomError(601, "CacheEnabled 值错误, 1是 2否")
	}

	if err := helper.CheckEnabled(info.RateLimitEnabled); err != nil {
		return cron.CreateCustomError(601, "RateLimitEnabled 值错误, 1是 2否")
	}

	if err := helper.CheckEnabled(info.TokenValidationEnabled); err != nil {
		return cron.CreateCustomError(601, "是否设置验证token的值设置错误, 1是 2否")
	}

	if err := helper.CheckParamItem(info.ReturnType, "string|json|list|page"); err != nil {
		return err
	}

	return nil
}

// 增加接口
func (m *Interface) AddInterface(info database.Interface) (string, error) {

	if err := m.InterfaceVerify(info); err != nil {
		return "", err
	}

	var existingInterface database.Interface

	if DB.Where("path = ? AND method = ?", info.Path, info.Method).First(&existingInterface); existingInterface.InterfaceId != "" {
		return "", cron.CreateCustomError(601, "接口["+info.Method+"]"+info.Path+" 已存在")
	}

	info.InterfaceId = uuid.New().String()

	if err := DB.Create(&info).Error; err != nil {
		return "", err
	}

	if len(info.Params) > 0 { //如果参数不为空. 保存数据
		m.SaveParams(info.InterfaceId, info.Params)
	}

	return info.InterfaceId, nil
}

// 修改修改
func (m *Interface) UpdateInterface(info database.Interface) error {
	if info.InterfaceId == "" {
		return cron.CreateCustomError(601, "InterfaceId 不能为空")
	}

	if err := m.InterfaceVerify(info); err != nil {
		return err
	}

	var existingInterface database.Interface
	if err := DB.Where("interface_id = ?", info.InterfaceId).First(&existingInterface).Error; err != nil {
		return cron.CreateCustomError(601, "数据不存在")
	}

	var existingInterface2 database.Interface
	if err := DB.Where("interface_id != ? AND path = ? AND method = ?", info.InterfaceId, info.Path, info.Method).First(&existingInterface2).Error; err == nil {
		return cron.CreateCustomError(601, "接口["+info.Method+"]"+info.Path+" 已存在")
	}

	if err := DB.Model(&database.Interface{}).Updates(info).Error; err != nil {
		return err
	}

	if len(info.Params) > 0 {
		m.SaveParams(info.InterfaceId, info.Params)
	} else {
		//如果空 则情况一次参数表
		if err := DB.Where("interface_id = ?", info.InterfaceId).Delete(&database.Params{}).Error; err != nil {
			return err
		}
	}

	return nil
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
