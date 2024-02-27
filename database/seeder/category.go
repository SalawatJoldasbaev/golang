package seeder

import (
	"errors"
	"gorm.io/gorm"
	"salawat/api/core/entity"
)

func CategorySeeder(db *gorm.DB) error {
	var dummyCategory = []entity.Category{
		{
			Title: "Category test 1",
		},
		{
			Title: "Category test 2",
		},
		{
			Title: "Category test 3",
		},
	}

	hasTable := db.Migrator().HasTable(&entity.Category{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Category{}); err != nil {
			return err
		}
	}

	for _, data := range dummyCategory {
		var category entity.Category
		err := db.Where(&entity.Category{Title: data.Title}).First(&category).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&category, "title = ?", data.Title).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
