package services

import (
	"muraragi/street-racer-arena-backend/internal/repositories"
)

type Services struct {
	BaseCarService BaseCarService
	UserService    UserService
}

func InitializeServices(repositories *repositories.Repositories) *Services {
	return &Services{
		BaseCarService: NewBaseCarService(repositories.BaseCarRepository),
		UserService:    NewUserService(repositories.UserRepository),
	}
}
