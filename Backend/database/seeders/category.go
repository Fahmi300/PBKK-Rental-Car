package seeders

import (
	"encoding/json"
	"car-rental-app/models"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

func CategorySeeder(db *gorm.DB) error {
	file, err := os.Open("database/seeders/data/category.json")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)	
	if err != nil {
		log.Fatal(err)
		return err
	}

	var users []models.Category
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, user := range users {
		if err != nil {
			return err
		}

		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
