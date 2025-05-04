package models

import (
	"gorm.io/gorm"
)

// BaseCarModel is the base model for all cars.
type BaseCarModel struct {
	gorm.Model
	Name         string  `gorm:"uniqueIndex;not null;size:100"`
	BasePower    float64 `gorm:"not null;default:100"`
	BaseHandling float64 `gorm:"not null;default:100"`
	CarInstances []Car   `gorm:"foreignKey:BaseCarModelID"`
}

// Car is the model for a car that a user can own.
type Car struct {
	gorm.Model
	UserID              uint                    `gorm:"index;not null"`
	BaseCarModelID      uint                    `gorm:"index;not null"`
	Nickname            string                  `gorm:"size:100"`
	User                User                    `gorm:"foreignKey:UserID"`
	BaseCarModel        BaseCarModel            `gorm:"foreignKey:BaseCarModelID"`
	InstalledComponents []InstalledCarComponent `gorm:"foreignKey:CarID"`
}

// Base car component
type CarComponent struct {
	gorm.Model
	Name             string  `gorm:"uniqueIndex;not null;size:100"`
	Description      string  `gorm:"type:text"`
	Type             string  `gorm:"index;not null;size:50"`
	PowerModifier    float64 `gorm:"default:0"`
	HandlingModifier float64 `gorm:"default:0"`
}

// Installed car component
type InstalledCarComponent struct {
	gorm.Model
	CarID          uint         `gorm:"index;not null"`
	CarComponentID uint         `gorm:"index;not null"`
	CarComponent   CarComponent `gorm:"foreignKey:CarComponentID"`
}
