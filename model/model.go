package model

import (
	"GoEasyApi/libraries"
	"GoEasyApi/structs"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 数据库初始化
func InitDB() *gorm.DB {
	dbPath, err := libraries.LoadDatabaseConfig()
	if err != nil {
		log.Fatal(err)
	}

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//建立数据表
	DB.AutoMigrate(&structs.WhiteList{}, &structs.BlackList{}, &structs.Database{}, &structs.Interface{}, &structs.Params{}, &structs.User{}, &structs.Token{})

	return DB
}

// 如果文件不存在, 则创建文件, 并且新建数据库
func CreateNewDBFile() {
	dbPath, err := libraries.LoadDatabaseConfig()
	if err != nil {
		log.Fatal(err)
	}

	// 判断文件是否存在.
	if _, err = os.Stat(dbPath); os.IsNotExist(err) {
		// 自动创建数据表
		DB.AutoMigrate(&structs.WhiteList{}, &structs.BlackList{}, &structs.Database{}, &structs.Interface{}, &structs.Params{}, &structs.User{}, &structs.Token{})
	}
}
