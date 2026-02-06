package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/warlock1729/first-go-project/middlewares"
)

func RegisterRoutes(app *gin.Engine) {
	app.POST("/signup", createUser)
	app.POST("/login", login)

	app.GET("/events", getEvents)
	app.GET("/events/:eventID", getEventById)

	authenticatedGroup := app.Group("/")
	authenticatedGroup.Use(middlewares.Authenticate)
	authenticatedGroup.POST("/event", createEvent)
	authenticatedGroup.PUT("/event/:eventID", updateEvent)
	authenticatedGroup.DELETE("/event/:eventID", deleteEvent)
	
	authenticatedGroup.POST("/event/:eventID/register",registerEvent)
	authenticatedGroup.DELETE("/event/:eventID/register",cancelRegistration)

}
