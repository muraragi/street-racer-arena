package repositories

import (
	"muraragi/street-racer-arena-backend/internal/models"

	"gorm.io/gorm"
)

type BaseCarRepository interface {
	FindAll() (*[]models.BaseCarModel, error)
}

type baseCarRepository struct {
	db *gorm.DB
}

func NewBaseCarRepository(db *gorm.DB) BaseCarRepository {
	return &baseCarRepository{db: db}
}

func (r *baseCarRepository) FindAll() (*[]models.BaseCarModel, error) {
	var cars []models.BaseCarModel
	if err := r.db.Find(&cars).Error; err != nil {
		return nil, err
	}
	return &cars, nil
}
