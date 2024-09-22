package model

import (
	"GoEasyApi/database"
	"GoEasyApi/libraries"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

// 数据库初始化
func InitDB() *gorm.DB {
	dbPath, err := libraries.LoadDatabaseConfig()
	if err != nil {
		log.Fatal(err)
	}

	DB, err = gorm.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	//建立数据表
	DB.AutoMigrate(&database.WhiteList{}, &database.BlackList{}, &database.Database{}, &database.Interface{}, &database.Params{}, &database.User{}, &database.Token{})

	return DB
}

func CloseDB() {
	defer DB.Close()
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
		DB.AutoMigrate(&database.WhiteList{}, &database.BlackList{}, &database.Database{}, &database.Interface{}, &database.Params{}, &database.User{}, &database.Token{})
	}
}
