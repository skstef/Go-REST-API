package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skstef/Go-REST-API/db"
	"github.com/skstef/Go-REST-API/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
