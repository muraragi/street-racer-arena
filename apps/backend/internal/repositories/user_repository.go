package repositories

import (
	"errors"
	"muraragi/street-racing-arena-backend/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByProvider(providerName string, providerUserID string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	SetSelectedCar(userID uint, carID uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) FindByProvider(providerName string, providerUserID string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("provider = ? AND provider_id = ?", providerName, providerUserID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) SetSelectedCar(userID uint, carID uint) error {
	return r.DB.Model(&models.User{}).Where("id = ?", userID).Update("selected_car_id", carID).Error
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.Preload("Cars").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
