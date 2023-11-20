package controllers

import (
	"tubes-pat/ticket-service/initializers"
	"tubes-pat/ticket-service/models"
	"tubes-pat/ticket-service/utility"

	"github.com/gin-gonic/gin"
)

// Basic CRUD for Seats

func SeatsCreate(c *gin.Context) {
	var body struct {
		Seat_Number uint
		Event_Id    uint
		Status      string
	}

	c.Bind(&body)

	seat := models.Seat{Seat_Number: body.Seat_Number, Event_Id: body.Event_Id, Status: body.Status}

	result := initializers.DB.Create(&seat)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(201, gin.H{
		"seat": seat,
	})
}

func SeatsGetAll(c *gin.Context) {
	var seats []models.Seat
	result := initializers.DB.Find(&seats)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"seats": seats,
	})
}

func SeatsGetById(c *gin.Context) {
	id := c.Param("id")

	var seat models.Seat
	result := initializers.DB.First(&seat, id)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"seat": seat,
	})
}

func SeatsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Seat_Number uint
		Event_Id    uint
		Status      string
	}

	c.Bind(&body)

	var seat models.Seat
	initializers.DB.First(&seat, id)

	initializers.DB.Model(&seat).Updates(models.Seat{Seat_Number: body.Seat_Number, Event_Id: body.Event_Id, Status: body.Status})

	c.JSON(200, gin.H{
		"seat": seat,
	})
}

func SeatsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Seat{}, id)

	c.Status(200)
}

// Seats Hold Function
func SeatsHold(c *gin.Context) {
	id := c.Param("id")

	var seat models.Seat
	initializers.DB.First(&seat, id)

	if seat.Status == "BOOKED" {
		c.JSON(400, gin.H{
			"message": "Seat is not available",
		})
		return
	}

	if utility.SimulateAPICall() == "FAILED" {
		c.JSON(500, gin.H{
			"message": "Seat failed to hold",
		})
		return
	}

	// Invoice ke Payment App

	BookingsCreate(utility.StrToUint(id))
	initializers.DB.Model(&seat).Update("Status", "ON GOING")

	c.JSON(200, gin.H{
		"message": "Seat successfully hold, waiting for payment",
		"seat":    seat,
	})
}
