package service

import (
	"car-rental-app/dto"
	"car-rental-app/models"
	"car-rental-app/repository"
	"context"

	"gorm.io/gorm"

	"github.com/mashingan/smapping"
)

type CarService interface {
	RegisterCar(ctx context.Context, carDTO dto.CarCreateDto) (models.Car, error)
	GetAllCar(ctx context.Context) ([]models.Car, error)
	GetCarByID(ctx context.Context, carID int) (models.Car, error)
	DeleteCar(ctx context.Context, carID int) error
	UpdateCar(ctx context.Context, carDTO dto.CarUpdateDto) error
	CheckCar(ctx context.Context, carName string) (bool, error)
	GetCarWithCategory(ctx context.Context, carID int) (*models.Car, error)
}

type carService struct {
	db            *gorm.DB
	carRepository repository.CarRepository
}

func NewCarService(cr repository.CarRepository) CarService {
	return &carService{
		carRepository: cr,
	}
}

func (cs *carService) RegisterCar(ctx context.Context, carDTO dto.CarCreateDto) (models.Car, error) {
	car := models.Car{}
	err := smapping.FillStruct(&car, smapping.MapFields(carDTO))
	if err != nil {
		return car, err
	}
	return cs.carRepository.RegisterCar(ctx, car)
}

func (cs *carService) GetAllCar(ctx context.Context) ([]models.Car, error) {
	return cs.carRepository.GetAllCar(ctx)
}

func (cs *carService) GetCarByID(ctx context.Context, carID int) (models.Car, error) {
	return cs.carRepository.FindCarByID(ctx, carID)
}

func (cs *carService) DeleteCar(ctx context.Context, carID int) error {
	return cs.carRepository.DeleteCar(ctx, carID)
}

func (cs *carService) UpdateCar(ctx context.Context, carDTO dto.CarUpdateDto) error {
	car := models.Car{}
	err := smapping.FillStruct(&car, smapping.MapFields(carDTO))
	if err != nil {
		return err
	}
	return cs.carRepository.UpdateCar(ctx, car)
}

func (cs *carService) CheckCar(ctx context.Context, carName string) (bool, error) {
	result, err := cs.carRepository.FindCarByName(ctx, carName)
	if err != nil {
		return false, err
	}

	if result.Name != "" {
		return false, nil
	}
	return true, nil
}

// In your service layer:
func (cs *carService) GetCarWithCategory(ctx context.Context, carID int) (*models.Car, error) {
	var car models.Car

	// Preload the Category and get the car by ID
	err := cs.db.Preload("Category").First(&car, carID).Error
	if err != nil {
		return nil, err
	}

	// Return the car along with the associated category
	return &car, nil
}
