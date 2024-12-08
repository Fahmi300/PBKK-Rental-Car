package seeders

import (
	"encoding/json"
	"car-rental-app/models"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

func CarSeeder(db *gorm.DB) error {
	file, err := os.Open("database/seeders/data/car.json")
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

	var users []models.Car
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
