package repository

import (
	"goupride_bot/internal/database/table"

	"gorm.io/gorm"
)

type RideParticipantRepository struct {
	database *gorm.DB
}

func NewRideParticipantRepository(db *gorm.DB) *RideParticipantRepository {
	return &RideParticipantRepository{database: db}
}

func (r *RideParticipantRepository) Init() *RideParticipantRepository {
	r.database.AutoMigrate(&table.RideParticipant{})
	return r
}

func (r *RideParticipantRepository) CreateRide(item table.RideParticipant) error {
	err := r.database.Create(&item).Error
	return err
}
