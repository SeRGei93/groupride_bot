package table

import (
	"time"
)

// Ride — велогонка/мероприятие
type Ride struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint
	Active       bool `gorm:"default:true"`
	StartDate    time.Time
	Description  string            `gorm:"type:text"`
	Participants []RideParticipant `gorm:"constraint:OnDelete:CASCADE;"`
	Files        []File            `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
