package main

import (
	"github.com/Prateek462003/trello-backend/controllers"
	"github.com/Prateek462003/trello-backend/database"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Activity struct {
	ID                  int    `json:"id"`
	TaskID              int    `json:"task_id"`
	ActivityName        string `json:"activity_name"`
	ActivityDescription string `json:"activity_description"`
}

func main() {

	database.Init()

	router := gin.Default()
	// Setup the routes
	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.CreateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	router.Run("localhost:8080")
}
