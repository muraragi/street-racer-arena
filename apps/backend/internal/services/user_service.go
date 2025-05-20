package services

import (
	"fmt"
	"muraragi/street-racing-arena-backend/internal/models"
	"muraragi/street-racing-arena-backend/internal/repositories"

	"github.com/markbates/goth"
)

type UserService interface {
	CreateUser(user goth.User) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	SetSelectedCar(userID uint, carID uint) error
}

type userService struct {
	userRepository repositories.UserRepository
	carRepository  repositories.CarRepository
}

func NewUserService(userRepository repositories.UserRepository, carRepository repositories.CarRepository) UserService {
	return &userService{userRepository: userRepository, carRepository: carRepository}
}

func (s *userService) CreateUser(gothUser goth.User) (*models.User, error) {
	user, err := s.userRepository.FindByProvider(gothUser.Provider, gothUser.UserID)
	if err == nil {
		return user, nil
	}

	user = &models.User{
		Provider:   gothUser.Provider,
		ProviderID: gothUser.UserID,
		Username:   gothUser.NickName,
		AvatarURL:  gothUser.AvatarURL,
		Cars:       []models.Car{},
	}

	user, err = s.userRepository.Create(user)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return user, nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting user by ID: %w", err)
	}
	return user, nil
}

func (s *userService) SetSelectedCar(userID uint, carID uint) error {
	car, err := s.carRepository.GetByID(carID)
	if err != nil {
		return fmt.Errorf("error getting car by ID: %w", err)
	}

	if car.UserID != userID {
		return fmt.Errorf("car does not belong to user")
	}

	return s.userRepository.SetSelectedCar(userID, carID)
}
