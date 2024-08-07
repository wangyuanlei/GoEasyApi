package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// mysql的连接类
var DB *gorm.DB

func InitMysql() {
	var err error
	db, err := gorm.Open(mysql.Open("admin:Qwert!@#456@tcp(192.168.1.54:33066)/gwSysBase?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		fmt.Println("mysql连接失败")
		panic(err)
	}
	fmt.Println("mysql连接成功")

	DB = db
}

func GetList() []map[string]interface{} {
	var results []map[string]interface{}
	sql := "SELECT DeptId, DeptName, OrgId FROM Department limit 10"
	DB.Raw(sql).Scan(&results)
	return results
}
