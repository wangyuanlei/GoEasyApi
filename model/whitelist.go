package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/helper"
	"GoEasyApi/structs"
)

type WhiteList struct{}

// 获得白名单列表
func (m *WhiteList) GetAllWhiteList() ([]structs.WhiteList, error) {
	var list []structs.WhiteList
	err := DB.Find(&list).Error
	return list, err
}

// 添加白名单
func (m *WhiteList) AddWhiteList(ip string, description string) error {
	//判断ip
	if !helper.IsValidIP(ip) {
		return cron.CreateCustomError(601, "ip 不合法")
	}
	//判断ip 是否存在, 如果不存在 ,则创建, 存在则更新 Description 值
	var existingData structs.WhiteList
	DB.First(&existingData, "ip = ?", ip)
	if existingData.IP == "" {
		return DB.Create(&structs.WhiteList{IP: ip, Description: description}).Error
	} else {
		return DB.Model(&existingData).Where("ip = ?", ip).Update("description", description).Error
	}
}

// 删除白名单ip
func (m *WhiteList) DeleteWhiteList(ip string) error {
	var data structs.WhiteList
	return DB.Where("ip = ?", ip).Delete(&data).Error
}

// 获得黑名单列表
func (m *WhiteList) GetAllBlackList() ([]structs.BlackList, error) {
	var list []structs.BlackList
	err := DB.Find(&list).Error
	return list, err
}

// 添加黑名单
func (m *WhiteList) AddBlackList(ip string, description string) error {
	//判断ip
	if !helper.IsValidIP(ip) {
		return cron.CreateCustomError(601, "ip 不合法")
	}
	//判断ip 是否存在, 如果不存在 ,则创建, 存在则更新 Description 值
	var existingData structs.BlackList
	DB.First(&existingData, "ip = ?", ip)
	if existingData.IP == "" {
		return DB.Create(&structs.BlackList{IP: ip, Description: description}).Error
	} else {
		return DB.Model(&existingData).Where("ip = ?", ip).Update("description", description).Error
	}
}

// 删除黑名单ip
func (m *WhiteList) DeleteBlackList(ip string) error {
	var data structs.BlackList
	return DB.Where("ip = ?", ip).Delete(&data).Error
}
