package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	Seat_Id uint
	Status  string
}
