package repository

import (
	"car-rental-app/models"
	"context"

	"gorm.io/gorm"
)

type CarRepository interface {
	RegisterCar(ctx context.Context, car models.Car) (models.Car, error)
	GetAllCar(ctx context.Context) ([]models.Car, error)
	FindCarByID(ctx context.Context, carID int) (models.Car, error)
	FindCarByName(ctx context.Context, carName string) (models.Car, error)
	DeleteCar(ctx context.Context, carID int) error
	UpdateCar(ctx context.Context, car models.Car) error
}

type carConnection struct {
	connection *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carConnection{
		connection: db,
	}
}

func (db *carConnection) RegisterCar(ctx context.Context, car models.Car) (models.Car, error) {
	result := db.connection.Create(&car)
	if result.Error != nil {
		return models.Car{}, result.Error
	}
	return car, nil
}

func (db *carConnection) GetAllCar(ctx context.Context) ([]models.Car, error) {
	var cars []models.Car
	tx := db.connection.Preload("Category").Find(&cars) // Preload Category
	if tx.Error != nil {
		return nil, tx.Error
	}
	return cars, nil
}

func (db *carConnection) FindCarByID(ctx context.Context, carID int) (models.Car, error) {
	var car models.Car
	tx := db.connection.Preload("Category").Where("id = ?", carID).Take(&car) // Preload Category
	if tx.Error != nil {
		return car, tx.Error
	}
	return car, nil
}

func (db *carConnection) FindCarByName(ctx context.Context, carName string) (models.Car, error) {
	var car models.Car
	tx := db.connection.Where("name = ?", carName).Take(&car)
	if tx.Error != nil {
		return car, tx.Error
	}
	return car, nil
}

func (db *carConnection) DeleteCar(ctx context.Context, carID int) error {
	tx := db.connection.Delete(&models.Car{}, carID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (db *carConnection) UpdateCar(ctx context.Context, car models.Car) error {
	tx := db.connection.Updates(&car)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
