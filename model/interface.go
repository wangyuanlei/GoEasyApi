package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/helper"
	"GoEasyApi/libraries"
	"GoEasyApi/structs"
	"strings"

	"github.com/google/uuid"
)

type Interface struct{}

// 接口检验
func (m *Interface) InterfaceVerify(info structs.Interface) error {
	if err := m.CheckInterfaceRequired(info); err != nil {
		return err
	}

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

	if err := helper.CheckParamItem(info.ReturnType, "row|list|update|delete|insert"); err != nil {
		return err
	}

	if info.ReturnValMode != "" { //如果不为空, 值只能是 last_id|row
		if err := helper.CheckParamItem(info.ReturnValMode, "last_id|row"); err != nil {
			return err
		}
	}

	return nil
}

// 检查参数是否必传
func (m *Interface) CheckInterfaceRequired(params structs.Interface) error {
	if params.Name == "" {
		return cron.CreateCustomError(601, "interfacename 必传参数未提交")
	}
	if params.Path == "" {
		return cron.CreateCustomError(601, "path 必传参数未提交")
	}
	if params.Method == "" {
		return cron.CreateCustomError(601, "method 必传参数未提交")
	}
	return nil
}

// 增加接口
func (m *Interface) AddInterface(info structs.Interface) (string, error) {

	if err := m.InterfaceVerify(info); err != nil {
		return "", err
	}

	var existingInterface structs.Interface
	if DB.Where("path = ? AND method = ?", info.Path, info.Method).First(&existingInterface); existingInterface.Id != "" {
		return "", cron.CreateCustomError(601, "接口["+info.Method+"]"+info.Path+" 已存在")
	}

	info.Id = uuid.New().String()

	if err := DB.Create(&info).Error; err != nil {
		return "", err
	}

	if len(info.Params) > 0 { //如果参数不为空. 保存数据
		for _, param := range info.Params {
			if err := m.ParamsVerify(param); err != nil {
				return "", err
			}
		}
	}

	//更新缓存
	if err := m.UpdateCacheByPath(info.Id); err != nil {
		return "", err
	}

	return info.Id, nil
}

// 修改修改
func (m *Interface) UpdateInterface(info structs.Interface) error {
	if info.Id == "" {
		return cron.CreateCustomError(601, "InterfaceId 不能为空")
	}

	if err := m.InterfaceVerify(info); err != nil {
		return err
	}

	var existingInterface structs.Interface
	if err := DB.Where("id = ?", info.Id).First(&existingInterface).Error; err != nil {
		return cron.CreateCustomError(601, "数据不存在")
	}

	var existingInterface2 structs.Interface
	if err := DB.Where("id != ? AND path = ? AND method = ?", info.Id, info.Path, info.Method).First(&existingInterface2).Error; err == nil {
		return cron.CreateCustomError(601, "接口["+info.Method+"]"+info.Path+" 已存在")
	}

	if len(info.Params) > 0 {
		for _, param := range info.Params {
			if err := m.ParamsVerify(param); err != nil {
				return err
			}
		}
	}

	if err := DB.Model(&structs.Interface{}).Where("id = ?", info.Id).Updates(info).Error; err != nil {
		return err
	}
	//保存Params信息
	//清空一次参数表
	if err := DB.Where("interface_id = ?", info.Id).Delete(&structs.Params{}).Error; err != nil {
		return err
	}
	if len(info.Params) > 0 {
		for _, param := range info.Params {
			param.InterfaceId = info.Id
			if err := DB.Create(&param).Error; err != nil {
				return err
			}
		}
	}

	//更新缓存
	if err := m.UpdateCacheByPath(info.Id); err != nil {
		return err
	}

	return nil
}

// 删除接口
func (m *Interface) DeleteInterface(interfaceId string) error {
	var existingInterface structs.Interface
	if err := DB.Where("id = ?", interfaceId).First(&existingInterface).Error; err != nil {
		return cron.CreateCustomError(601, "数据不存在")
	}

	if err := DB.Delete(&existingInterface).Error; err != nil {
		return err
	}

	if err := DB.Where("interface_id = ?", interfaceId).Delete(&structs.Params{}).Error; err != nil {
		return err
	}

	//清除缓存
	cacheKey := m.GetCacheKeyByPath(existingInterface.Path, existingInterface.Method)
	libraries.DeleteCache(cacheKey)

	return nil
}

// 获得接口列表
func (m *Interface) GetList() ([]structs.Interface, error) {
	/*
		实现逻辑以下逻辑
		1. 获得接口列表.
		2. 数据结构是 structs.Interface
		3. 输出的字段根据 InterfaceList 结构来
		3. 根据 Path 字段顺序排序
	*/
	var interfaceList []structs.Interface
	if err := DB.Find(&interfaceList).Order("path asc").Error; err != nil {
		return nil, err
	}
	return interfaceList, nil
}

// 获得接口详情
func (m *Interface) GetInfo(InterfaceId string) (structs.Interface, error) {
	/*
		现逻辑以下逻辑
		1. 获得接口数据详情
		2. 数据结构是 structs.Interface
		3. 同时输出 Params 详细内容
	*/
	var interfaceInfo structs.Interface
	if err := DB.Model(&structs.Interface{}).Where("id = ?", InterfaceId).First(&interfaceInfo).Error; err != nil {
		return structs.Interface{}, cron.CreateCustomError(601, "接口不存在")
	}

	var params []structs.Params
	if err := DB.Model(&structs.Params{}).Where("interface_id = ?", InterfaceId).Find(&params).Error; err != nil {
		return structs.Interface{}, err
	}

	interfaceInfo.Params = params
	return interfaceInfo, nil
}

func (m *Interface) GetInfoByPath(path string, method string) (structs.Interface, error) {
	var interfaceInfo structs.Interface

	if strings.HasPrefix(path, "/api") {
		path = path[4:]
	}

	cacheKey := m.GetCacheKeyByPath(path, method)
	_Interface, IsExists := libraries.GetCache(cacheKey)
	if IsExists { //存在缓存 从缓存中获取
		return _Interface.(structs.Interface), nil
	}

	if err := DB.Model(&structs.Interface{}).Where("path = ? and method = ?", path, method).First(&interfaceInfo).Error; err != nil {
		return structs.Interface{}, cron.CreateCustomError(601, "接口不存在")
	}

	var params []structs.Params
	if err := DB.Model(&structs.Params{}).Where("interface_id = ?", interfaceInfo.Id).Find(&params).Error; err != nil {
		return structs.Interface{}, err
	}

	interfaceInfo.Params = params
	libraries.AddCache(cacheKey, interfaceInfo, 0)

	return interfaceInfo, nil
}

// 更新接口缓存
func (m *Interface) UpdateCacheByPath(interfaceId string) error {

	var interfaceInfo structs.Interface
	if err := DB.Model(&structs.Interface{}).Where("id = ?", interfaceId).First(&interfaceInfo).Error; err != nil {
		return cron.CreateCustomError(601, "接口不存在")
	}

	var params []structs.Params
	if err := DB.Model(&structs.Params{}).Where("interface_id = ?", interfaceId).Find(&params).Error; err != nil {
		return err
	}

	interfaceInfo.Params = params
	cacheKey := m.GetCacheKeyByPath(interfaceInfo.Path, interfaceInfo.Method)
	libraries.AddCache(cacheKey, interfaceInfo, 0)

	return nil
}

// 获得接口详情的缓存key
func (m *Interface) GetCacheKeyByPath(path string, method string) string {
	return "Interface_Info_" + helper.HashMD5(path+"_"+method)
}
