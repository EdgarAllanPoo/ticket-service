package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name    string
	API_URL string
}
