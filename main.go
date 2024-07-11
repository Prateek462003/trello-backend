package main

import (
	"os"

	"github.com/Prateek462003/trello-backend/controllers"
	"github.com/Prateek462003/trello-backend/database"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	database.Init()

	router := gin.Default()
	// Setup the routes
	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.CreateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
