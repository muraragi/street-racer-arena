package repositories

import "gorm.io/gorm"

type Repositories struct {
	BaseCarRepository BaseCarRepository
	UserRepository    UserRepository
	CarRepository     CarRepository
}

func InitializeRepositories(db *gorm.DB) *Repositories {
	baseCarRepository := NewBaseCarRepository(db)
	userRepository := NewUserRepository(db)
	carRepository := NewCarRepository(db)

	return &Repositories{
		BaseCarRepository: baseCarRepository,
		UserRepository:    userRepository,
		CarRepository:     carRepository,
	}
}
