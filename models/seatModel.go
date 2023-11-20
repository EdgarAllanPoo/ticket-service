package models

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	Seat_Number uint
	Event_Id    uint
	Status      string
}
