package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/database"
	"regexp"
	"strings"

	"github.com/cengsin/oracle"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type ApiModel struct {
	Interface database.Interface
	DBType    string
	DB        *gorm.DB
}

// 初始化
func (m *ApiModel) Init() error {

	var DBModel = DataBase{}
	dbConfig, err := DBModel.GetUserDBConf()
	if err != nil {
		return err
	}

	m.DBType = dbConfig.OrmType

	//获得数据库连接句柄
	if m.DBType == "mysql" {
		//dsn := "username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		m.DB, err = gorm.Open(mysql.Open(dbConfig.Dns), &gorm.Config{})
		if err != nil {
			panic("failed to connect database:")
		}
	} else if m.DBType == "postgresql" {
		//dsn := "host=localhost user=your_username password=your_password dbname=your_db port=5432 sslmode=disable"
		m.DB, err = gorm.Open(postgres.Open(dbConfig.Dns), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	} else if m.DBType == "sqlite" {
		//dsn := "test.db"
		m.DB, err = gorm.Open(sqlite.Open(dbConfig.Dns), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	} else if m.DBType == "sqlserver" {
		//dsn := "sqlserver://username:password@localhost:1433?database=dbname"
		m.DB, err = gorm.Open(sqlserver.Open(dbConfig.Dns), &gorm.Config{})
		if err != nil {
			panic("failed to connect to database")
		}
	} else if m.DBType == "oracle" {
		//dsn := "system/oracle@127.0.0.1:1521/XE"
		m.DB, err = gorm.Open(oracle.Open(dbConfig.Dns), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	} else {
		return cron.CreateCustomError(602, "数据库类型错误")
	}

	return nil
}

func (m *ApiModel) Run(Interface database.Interface, params map[string]string) (interface{}, error) {
	m.Interface = Interface
	if m.Interface.ReturnType == "insert" {
		return m.Insert(Interface.SqlContent, params)
	} else if m.Interface.ReturnType == "update" {
		return m.Update(Interface.SqlContent, params)
	} else if m.Interface.ReturnType == "delete" {
		return m.Delete(Interface.SqlContent, params)
	} else if m.Interface.ReturnType == "row" {
		return m.GetOne(Interface.SqlContent, params)
	} else if m.Interface.ReturnType == "list" {
		return m.GetList(Interface.SqlContent, params)
	} else {
		return nil, cron.CreateCustomError(602, "返回类型错误")
	}
}

// 插入数据
func (m *ApiModel) Insert(sql_content string, params map[string]string) (string, error) {
	newSql, newParams, err := m.HandleSql(sql_content, params)
	if err != nil {
		return "", err
	}
	if err := m.DB.Exec(newSql, newParams...).Error; err != nil {
		return "", err
	}

	var id string
	if m.Interface.ReturnType == "insert" && m.Interface.ReturnValMode == "last_id" {
		if m.DBType == "mysql" {
			if err := m.DB.Raw("SELECT LAST_INSERT_ID()").Row().Scan(&id); err != nil {
				return "", err
			}
		} else if m.DBType == "pgsql" {
			if err := m.DB.Raw("SELECT LASTVAL()").Row().Scan(&id); err != nil {
				return "", err
			}
		}

		return id, nil
	}

	return "success", nil
}

// 修改数据
func (m *ApiModel) Update(sql_content string, params map[string]string) ([]interface{}, error) {
	newSql, newParams, err := m.HandleSql(sql_content, params)
	if err != nil {
		return nil, err
	}
	if err := m.DB.Exec(newSql, newParams...).Error; err != nil {
		return nil, err
	}

	if RowsAffected := m.DB.RowsAffected; RowsAffected == 0 {
		return nil, cron.CreateCustomError(602, "更新数据失败")
	}

	if m.Interface.ReturnValMode == "row" {
		var data []interface{}
		if err := m.DB.Last(&data).Error; err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, nil
}

// 删除数据
func (m *ApiModel) Delete(sql_content string, params map[string]string) (string, error) {
	newSql, newParams, err := m.HandleSql(sql_content, params)
	if err != nil {
		return "", err
	}
	if err := m.DB.Exec(newSql, newParams...).Error; err != nil {
		return "", err
	}

	return "success", nil
}

// 获得单条数据
func (m *ApiModel) GetOne(sql_content string, params map[string]string) ([]interface{}, error) {
	newSql, newParams, err := m.HandleSql(sql_content, params)
	if err != nil {
		return nil, err
	}
	var data []interface{}
	if err := m.DB.Raw(newSql, newParams...).Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// 获得多条数据
func (m *ApiModel) GetList(sql_content string, params map[string]string) ([]interface{}, error) {
	newSql, newParams, err := m.HandleSql(sql_content, params)
	if err != nil {
		return nil, err
	}
	var data []interface{}
	if err := m.DB.Raw(newSql, newParams...).Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// 处理sql 语句中的 key值 和 对应的参数
func (m *ApiModel) HandleSql(sql_content string, params map[string]string) (newSql string, paramValues []interface{}, err error) {
	re := regexp.MustCompile(`{{(.*?)}}`)
	keys := re.FindAllStringSubmatch(sql_content, -1)

	newSql = sql_content
	for _, key := range keys {
		paramValues = append(paramValues, params[key[1]])

		if params[key[1]] == "" {
			return "", []interface{}{}, cron.CreateCustomError(602, "参数错误")
		}

		paramValues = append(paramValues, params[key[1]])

		newSql = strings.Replace(newSql, "{{"+key[1]+"}}", "?", -1)
	}

	return newSql, paramValues, nil
}
