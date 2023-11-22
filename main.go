package main

import (
	"tubes-pat/ticket-service/controllers"
	"tubes-pat/ticket-service/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.ConnectToQueue()
}

func main() {
	defer initializers.Ch.Close()
	defer initializers.Conn.Close()

	r := gin.Default()

	// Route for events
	r.POST("/events", controllers.EventsCreate)

	r.GET("/events", controllers.EventsGetAll)
	r.GET("/events/:id", controllers.EventsGetById)

	r.PUT("/events/:id", controllers.EventsUpdate)

	r.DELETE("/events/:id", controllers.EventsDelete)

	// Route for seats
	r.POST("/seats", controllers.SeatsCreate)

	r.GET("/seats", controllers.SeatsGetAll)
	r.GET("/seats/:id", controllers.SeatsGetById)

	r.PUT("/seats/:id", controllers.SeatsUpdate)

	r.DELETE("/seats/:id", controllers.SeatsDelete)

	// Route for holding seats
	r.POST("/seats/hold/:id", controllers.SeatsHold)

	// WebHook API Call
	r.POST("/webhook/:id", controllers.BookingUpdate)

	// Route for Bookings
	r.GET("/bookings", controllers.BookingsGetAll)

	r.Run()
}
