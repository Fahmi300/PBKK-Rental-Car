package service

import (
	"car-rental-app/dto"
	"car-rental-app/models"
	"car-rental-app/repository"
	"context"

	"github.com/mashingan/smapping"
)

type CarService interface {
	RegisterCar(ctx context.Context, carDTO dto.CarCreateDto) (models.Car, error)
	GetAllCar(ctx context.Context) ([]models.Car, error)
	GetCarByID(ctx context.Context, carID string) (models.Car, error)
	DeleteCar(ctx context.Context, carID string) error
	UpdateCar(ctx context.Context, carDTO dto.CarUpdateDto) error
	CheckCar(ctx context.Context, carName string) (bool, error)
}

type carService struct {
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

func (cs *carService) GetCarByID(ctx context.Context, carID string) (models.Car, error) {
	return cs.carRepository.FindCarByID(ctx, carID)
}

func (cs *carService) DeleteCar(ctx context.Context, carID string) error {
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
	car, err := cs.carRepository.FindCarByName(ctx, carName)
	if err != nil {
		return false, err
	}
	return car.Name != "", nil
}
