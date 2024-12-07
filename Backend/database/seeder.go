package database

import (
	"gorm.io/gorm"
	"car-rental-app/database/seeders"
)

func Seeder(db *gorm.DB) error {
	if err := seeders.UserSeeder(db); err != nil {
		return err
	}

	return nil
}

