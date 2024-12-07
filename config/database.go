package config

import (
	"fmt"
	"car-rental-app/helpers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	dbUser := helpers.MustGetenv("DB_USER")
	dbPass := helpers.MustGetenv("DB_PASS")
	dbHost := helpers.MustGetenv("DB_HOST")
	dbName := helpers.MustGetenv("DB_NAME")
	dbPort := helpers.MustGetenvInt("DB_PORT")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// if err := db.AutoMigrate(
	// 	entity.User{},
	// ); err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbSQL.Close()
}