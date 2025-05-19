package repositories

import (
	"muraragi/street-racing-arena-backend/internal/models"

	"gorm.io/gorm"
)

type CarRepository interface {
	GetAll(userID uint) (*[]models.Car, error)
	Create(car *models.Car) (*models.Car, error)
	Update(car *models.Car) (*models.Car, error)
	Delete(id uint) error
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{db: db}
}

func (r *carRepository) Create(car *models.Car) (*models.Car, error) {
	if err := r.db.Create(car).Error; err != nil {
		return nil, err
	}
	return car, nil
}

func (r *carRepository) GetAll(userID uint) (*[]models.Car, error) {
	var cars []models.Car
	if err := r.db.Where("user_id = ?", userID).Find(&cars).Error; err != nil {
		return nil, err
	}
	return &cars, nil
}

func (r *carRepository) Update(car *models.Car) (*models.Car, error) {
	if err := r.db.Save(car).Error; err != nil {
		return nil, err
	}
	return car, nil
}

func (r *carRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Car{}, id).Error; err != nil {
		return err
	}
	return nil
}
