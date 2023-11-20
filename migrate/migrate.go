package main

import (
	"tubes-pat/ticket-service/initializers"
	"tubes-pat/ticket-service/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Event{})
	initializers.DB.AutoMigrate(&models.Seat{})
	initializers.DB.AutoMigrate(&models.Booking{})
}
