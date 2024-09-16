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
	db, err := gorm.Open(mysql.Open("admin:Qwert!@#456@tcp(43.254.227.159:33066)/gwSysBase?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		fmt.Println("mysql连接失败")
		panic(err)
	}
	fmt.Println("mysql连接成功")

	DB = db
}

// PageResult 定义分页的数据结构体
type PageResult struct {
	PageNo    int                      `json:"pageNo"`    // 当前页码
	PageCount int                      `json:"pageCount"` // 总页数
	DataCount int                      `json:"dataCount"` // 总数据条数
	Data      []map[string]interface{} `json:"data"`      // 当前页数据
}

func GetList() []map[string]interface{} {
	var results []map[string]interface{}
	sql := "SELECT DeptId, DeptName, OrgId FROM Department limit 10"
	DB.Raw(sql).Scan(&results)
	return results
}

func GetPage(sql string, pageNo int, pageSize int) PageResult {
	var results []map[string]interface{}
	var totalCount int64
	offset := (pageNo - 1) * pageSize

	// Get total count
	DB.Raw(fmt.Sprintf("SELECT COUNT(*) FROM (%s) as count_table", sql)).Scan(&totalCount)

	// Get paginated data
	paginatedSQL := fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, pageSize, offset)
	DB.Raw(paginatedSQL).Scan(&results)

	// Calculate total pages
	pageCount := int((totalCount + int64(pageSize) - 1) / int64(pageSize))

	return PageResult{
		PageNo:    pageNo,
		PageCount: pageCount,
		DataCount: int(totalCount),
		Data:      results,
	}
}

func GetListPage(pageNo int, pageSize int) PageResult {
	var results []map[string]interface{}
	var totalCount int64
	offset := (pageNo - 1) * pageSize

	// Get total count
	DB.Table("Department").Count(&totalCount)

	// Get paginated data
	sql := fmt.Sprintf("SELECT DeptId, DeptName, OrgId FROM Department LIMIT %d OFFSET %d", pageSize, offset)
	DB.Raw(sql).Scan(&results)

	// Calculate total pages
	pageCount := int((totalCount + int64(pageSize) - 1) / int64(pageSize))

	return PageResult{
		PageNo:    pageNo,
		PageCount: pageCount,
		DataCount: int(totalCount),
		Data:      results,
	}
}
