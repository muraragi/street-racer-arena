package services

import (
	"muraragi/street-racing-arena-backend/internal/models"
	"muraragi/street-racing-arena-backend/internal/repositories"
)

type CarService interface {
	AddCarToUser(car *models.Car) (*models.Car, error)
	GetUsersCars(userID uint) (*[]models.Car, error)
	UpdateUserCar(car *models.Car) (*models.Car, error)
	DeleteUserCar(carID uint) error
}

type carService struct {
	carRepository repositories.CarRepository
}

func NewCarService(carRepository repositories.CarRepository) CarService {
	return &carService{carRepository: carRepository}
}

func (s *carService) AddCarToUser(car *models.Car) (*models.Car, error) {
	return s.carRepository.Create(car)
}

func (s *carService) GetUsersCars(userID uint) (*[]models.Car, error) {
	return s.carRepository.GetAll(userID)
}

func (s *carService) UpdateUserCar(car *models.Car) (*models.Car, error) {
	return s.carRepository.Update(car)
}

func (s *carService) DeleteUserCar(carID uint) error {
	return s.carRepository.Delete(carID)
}
