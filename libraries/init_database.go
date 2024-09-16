package libraries

import (
	"GoEasyApi/database"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDatabase() {
	dbPath, err := LoadDatabaseConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Check if the database file exists
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		// If the database file does not exist, create it and initialize the tables
		db, err := gorm.Open("sqlite3", dbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Auto migrate the tables
		db.AutoMigrate(&database.WhiteList{}, &database.BlackList{}, &database.Database{}, &database.Interface{}, &database.Params{}, &database.User{}, &database.Token{})
	}
}
