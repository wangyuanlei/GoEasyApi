package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/database"
	"GoEasyApi/helper"
)

type WhiteList struct{}

// 获得白名单列表
func (m *WhiteList) GetAllWhiteList() ([]database.WhiteList, error) {
	var list []database.WhiteList
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
	var existingData database.WhiteList
	DB.First(&existingData, "ip = ?", ip)
	if existingData.IP == "" {
		return DB.Create(&database.WhiteList{IP: ip, Description: description}).Error
	} else {
		return DB.Model(&existingData).Update("description", description).Error
	}
}

// 删除白名单ip
func (m *WhiteList) DeleteWhiteList(ip string) error {
	var data database.WhiteList
	return DB.Where("ip = ?", ip).Delete(&data).Error
}

// 获得黑名单列表
func (m *WhiteList) GetAllBlackList() ([]database.BlackList, error) {
	var list []database.BlackList
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
	var existingData database.BlackList
	DB.First(&existingData, "ip = ?", ip)
	if existingData.IP == "" {
		return DB.Create(&database.BlackList{IP: ip, Description: description}).Error
	} else {
		return DB.Model(&existingData).Update("description", description).Error
	}
}

// 删除黑名单ip
func (m *WhiteList) DeleteBlackList(ip string) error {
	var data database.BlackList
	return DB.Where("ip = ?", ip).Delete(&data).Error
}
