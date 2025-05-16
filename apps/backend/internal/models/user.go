package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `gorm:"uniqueIndex;not null;size:50"`
	Provider      string `gorm:"index;size:50"`
	ProviderID    string `gorm:"index;size:255"`
	AvatarURL     string `gorm:"size:512"`
	ProfileBio    string `gorm:"type:text"`
	Score         int    `gorm:"default:1000"`
	RacesWon      uint   `gorm:"default:0"`
	TotalRaces    uint   `gorm:"default:0"`
	Cars          []Car  `gorm:"foreignKey:UserID"`
	SelectedCarID *uint  `gorm:"default:null;constraint:OnDelete:SET NULL,OnUpdate:CASCADE;foreignKey:SelectedCarID;references:ID"`
}

func (u *User) BeforeSave(db *gorm.DB) error {
	if u.SelectedCarID == nil {
		return nil
	}

	var count int64
	err := db.
		Model(&Car{}).
		Where("id = ? AND user_id = ?", *u.SelectedCarID, u.ID).
		Count(&count).
		Error

	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("Car doesn't belong to user")
	}

	return nil
}
