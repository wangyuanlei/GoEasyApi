package libraries

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create admin table and insert default user
	type Admin struct {
		Account  string
		Password string
	}
	db.AutoMigrate(&Admin{})
	db.Create(&Admin{Account: "admin", Password: "admin"})

	// Create config table
	db.AutoMigrate(struct {
		Key   string
		Value string
	}{})
	// Add a configuration. Currently using whitelist or blacklist
	db.Create(&struct {
		Key   string
		Value string
	}{Key: "ipFilterMode", Value: "whitelist"})

	// Create redis configuration
	db.AutoMigrate(struct {
		Ip       string
		Port     int
		Password string
		Db       int
	}{})

	// Add a redis configuration
	db.Create(&struct {
		Ip       string
		Port     int
		Password string
		Db       int
	}{Ip: "127.0.0.1", Port: 6379, Password: "Qwert!2024", Db: 0})

	// Create whitelist table
	db.AutoMigrate(struct {
		IP string
	}{})

	// Create blacklist table
	db.AutoMigrate(struct {
		IP string
	}{})

	return db
}
