package models

import (
	"gorm.io/gorm"
)

type WeatherCondition string

const (
	WeatherDry  WeatherCondition = "Dry"
	WeatherWet  WeatherCondition = "Wet"
	WeatherSnow WeatherCondition = "Snow"
)

type Lobby struct {
	gorm.Model
	LobbyTrackID     uint               `gorm:"index;not null"`
	Status           string             `gorm:"index;not null;size:50;default:'Waiting'"`
	MaxPlayers       uint               `gorm:"not null;default:8"`
	WeatherCondition WeatherCondition   `gorm:"type:varchar(10);not null;default:'Dry'"`
	LobbyTrack       LobbyTrack         `gorm:"foreignKey:LobbyTrackID"`
	Participants     []LobbyParticipant `gorm:"foreignKey:LobbyID"`
	RaceResult       *LobbyResult       `gorm:"foreignKey:LobbyID"`
}

type LobbyResult struct {
	gorm.Model
	LobbyID      uint               `gorm:"uniqueIndex;not null"`
	WinnerUserID uint               `gorm:"not null"`
	Results      []LobbyResultEntry `gorm:"foreignKey:LobbyResultID"`
	Lobby        Lobby              `gorm:"foreignKey:LobbyID"`
	Winner       User               `gorm:"foreignKey:WinnerUserID"`
}

type LobbyResultEntry struct {
	gorm.Model
	LobbyResultID uint `gorm:"index;not null"`
	UserID        uint `gorm:"not null"`
	CarID         uint `gorm:"not null"`
	Position      uint `gorm:"not null"`
	ScoreChange   int
	LapTime       float64 // Best lap time in seconds, optional
	TotalTime     float64 // Total race time in seconds, optional
	User          User    `gorm:"foreignKey:UserID"`
	Car           Car     `gorm:"foreignKey:CarID"`
}

type DrivingStrategy string

const (
	StrategyAggressive DrivingStrategy = "Aggressive"
	StrategyDefensive  DrivingStrategy = "Defensive"
	StrategyBalanced   DrivingStrategy = "Balanced"
	StrategyTechnical  DrivingStrategy = "Technical"
)

type LobbyParticipant struct {
	gorm.Model
	LobbyID          uint            `gorm:"index;not null"`
	UserID           uint            `gorm:"index;not null"`
	CarID            uint            `gorm:"index;not null"`
	IsReady          bool            `gorm:"default:false"`
	SelectedStrategy DrivingStrategy `gorm:"type:varchar(20);not null;default:'Balanced'"`
	InitialStanding  uint
	FinalPosition    *uint
	ScoreChange      *int
	Lobby            Lobby `gorm:"foreignKey:LobbyID"`
	User             User  `gorm:"foreignKey:UserID"`
	Car              Car   `gorm:"foreignKey:CarID"`
}
