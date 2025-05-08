package services

import (
	"muraragi/street-racer-arena-backend/internal/repositories"
)

type Services struct {
	BaseCarService BaseCarService
}

func InitializeServices(repositories *repositories.Repositories) *Services {
	return &Services{
		BaseCarService: NewBaseCarService(repositories.BaseCarRepository),
	}
}
