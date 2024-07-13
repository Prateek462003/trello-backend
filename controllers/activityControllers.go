package controllers

import (
	"log"
	"net/http"

	models "github.com/Prateek462003/trello-backend/Models"
	"github.com/Prateek462003/trello-backend/database"
	"github.com/gin-gonic/gin"
)

func GetActivities(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	rows, err := database.DB.Query("SELECT id, name FROM activities")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var activities []models.Activity
	for rows.Next() {
		var activity models.Activity
		if err := rows.Scan(&activity.ID, &activity.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		activities = append(activities, activity)
	}
	c.JSON(http.StatusOK, activities)
}

func CreateActivity(c *gin.Context) {
	var activity models.Activity
	if err := c.BindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec("INSERT INTO activities (name) VALUES ($1)", activity.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func GetTasksByActivity(c *gin.Context) {
	activityID := c.Param("activity_id")

	rows, err := database.DB.Query("SELECT id, title, description, image, activity_id FROM tasks WHERE activity_id = $1", activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Image, &task.ActivityId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
