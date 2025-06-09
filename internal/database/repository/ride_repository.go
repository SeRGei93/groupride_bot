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

func (r *RideRepository) FindNoReadyRideByUser(userID int64) (*table.Ride, error) {
	var row table.Ride
	err := r.database.Where("user_id = ?", userID).Order("id DESC").First(&row).Error
	if err != nil {
		return nil, err
	}
	return &row, nil
}

func (r *RideRepository) DeleteRide(item table.Ride) error {
	err := r.database.Delete(&item).Error
	return err
}

func (r *RideRepository) UpdateRide(item table.Ride) error {
	err := r.database.Save(&item).Error
	return err
}
