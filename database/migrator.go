package database

import (
	"fmt"
	"gorm.io/gorm"
	"salawat/api/core/entity"
	"salawat/api/database/seeder"
)

func DBMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.Category{},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := DBSeed(db); err != nil {
		panic(err)
	}
}

func DBSeed(db *gorm.DB) error {
	if err := seeder.CategorySeeder(db); err != nil {
		return err
	}

	return nil
}
