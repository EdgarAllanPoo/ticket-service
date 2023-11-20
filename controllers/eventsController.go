package controllers

import (
	"tubes-pat/ticket-service/initializers"
	"tubes-pat/ticket-service/models"

	"github.com/gin-gonic/gin"
)

func EventsCreate(c *gin.Context) {
	var body struct {
		Name    string
		API_URL string
	}

	c.Bind(&body)

	event := models.Event{Name: body.Name, API_URL: body.API_URL}

	result := initializers.DB.Create(&event)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(201, gin.H{
		"event": event,
	})
}

func EventsGetAll(c *gin.Context) {
	var events []models.Event
	result := initializers.DB.Find(&events)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"events": events,
	})
}

func EventsGetById(c *gin.Context) {
	id := c.Param("id")

	var event models.Event
	result := initializers.DB.First(&event, id)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"event": event,
	})
}

func EventsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name    string
		API_URL string
	}

	c.Bind(&body)

	var event models.Event
	initializers.DB.First(&event, id)

	initializers.DB.Model(&event).Updates(models.Event{
		Name:    body.Name,
		API_URL: body.API_URL,
	})

	c.JSON(200, gin.H{
		"event": event,
	})
}

func EventsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Event{}, id)

	c.Status(200)
}
