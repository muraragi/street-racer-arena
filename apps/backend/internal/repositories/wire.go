package repositories

import "gorm.io/gorm"

type Repositories struct {
	BaseCarRepository BaseCarRepository
}

func InitializeRepositories(db *gorm.DB) *Repositories {
	baseCarRepository := NewBaseCarRepository(db)

	return &Repositories{
		BaseCarRepository: baseCarRepository,
	}
}
