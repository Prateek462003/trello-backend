package main

import (
	"github.com/Prateek462003/trello-backend/controllers"
	"github.com/Prateek462003/trello-backend/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	database.Init()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	router.GET("/activities", controllers.GetActivities)
	router.POST("/activities", controllers.CreateActivity)
	router.GET("/activities/:activity_id", controllers.GetTasksByActivity)

	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.CreateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
