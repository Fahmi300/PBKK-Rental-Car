package database

import (
	"fmt"
	"car-rental-app/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		models.User{},
		models.Category{},
		models.Car{},
		models.Booking{},
	); err != nil {
		return err
	}
	fmt.Println("Migration success!")

	return nil
}