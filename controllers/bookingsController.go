package controllers

import (
	"time"
	"tubes-pat/ticket-service/initializers"
	"tubes-pat/ticket-service/models"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

func BookingsCreate(seat_id uint) {
	booking := models.Booking{Seat_Id: seat_id, Status: "ON PROGRESS"}

	initializers.DB.Create(&booking)
}

func BookingsGetAll(c *gin.Context) {
	var bookings []models.Booking
	initializers.DB.Find(&bookings)

	c.JSON(200, gin.H{
		"bookings": bookings,
	})
}

func BookingsGetById(c *gin.Context) {
	id := c.Param("id")

	var booking models.Booking
	initializers.DB.First(&booking, id)

	c.JSON(200, gin.H{
		"booking": booking,
	})
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

	// Post ke MQ
	PublishToQueue("BOOKING ID " + id + " SUKSES DIUBAH MENJADI " + body.Status)
}

func PublishToQueue(msg string) {
	err := initializers.Ch.Publish(
		"booking-result", // exchange
		"go-booking-key", // routing key
		false,            // mandatory
		false,            // immediate
		amqp091.Publishing{
			DeliveryMode: amqp091.Transient,
			ContentType:  "text/plain",
			Body:         []byte(msg),
			Timestamp:    time.Now(),
		})

	initializers.FailOnError(err, "Failed to Publish on RabbitMQ")
}
