package main

import (
	// "car-rental-app/database"
	"car-rental-app/models"
	"car-rental-app/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database

	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	// Auto-migrate models
	db.AutoMigrate(&models.User{},  &models.Category{}, &models.Car{}, &models.Booking{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Car Rental App!"})
	})

	r.Run(":8080")
}
