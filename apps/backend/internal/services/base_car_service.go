package services

import (
	"muraragi/street-racing-arena-backend/internal/models"
	"muraragi/street-racing-arena-backend/internal/repositories"
)

type BaseCarService interface {
	GetAllBaseCars() (*[]models.BaseCarModel, error)
}

type baseCarService struct {
	baseCarRepository repositories.BaseCarRepository
}

func NewBaseCarService(baseCarRepository repositories.BaseCarRepository) BaseCarService {
	return &baseCarService{baseCarRepository: baseCarRepository}
}

func (s *baseCarService) GetAllBaseCars() (*[]models.BaseCarModel, error) {
	cars, err := s.baseCarRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return cars, nil
}
