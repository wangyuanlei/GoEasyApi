package libraries

import (
	"GoEasyApi/database"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// 数据库初始化
func InitDB() (*gorm.DB, error) {
	dbPath, err := LoadDatabaseConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db, nil
}

// 如果文件不存在, 则创建文件, 并且新建数据库
func CreateNewDBFile() {
	dbPath, err := LoadDatabaseConfig()
	if err != nil {
		log.Fatal(err)
	}

	// 判断文件是否存在.
	if _, err = os.Stat(dbPath); os.IsNotExist(err) {
		db, err := InitDB()
		if err != nil {
			log.Fatal(err)
		}
		// 自动创建数据表
		db.AutoMigrate(&database.WhiteList{}, &database.BlackList{}, &database.Database{}, &database.Interface{}, &database.Params{}, &database.User{}, &database.Token{})
	}
}
