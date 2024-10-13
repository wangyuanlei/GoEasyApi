package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/database"
	"GoEasyApi/helper"

	"github.com/google/uuid"
)

type DataBase struct{}

// 获得数据库配置信息
func (a *DataBase) GetUserDBConf() (database.Database, error) {
	var condfig database.Database
	err := DB.First(&condfig).Error
	return condfig, err
}

// 修改数据库配置信息
func (a *DataBase) UpdateUserDBConf(name string, description string, orm_type string, dns string) error {

	supportedOrmTypes := []string{"mysql", "postgresql", "oracle", "sqlserver", "sqlite"}
	if !helper.InArray(supportedOrmTypes, orm_type) {
		return cron.CreateCustomError(601, "类型不支持")
	}

	dbconf, _ := a.GetUserDBConf()
	if dbconf.DatabaseId == "" { //增加
		dbconf.DatabaseId = uuid.New().String()
		dbconf.DatabaseName = name
		dbconf.Description = description
		dbconf.OrmType = orm_type
		dbconf.Dns = dns

		return DB.Create(&dbconf).Error
	} else { //修改
		dbconf.DatabaseName = name
		dbconf.Description = description
		dbconf.OrmType = orm_type
		dbconf.Dns = dns
		return DB.Save(&dbconf).Error
	}
}
