package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	// 连接数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建表
	db.AutoMigrate(&User{})

	// 插入数据
	user := User{Name: "张三"}
	db.Create(&user)

	// 查询数据
	var users []User
	db.Find(&users)
	fmt.Println(users)

	// 更新数据
	db.Model(&User{}).Where("name = ?", "张三").Update("name", "李四")

	// 删除数据
	db.Delete(&User{}, "name = ?", "李四")
}
