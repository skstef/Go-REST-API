package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skstef/Go-REST-API/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerEvent)
	authenticated.POST("/events/:id/cancel", cancelRegistration)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
