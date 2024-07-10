package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/igor570/eventexplorer/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch events"})
	}
	context.JSON(http.StatusOK, events)
}

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64) //get id from param and cast int64

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event data", "error": err.Error()})
		return
	}

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not save Event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Message": "Event has been created!"})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64) //get id from param and cast int64

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event id"})
		return
	}

	_, err = models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch event"})
		return
	}

	var updatedEvent models.Event
	context.ShouldBindJSON(&updatedEvent)

	updatedEvent.ID = id
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event has been updated!"})

}
