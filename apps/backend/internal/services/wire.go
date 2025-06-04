package services

import (
	"muraragi/street-racing-arena-backend/internal/repositories"
)

type Services struct {
	BaseCarService BaseCarService
	UserService    UserService
	CarService     CarService
}

func InitializeServices(repositories *repositories.Repositories) *Services {
	return &Services{
		BaseCarService: NewBaseCarService(repositories.BaseCarRepository),
		UserService:    NewUserService(repositories.UserRepository, repositories.CarRepository),
		CarService:     NewCarService(repositories.CarRepository),
	}
}
