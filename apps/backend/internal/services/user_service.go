package services

import (
	"fmt"
	"muraragi/street-racer-arena-backend/internal/models"
	"muraragi/street-racer-arena-backend/internal/repositories"

	"github.com/markbates/goth"
)

type UserService interface {
	CreateUser(user goth.User) (*models.User, error)
	GetUserFromSession(user goth.User) (*models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) CreateUser(gothUser goth.User) (*models.User, error) {
	user, err := s.userRepository.FindByProvider(gothUser.Provider, gothUser.UserID)
	if err == nil {
		return user, nil
	}

	user = &models.User{
		Provider:   gothUser.Provider,
		ProviderID: gothUser.UserID,
		Username:   gothUser.Name,
		AvatarURL:  gothUser.AvatarURL,
		Cars:       []models.Car{},
	}

	user, err = s.userRepository.Create(user)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return user, nil
}

func (s *userService) GetUserFromSession(gothUser goth.User) (*models.User, error) {
	user, err := s.userRepository.FindByProvider(gothUser.Provider, gothUser.UserID)
	if err != nil {
		return nil, fmt.Errorf("error getting user from session: %w", err)
	}
	return user, nil
}
