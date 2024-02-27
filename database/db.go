package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"salawat/api/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=joldasbaevsalawat password=Salawat&1909 dbname=golang port=5432 sslmode=disable TimeZone=Asia/Tashkent"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	err = db.AutoMigrate(models.Category{})
	if err != nil {
		return
	}

	DB = db
}
