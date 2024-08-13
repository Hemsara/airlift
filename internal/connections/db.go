package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DBPath = "DB_PATH"
)

var DB *gorm.DB

func New() {
	if DB != nil {
		return
	}

	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc",
		os.Getenv(DBPath),
	)

	var err error
	d, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Fatalf("Error opening database: %v", err)
	}
	DB = d
}

func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
