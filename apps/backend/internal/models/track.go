package models

import (
	"gorm.io/gorm"
)

type LobbyTrack struct {
	gorm.Model
	TrackID uint  `gorm:"index;not null"`
	LobbyID uint  `gorm:"index;not null"`
	Track   Track `gorm:"foreignKey:TrackID"`
}

type Track struct {
	gorm.Model
	Name   string  `gorm:"uniqueIndex;not null;size:100"`
	Length float64 `gorm:"not null"`
	Type   string  `gorm:"index;not null;size:50"`
}
