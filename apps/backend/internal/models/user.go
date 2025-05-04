package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"uniqueIndex;not null;size:50"`
	Email      string `gorm:"uniqueIndex;not null;size:255"`
	Provider   string `gorm:"index;size:50"`
	ProviderID string `gorm:"index;size:255"`
	AvatarURL  string `gorm:"size:512"`
	ProfileBio string `gorm:"type:text"`
	Score      int    `gorm:"default:1000"`
	RacesWon   uint   `gorm:"default:0"`
	TotalRaces uint   `gorm:"default:0"`
	Cars       []Car  `gorm:"foreignKey:UserID"`
}
