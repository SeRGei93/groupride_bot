package repository

import (
	"goupride_bot/internal/database/table"

	"gorm.io/gorm"
)

type RideRepository struct {
	database *gorm.DB
}

func NewRideRepository(db *gorm.DB) *RideRepository {
	return &RideRepository{database: db}
}

func (r *RideRepository) Init() *RideRepository {
	r.database.AutoMigrate(&table.File{})
	return r
}

func (r *RideRepository) CreateRide(item table.Ride) error {
	err := r.database.Create(&item).Error
	return err
}
