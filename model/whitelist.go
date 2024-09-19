package model

import (
	"GoEasyApi/database"

	"github.com/jinzhu/gorm"
)

type WhiteList struct{}

// 获得白名单列表
func (m *WhiteList) GetAllWhiteList(db *gorm.DB) ([]database.WhiteList, error) {
	var list []database.WhiteList
	db.Find(&list)
	return list, nil
}

// 添加白名单
func (m *WhiteList) AddWhiteList(db *gorm.DB, data database.WhiteList) {
	//判断ip 是否存在, 如果不存在 ,则创建, 存在则更新 Description 值
	var existingData database.WhiteList
	db.First(&existingData, "ip = ?", data.IP)
	if existingData.IP == "" {
		db.Create(data)
	} else {
		db.Model(existingData).Update("description", data.Description)
	}
}

// 删除白名单ip
func (m *WhiteList) DeleteWhiteList(db *gorm.DB, ip string) error {
	var data database.WhiteList
	db.Where("ip = ?", ip).Delete(&data)
	return nil
}

// 获得黑名单列表
func (m *WhiteList) GetAllBlackList(db *gorm.DB) ([]database.BlackList, error) {
	var list []database.BlackList
	db.Find(&list)
	return list, nil
}

// 添加黑名单
func (m *WhiteList) AddBlackList(db *gorm.DB, data database.BlackList) {
	var existingData database.BlackList
	db.First(&existingData, "ip = ?", data.IP)
	if existingData.IP == "" {
		db.Create(data)
	} else {
		db.Model(existingData).Update("description", data.Description)
	}
}

// 删除黑名单ip
func (m *WhiteList) DeleteBlackList(db *gorm.DB, ip string) error {
	var data database.BlackList
	db.Where("ip = ?", ip).Delete(&data)
	return nil
}
