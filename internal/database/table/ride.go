package table

import (
	"time"
)

// Ride — велогонка/мероприятие
type Ride struct {
	ID           uint `gorm:"primaryKey"`
	UserID       int64
	Active       bool `gorm:"default:true"`
	Ready        bool `gorm:"default:false"` // true - после того как пользователь заполнил все данные
	StartDate    time.Time
	Link         string
	Description  string            `gorm:"type:text"`
	Participants []RideParticipant `gorm:"constraint:OnDelete:CASCADE;"`
	Files        []File            `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
