package controllers

import (
	"tubes-pat/ticket-service/initializers"
	"tubes-pat/ticket-service/models"

	"github.com/gin-gonic/gin"
)

func BookingsCreate(seat_id uint) {
	booking := models.Booking{Seat_Id: seat_id, Status: "ON PROGRESS"}

	initializers.DB.Create(&booking)
}

func BookingUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Status string
	}

	c.Bind(&body)

	var booking models.Booking
	initializers.DB.First(&booking, id)

	initializers.DB.Model(&booking).Update("Status", body.Status)

	// Bad Coding Practices
	var seat models.Seat
	initializers.DB.First(&seat, booking.Seat_Id)
	if body.Status == "SUCCESS" {
		initializers.DB.Model(&seat).Update("Status", "BOOKED")
	} else {
		initializers.DB.Model(&seat).Update("Status", "OPEN")
	}

	c.JSON(200, gin.H{
		"message": "Booking Status Successfully Updated",
	})
}
