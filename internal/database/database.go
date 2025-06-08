package database

import (
	"fmt"
	Repository "goupride_bot/internal/database/repository"

	"gorm.io/gorm"
)

type Database struct {
	User            Repository.UserRepository
	Ride            Repository.RideRepository
	File            Repository.FileRepository
	RideParticipant Repository.RideParticipantRepository
}

func InitDatabase(dialector gorm.Dialector) Database {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	return Database{
		User:            *Repository.NewUserRepository(db).Init(),
		Ride:            *Repository.NewRideRepository(db).Init(),
		File:            *Repository.NewFileRepository(db).Init(),
		RideParticipant: *Repository.NewRideParticipantRepository(db).Init(),
	}
}
