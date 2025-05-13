package repositories

import "gorm.io/gorm"

type Repositories struct {
	BaseCarRepository BaseCarRepository
	UserRepository    UserRepository
}

func InitializeRepositories(db *gorm.DB) *Repositories {
	baseCarRepository := NewBaseCarRepository(db)
	userRepository := NewUserRepository(db)

	return &Repositories{
		BaseCarRepository: baseCarRepository,
		UserRepository:    userRepository,
	}
}
