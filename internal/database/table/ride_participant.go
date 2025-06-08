package table

import (
	"time"
)

type RideParticipant struct {
	UserID    int64 `gorm:"primaryKey"`
	RideID    uint  `gorm:"primaryKey"`
	Bike      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
