package libraries

import (
	"os"
	"testing"
)

func TestInitDatabase(t *testing.T) {
	InitDatabase()

	// Check if the database file exists
	dbPath, err := LoadDatabaseConfig()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		t.Fatal("数据库文件创建失败")
	}
}
