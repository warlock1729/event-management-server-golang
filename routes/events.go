package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/warlock1729/first-go-project/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events!"})
		return
	}
	// this is how a json reponse is returned using gin
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	// even though authorization and bearer words are just convention still we need to use them as they are expected in most tools and libraries

	var event models.Event

	// gin will read req body, decode the json,write to our variable
	// gin will read the struct tags on the struct type of the passed variable to enforce validation
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse body"})
		return
	}

	userID := context.GetInt64("userID")
	event.UserID = userID
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func getEventById(context *gin.Context) {
	var ID = context.Param("eventID")
	var intID, err = strconv.ParseInt(ID, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	event, err := models.GetEventByID(intID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event fetched successfully", "event": *event})
}

func updateEvent(context *gin.Context) {
	var ID = context.Param("eventID")
	var intID, err = strconv.ParseInt(ID, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	event, err := models.GetEventByID(intID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event!"})
		return
	}

	userID := context.GetInt64("userID")
	if event.UserID != userID {
		context.JSON(http.StatusForbidden, gin.H{"message": "Unauthorized access to event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse body!"})
		return
	}
	updatedEvent.ID = intID
	updatedEvent.UserID = userID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})

}

func deleteEvent(context *gin.Context) {
	var ID = context.Param("eventID")
	var intID, err = strconv.ParseInt(ID, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	event, err := models.GetEventByID(intID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event!"})
		return
	}

	userID := context.GetInt64("userID")
	if event.UserID != userID {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Forbidden!"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}

func registerEvent(context *gin.Context) {
	var ID = context.Param("eventID")
	var intID, err = strconv.ParseInt(ID, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	event, err := models.GetEventByID(intID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event!"})
		return
	}
	userID := context.GetInt64("userID")
	err = event.RegisterUser(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered for event"})
}

func cancelRegistration(context *gin.Context) {
	var ID = context.Param("eventID")
	var intID, err = strconv.ParseInt(ID, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	var event models.Event

	userID := context.GetInt64("userID")
	event.ID = intID

	err = event.CancelRegistration(userID)
	fmt.Println(err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration!"})
		return

	}

	context.JSON(http.StatusOK,gin.H{"message":"Registration cancelled"})
}
